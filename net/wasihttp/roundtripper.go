package wasihttp

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ydnar/wasm-tools-go/cm"
	outgoinghandler "go.wasmcloud.dev/component/gen/wasi/http/outgoing-handler"
	"go.wasmcloud.dev/component/gen/wasi/http/types"
)

// Transport implements http.RoundTripper
type Transport struct{}

var _ http.RoundTripper = &Transport{}

var DefaultClient = NewClient()

// NewClient returns a new HTTP client compatible with wasi preview 2
func NewClient() *http.Client {
	return &http.Client{
		Transport: &Transport{},
	}
}

func (r *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	or, err := NewOutgoingHttpRequest(req)
	if err != nil {
		return nil, err
	}

	result := outgoinghandler.Handle(or, cm.None[types.RequestOptions]())
	if result.Err() != nil {
		return nil, fmt.Errorf("TODO: convert to readable error")
	}

	if result.IsErr() {
		return nil, fmt.Errorf("error is %v", result.Err())
	}

	okresult := result.OK()

	// wait until resp is returned
	okresult.Subscribe().Block()

	incomingResp := okresult.Get()
	if incomingResp.None() {
		return nil, fmt.Errorf("incoming resp is None")
	}

	if incomingResp.Some().IsErr() {
		return nil, fmt.Errorf("error is %v", incomingResp.Some().Err())
	}

	if incomingResp.Some().OK().IsErr() {
		return nil, fmt.Errorf("error is %v", incomingResp.Some().OK().Err())
	}

	okresp := incomingResp.Some().OK().OK()
	var body io.ReadCloser
	if consumeResult := okresp.Consume(); consumeResult.IsErr() {
		return nil, fmt.Errorf("failed to consume incoming request %s", *consumeResult.Err())
	} else if streamResult := consumeResult.OK().Stream(); streamResult.IsErr() {
		return nil, fmt.Errorf("failed to consume incoming requests's stream %s", streamResult.Err())
	} else {
		body = NewReadCloser(*streamResult.OK())
	}

	resp := &http.Response{
		StatusCode: int(okresp.Status()),
		Body:       body,
	}

	return resp, nil
}
