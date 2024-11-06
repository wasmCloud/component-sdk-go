package wasihttp

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bytecodealliance/wasm-tools-go/cm"
	monotonicclock "go.wasmcloud.dev/component/gen/wasi/clocks/monotonic-clock"
	outgoinghandler "go.wasmcloud.dev/component/gen/wasi/http/outgoing-handler"
	"go.wasmcloud.dev/component/gen/wasi/http/types"
)

// Transport implements [http.RoundTripper] for [wasi:http].
//
// [wasi:http]: https://github.com/WebAssembly/wasi-http/tree/v0.2.0
type Transport struct {
	ConnectTimeout time.Duration
}

var _ http.RoundTripper = (*Transport)(nil)

// DefaultTransport is the default implementation of [Transport] and is used by [DefaultClient].
// It is configured use the same timeout value as [net/http.DefaultTransport].
var DefaultTransport = &Transport{
	ConnectTimeout: 30 * time.Second, // NOTE(lxf): Same as stdlib http.Transport
}

// DefaultClient is the default [net/http.Client] that uses [DefaultTransport] to adapt [net/http] to [wasi:http].
//
// [wasi:http]: https://github.com/WebAssembly/wasi-http/tree/v0.2.0
var DefaultClient = &http.Client{Transport: DefaultTransport}

func (r *Transport) requestOptions() types.RequestOptions {
	options := types.NewRequestOptions()
	options.SetConnectTimeout(cm.Some(monotonicclock.Duration(r.ConnectTimeout)))
	return options
}

// RoundTrip implements the [net/http.RoundTripper] interface.
func (r *Transport) RoundTrip(incomingRequest *http.Request) (*http.Response, error) {
	var err error

	outHeaders := types.NewFields()
	if err := HTTPtoWASIHeader(incomingRequest.Header, outHeaders); err != nil {
		return nil, err
	}

	outRequest := types.NewOutgoingRequest(outHeaders)

	outRequest.SetAuthority(cm.Some(incomingRequest.Host))
	outRequest.SetMethod(toWasiMethod(incomingRequest.Method))

	pathWithQuery := incomingRequest.URL.Path
	if incomingRequest.URL.RawQuery != "" {
		pathWithQuery = pathWithQuery + "?" + incomingRequest.URL.Query().Encode()
	}
	outRequest.SetPathWithQuery(cm.Some(pathWithQuery))

	switch incomingRequest.URL.Scheme {
	case "http":
		outRequest.SetScheme(cm.Some(types.SchemeHTTP()))
	case "https":
		outRequest.SetScheme(cm.Some(types.SchemeHTTPS()))
	default:
		outRequest.SetScheme(cm.Some(types.SchemeOther(incomingRequest.URL.Scheme)))
	}

	var adaptedBody io.WriteCloser
	var body *types.OutgoingBody
	if incomingRequest.Body != nil {
		bodyRes := outRequest.Body()
		if bodyRes.IsErr() {
			return nil, fmt.Errorf("failed to acquire resource handle to request body: %s", bodyRes.Err())
		}
		body = bodyRes.OK()
		adaptedBody, err = NewOutgoingBody(body)
		if err != nil {
			return nil, fmt.Errorf("failed to adapt body: %s", err)
		}
	}

	handleResp := outgoinghandler.Handle(outRequest, cm.Some(r.requestOptions()))
	if handleResp.Err() != nil {
		return nil, fmt.Errorf("%v", handleResp.Err())
	}

	// NOTE(lxf): If request includes a body, copy it to the adapted wasi body
	if body != nil {
		if _, err := io.Copy(adaptedBody, incomingRequest.Body); err != nil {
			return nil, fmt.Errorf("failed to copy body: %v", err)
		}

		if err := adaptedBody.Close(); err != nil {
			return nil, fmt.Errorf("failed to close body: %v", err)
		}

		outTrailers := types.NewFields()
		if err := HTTPtoWASIHeader(incomingRequest.Trailer, outTrailers); err != nil {
			return nil, err
		}

		maybeTrailers := cm.None[types.Fields]()
		if len(incomingRequest.Trailer) > 0 {
			maybeTrailers = cm.Some(outTrailers)
		}

		outFinish := types.OutgoingBodyFinish(*body, maybeTrailers)
		if outFinish.IsErr() {
			return nil, fmt.Errorf("failed to finish body: %v", outFinish.Err())
		}
	}

	// NOTE(lxf): Request is fully sent. Processing response.
	futureResponse := handleResp.OK()

	// wait until resp is returned
	futureResponse.Subscribe().Block()

	pollableOption := futureResponse.Get()
	if pollableOption.None() {
		return nil, fmt.Errorf("incoming resp is None")
	}

	pollableResult := pollableOption.Some()
	if pollableResult.IsErr() {
		return nil, fmt.Errorf("error is %v", pollableResult.Err())
	}

	resultOption := pollableResult.OK()
	if resultOption.IsErr() {
		return nil, fmt.Errorf("%v", resultOption.Err())
	}

	incomingResponse := resultOption.OK()
	incomingBody, incomingTrailers, err := NewIncomingBodyTrailer(incomingResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to consume incoming request %s", err)
	}

	incomingHeaders := http.Header{}
	headers := incomingResponse.Headers()
	WASItoHTTPHeader(headers, &incomingHeaders)
	headers.ResourceDrop()

	resp := &http.Response{
		StatusCode: int(incomingResponse.Status()),
		Status:     http.StatusText(int(incomingResponse.Status())),
		Request:    incomingRequest,
		Header:     incomingHeaders,
		Body:       incomingBody,
		Trailer:    incomingTrailers,
	}

	return resp, nil
}
