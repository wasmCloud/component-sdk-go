package wasihttp

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"go.wasmcloud.dev/component/gen/wasi/http/types"
	"go.wasmcloud.dev/component/gen/wasi/io/streams"
	"github.com/ydnar/wasm-tools-go/cm"
)

var _ http.ResponseWriter = &responseOutparamWriter{}

type responseOutparamWriter struct {
	// wasi response outparam is set at the end of http_trigger_handle
	outparam types.ResponseOutparam
	// wasi response
	response types.OutgoingResponse
	// wasi http headers
	wasiHeaders types.Fields
	// go httpHeaders are reconciled on call to WriteHeader, Flush or at the end of http_trigger_handle
	httpHeaders http.Header
	// wasi response body is set on first write because it can only be called once
	body *types.OutgoingBody
	// wasi response stream is set on first write because it can only be called once
	stream *streams.OutputStream

	statuscode int
}

func (row *responseOutparamWriter) Header() http.Header {
	return row.httpHeaders
}

func (row *responseOutparamWriter) Write(buf []byte) (int, error) {
	// acquire the response body's resource handle on first call to write
	if row.body == nil {
		bodyResult := row.response.Body()
		if bodyResult.IsErr() {
			return 0, fmt.Errorf("failed to acquire resource handle to response body: %s", bodyResult.Err())
		}
		row.body = bodyResult.OK()

		writeResult := row.body.Write()
		if writeResult.IsErr() {
			return 0, fmt.Errorf("failed to acquire resource handle for response body's stream: %s", writeResult.Err())
		}
		row.stream = writeResult.OK()
	}

	// //TODO: determine if we need to do these to fulfill the ResponseWriter contract
	// // call WriteHeader(http.StatusOK) if it hasn't been called yet
	// // call DetectContentType if headers doesn't contain content-type yet
	// // if total data is under "a few" KB and there are no flush calls, Content-Length is added automatically

	contents := cm.ToList(buf)
	writeResult := row.stream.Write(contents)
	if writeResult.IsErr() {
		if writeResult.Err().Closed() {
			return 0, fmt.Errorf("failed to write to response body's stream: closed")
		}

		// TODO: possible nil error here
		return 0, fmt.Errorf("failed to write to response body's stream: %s", writeResult.Err().LastOperationFailed().ToDebugString())
	}

	result := cm.OK[cm.Result[types.ErrorCodeShape, types.OutgoingResponse, types.ErrorCode]](row.response)
	types.ResponseOutparamSet(row.outparam, result)

	return int(contents.Len()), nil
}

func (row *responseOutparamWriter) WriteHeader(statusCode int) {
	row.statuscode = statusCode
	row.reconcile()
}

// reconcile headers from go to wasi
func (row *responseOutparamWriter) reconcileHeaders() error {
	for key, vals := range row.httpHeaders {
		// convert each value distincly
		fieldVals := []types.FieldValue{}
		for _, val := range vals {
			fieldVals = append(fieldVals, types.FieldValue(cm.ToList([]uint8(val))))
		}

		if result := row.wasiHeaders.Set(types.FieldKey(key), cm.ToList(fieldVals)); result.IsErr() {
			switch *result.Err() {
			case types.HeaderErrorInvalidSyntax:
				return fmt.Errorf("failed to set header %s to [%s]: invalid syntax", key, strings.Join(vals, ","))
			case types.HeaderErrorForbidden:
				return fmt.Errorf("failed to set forbidden header key %s", key)
			case types.HeaderErrorImmutable:
				return fmt.Errorf("failed to set header on immutable header fields")
			default:
				return fmt.Errorf("not sure what happened here?")
			}
		}
	}

	// TODO: handle deleted headers

	return nil
}

// convert the ResponseOutparam to http.ResponseWriter
func NewHttpResponseWriter(out types.ResponseOutparam) *responseOutparamWriter {
	row := &responseOutparamWriter{
		outparam:    out,
		httpHeaders: http.Header{},
		wasiHeaders: types.NewFields(),
	}

	return row
}

func (row *responseOutparamWriter) reconcile() {
	err := row.reconcileHeaders()
	if err != nil {
		// TODO
	}

	// setting headers after this cause panic
	// TODO: debug
	row.response = types.NewOutgoingResponse(row.wasiHeaders)

	// set status code
	row.response.SetStatusCode(types.StatusCode(row.statuscode))
}

func println(msg string) {
	fmt.Println(cm.ToList([]byte(msg)))
	fmt.Println(cm.ToList([]byte("\n")))
}

type IncomingRequest = types.IncomingRequest

// convert the IncomingRequest to http.Request
func NewHttpRequest(ir IncomingRequest) (req *http.Request, err error) {
	// convert the http method to string
	method, err := methodToString(ir.Method())
	if err != nil {
		return nil, err
	}

	// convert the path with query to a url
	var url string
	if pathWithQuery := ir.PathWithQuery(); pathWithQuery.None() {
		url = ""
	} else {
		url = *pathWithQuery.Some()
	}

	// convert the body to a reader
	var body io.Reader
	if consumeResult := ir.Consume(); consumeResult.IsErr() {
		return nil, fmt.Errorf("failed to consume incoming request %s", *consumeResult.Err())
	} else if streamResult := consumeResult.OK().Stream(); streamResult.IsErr() {
		return nil, fmt.Errorf("failed to consume incoming requests's stream %s", streamResult.Err())
	} else {
		body = NewReader(*streamResult.OK())
	}

	// create a new request
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// update additional fields
	toHttpHeader(ir.Headers(), &req.Header)

	return req, nil
}

func methodToString(m types.Method) (string, error) {
	if m.Connect() {
		return "CONNECT", nil
	}
	if m.Delete() {
		return "DELETE", nil
	}
	if m.Get() {
		return "GET", nil
	}
	if m.Head() {
		return "HEAD", nil
	}
	if m.Options() {
		return "OPTIONS", nil
	}
	if m.Patch() {
		return "PATCH", nil
	}
	if m.Post() {
		return "POST", nil
	}
	if m.Put() {
		return "PUT", nil
	}
	if m.Trace() {
		return "TRACE", nil
	}
	if other := m.Other(); other != nil {
		return *other, fmt.Errorf("unknown http method 'other'")
	}
	return "", fmt.Errorf("failed to convert http method")
}

func toHttpHeader(src types.Fields, dest *http.Header) {
	for _, f := range src.Entries().Slice() {
		key := string(f.F0)
		value := string(cm.List[uint8](f.F1).Slice())
		dest.Add(key, value)
	}
}

// convert the IncomingRequest to http.Request
func NewOutgoingHttpRequest(req *http.Request) (types.OutgoingRequest, error) {
	headers := types.NewFields()
	toWasiHeader(req.Header, headers)

	or := types.NewOutgoingRequest(headers)
	or.SetAuthority(cm.Some(req.Host))
	or.SetMethod(toWasiMethod(req.Method))
	or.SetPathWithQuery(cm.Some(req.URL.Path + "?" + req.URL.Query().Encode()))

	switch req.URL.Scheme {
	case "http":
		or.SetScheme(cm.Some(types.SchemeHTTP()))
	case "https":
		or.SetScheme(cm.Some(types.SchemeHTTPS()))
	default:
		or.SetScheme(cm.Some(types.SchemeOther(req.URL.Scheme)))
	}

	return or, nil
}

func toWasiHeader(src http.Header, dest types.Fields) {
	for k, v := range src {
		key := types.FieldKey(k)
		fieldVals := []types.FieldValue{}

		for _, val := range v {
			fieldVals = append(fieldVals, types.FieldValue(cm.ToList([]uint8(val))))
		}

		// TODO(rjindal): check error
		_ = dest.Set(key, cm.ToList(fieldVals))
	}
}

func toWasiMethod(s string) types.Method {
	switch s {
	case http.MethodConnect:
		return types.MethodConnect()
	case http.MethodDelete:
		return types.MethodDelete()
	case http.MethodGet:
		return types.MethodGet()
	case http.MethodHead:
		return types.MethodHead()
	case http.MethodOptions:
		return types.MethodOptions()
	case http.MethodPatch:
		return types.MethodPatch()
	case http.MethodPost:
		return types.MethodPost()
	case http.MethodPut:
		return types.MethodPut()
	case http.MethodTrace:
		return types.MethodTrace()
	default:
		return types.MethodOther(s)
	}
}
