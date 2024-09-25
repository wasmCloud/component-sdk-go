package wasihttp

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bytecodealliance/wasm-tools-go/cm"
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
func Handle(h http.Handler) {
	handler = h.ServeHTTP
}

// HandleFunc sets the [net/http.HandlerFunc] that will be called to handle the
// incoming request.
func HandleFunc(h http.HandlerFunc) {
	handler = h
}

func wasiHandle(request types.IncomingRequest, responseOut types.ResponseOutparam) {
	httpReq, err := WASItoHTTPRequest(request)
	if err != nil {
		types.ResponseOutparamSet(responseOut, cm.Err[cm.Result[types.ErrorCodeShape, types.OutgoingResponse, types.ErrorCode]](
			types.ErrorCodeInternalError(cm.Some(err.Error()))),
		)
		return
	}
	if httpReq.Body != nil {
		defer httpReq.Body.Close()
	}

	httpRes := WASItoHTTPResponseWriter(responseOut)
	defer httpRes.Close()

	handler(httpRes, httpReq)
}

func init() {
	incominghandler.Exports.Handle = wasiHandle
}
