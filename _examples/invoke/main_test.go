//go:generate go run go.wasmcloud.dev/wadge/cmd/wadge-bindgen-go -test

package main

import (
	"log"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.wasmcloud.dev/wadge"
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

func TestCall(t *testing.T) {
	wadge.RunTest(t, func() {
		buf := invokerCall()
		assert.Equal(t, InvokeResponse, buf)
	})
}
