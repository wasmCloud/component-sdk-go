package wasihttp

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/lxfontes/component"
	incominghandler "github.com/lxfontes/component/gen/wasi/http/incoming-handler"
	"github.com/lxfontes/component/gen/wasi/http/types"
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
	// convert the incoming request to go's net/http type
	httpReq, err := NewHttpRequest(request)
	if err != nil {
		fmt.Printf("failed to convert wasi/http/types.IncomingRequest to http.Request: %s\n", err)
		return
	}

	// convert the response outparam to go's net/http type
	httpRes := NewHttpResponseWriter(responseOut)

	// run the user's handler
	handler(httpRes, httpReq)
}

func init() {
	incominghandler.Exports.Handle = wasiHandle
}
