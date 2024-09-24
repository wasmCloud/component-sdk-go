//go:generate go run github.com/wasmCloud/west/cmd/west-bindgen-go

package main

import (
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wasmCloud/west"
	_ "github.com/wasmCloud/west/bindings"
	"github.com/wasmCloud/west/westhttp"
	incominghandler "go.wasmcloud.dev/component/gen/wasi/http/incoming-handler"
)

func init() {
	log.SetFlags(0)
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug, ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	})))
}

func TestIncomingHandler(t *testing.T) {
	west.RunTest(t, func() {
		req, err := http.NewRequest("", "/test", nil)
		if err != nil {
			t.Fatalf("failed to create new HTTP request: %s", err)
		}
		req.Header.Add("foo", "bar")
		req.Header.Add("foo", "baz")
		req.Header.Add("key", "value")
		resp, err := westhttp.HandleIncomingRequest(incominghandler.Exports.Handle, req)
		if err != nil {
			t.Fatalf("failed to handle incoming HTTP request: %s", err)
		}
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, http.Header{
			"content-type": {
				"text/plain",
			},
			"x-requested-path": {
				"/test",
			},
			"x-existing-paths": {
				"/error,/form,/headers,/post",
			},
		}, resp.Header)
		buf, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read HTTP response body: %s", err)
		}
		assert.Equal(t, []byte(Index), buf)
	})
}
