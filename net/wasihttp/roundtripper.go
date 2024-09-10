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

// Transport implements http.RoundTripper
type Transport struct {
	ConnectTimeout time.Duration
}

var _ http.RoundTripper = (*Transport)(nil)

var (
	DefaultTransport = &Transport{
		ConnectTimeout: 30 * time.Second,
	}
	DefaultClient = &http.Client{Transport: DefaultTransport}
)

func (r *Transport) requestOptions() types.RequestOptions {
	options := types.NewRequestOptions()
	options.SetConnectTimeout(cm.Some(monotonicclock.Duration(r.ConnectTimeout)))
	return options
}

func (r *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	or, err := NewOutgoingHttpRequest(req)
	if err != nil {
		return nil, err
	}

	var adaptedBody io.WriteCloser
	var body types.OutgoingBody
	if req.Body != nil {
		bodyRes := or.Body()
		if bodyRes.IsErr() {
			return nil, fmt.Errorf("failed to acquire resource handle to request body: %s", bodyRes.Err())
		}

		body = *bodyRes.OK()

		adaptedBody, err = NewOutgoingBody(body)
		if err != nil {
			return nil, fmt.Errorf("failed to adapt body: %s", err)
		}
	}

	handleResp := outgoinghandler.Handle(or, cm.Some(r.requestOptions()))
	if handleResp.Err() != nil {
		return nil, fmt.Errorf("%v", handleResp.Err())
	}

	if adaptedBody != nil {
		if _, err := io.Copy(adaptedBody, req.Body); err != nil {
			return nil, fmt.Errorf("failed to copy body: %s", err)
		}

		trailers := types.NewFields()
		if err := toWasiHeader(req.Trailer, trailers); err != nil {
			return nil, err
		}
		types.OutgoingBodyFinish(body, cm.Some(trailers))
	}

	top := handleResp.OK()
	// wait until resp is returned
	subscription := top.Subscribe()
	subscription.Block()
	subscription.ResourceDrop()

	pollableOption := top.Get()
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

	result := resultOption.OK()
	trailers := http.Header{}

	respBody, err := NewIncomingBodyTrailer(result, trailers)
	if err != nil {
		return nil, fmt.Errorf("failed to consume incoming request %s", err)
	}

	resp := &http.Response{
		StatusCode: int(result.Status()),
		Body:       respBody,
		Trailer:    trailers,
	}

	return resp, nil
}
