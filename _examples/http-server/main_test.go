//go:generate go run github.com/rvolosatovs/west/cmd/west-bindgen-go

package main

import (
	"log"
	"log/slog"
	"os"
	"testing"
	"unsafe"

	"github.com/bytecodealliance/wasm-tools-go/cm"
	"github.com/rvolosatovs/west"
	_ "github.com/rvolosatovs/west/bindings"
	testtypes "github.com/rvolosatovs/west/bindings/wasi/http/types"
	teststreams "github.com/rvolosatovs/west/bindings/wasi/io/streams"
	httpext "github.com/rvolosatovs/west/bindings/wasiext/http/ext"
	"github.com/stretchr/testify/assert"
	incominghandler "go.wasmcloud.dev/component/gen/wasi/http/incoming-handler"
	"go.wasmcloud.dev/component/gen/wasi/http/types"
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
		headers := testtypes.NewFields()
		headers.Append(
			testtypes.FieldKey("foo"),
			testtypes.FieldValue(cm.NewList(
				unsafe.SliceData([]byte("bar")),
				3,
			)),
		)
		headers.Append(
			testtypes.FieldKey("foo"),
			testtypes.FieldValue(cm.NewList(
				unsafe.SliceData([]byte("baz")),
				3,
			)),
		)
		headers.Set(
			testtypes.FieldKey("key"),
			cm.NewList(
				unsafe.SliceData(
					[]testtypes.FieldValue{
						testtypes.FieldValue(cm.NewList(
							unsafe.SliceData([]byte("value")),
							5,
						)),
					},
				),
				1,
			),
		)
		req := testtypes.NewOutgoingRequest(headers)
		req.SetPathWithQuery(cm.Some("/test"))
		req.SetMethod(testtypes.MethodGet())
		out := httpext.NewResponseOutparam()
		incominghandler.Exports.Handle(
			types.IncomingRequest(httpext.NewIncomingRequest(req)),
			types.ResponseOutparam(out.F0),
		)
		out.F1.Subscribe().Block()
		respOptResRes := out.F1.Get()
		respResRes := respOptResRes.Some()
		if !assert.NotNil(t, respResRes) || !assert.Nil(t, respResRes.Err()) {
			t.FailNow()
		}
		respRes := respResRes.OK()
		if !assert.Nil(t, respRes.Err()) {
			t.Fatal(*respRes.Err())
		}
		resp := respRes.OK()
		assert.Equal(t, testtypes.StatusCode(200), resp.Status())

		hs := map[string][][]byte{}
		for _, h := range resp.Headers().Entries().Slice() {
			k := string(h.F0)
			hs[k] = append(hs[k], h.F1.Slice())
		}
		assert.Equal(t, map[string][][]byte{
			"content-type": {
				[]byte("text/plain"),
			},
			"x-requested-path": {
				[]byte("/test"),
			},
			"x-existing-paths": {
				[]byte("/error,/form,/headers,/post"),
			},
		}, hs)
		bodyRes := resp.Consume()
		if !assert.Nil(t, bodyRes.Err()) {
			t.FailNow()
		}

		body := bodyRes.OK()
		bodyStreamRes := body.Stream()
		if !assert.Nil(t, bodyStreamRes.Err()) {
			t.FailNow()
		}

		bodyStream := bodyStreamRes.OK()
		var buf []byte
		for {
			bufRes := bodyStream.BlockingRead(4096)
			if bufRes.IsErr() && *bufRes.Err() == teststreams.StreamErrorClosed() {
				break
			}
			if !assert.Nil(t, bufRes.Err()) {
				if !assert.False(t, bufRes.Err().Closed()) {
					t.FailNow()
				} else {
					t.Fatal(*bufRes.Err().LastOperationFailed())
				}
			}
			buf = append(buf, bufRes.OK().Slice()...)
		}

		assert.Equal(t, []byte(Index), buf)
		bodyStream.ResourceDrop()
	})
}
