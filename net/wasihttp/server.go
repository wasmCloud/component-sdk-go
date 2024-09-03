package wasihttp

import (
	"fmt"
	"net/http"
	"os"

	incominghandler "go.wasmcloud.dev/component/gen/wasi/http/incoming-handler"
	"go.wasmcloud.dev/component/gen/wasi/http/types"
)

// handler is the function that will be called by the http server.
var handler = defaultHandler

// defaultHandler is a placeholder for returning a useful error to stderr when
// the handler is not set.
var defaultHandler = func(http.ResponseWriter, *http.Request) {
	fmt.Fprintln(os.Stderr, "http handler undefined")
}

// Handle sets the handler function for the http trigger.
// It must be set in an init() function.
func Handle(fn func(http.ResponseWriter, *http.Request)) {
	handler = fn
}

func wasiHandle(request types.IncomingRequest, responseOut types.ResponseOutparam) {
	defer responseOut.ResourceDrop()

	httpReq, err := NewHttpRequest(request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to convert wasi/http/types.IncomingRequest to http.Request: %s\n", err)
		return
	}
	defer httpReq.Body.Close()

	httpRes := NewHttpResponseWriter(responseOut)
	defer httpRes.Close()

	handler(httpRes, httpReq)
}

func init() {
	incominghandler.Exports.Handle = wasiHandle
}
