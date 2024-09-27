//go:generate go run github.com/wasmCloud/wadge/cmd/wadge-bindgen-go

package main

import (
	"log"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wasmCloud/wadge"
	_ "github.com/wasmCloud/wadge/bindings"
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
