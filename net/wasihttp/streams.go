package wasihttp

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/bytecodealliance/wasm-tools-go/cm"
	"go.wasmcloud.dev/component/gen/wasi/http/types"
	"go.wasmcloud.dev/component/gen/wasi/io/streams"
)

type BodyConsumer interface {
	Consume() (result cm.Result[types.IncomingBody, types.IncomingBody, struct{}])
	Headers() (result types.Fields)
}
type inputStreamReader struct {
	consumer    BodyConsumer
	body        types.IncomingBody
	stream      streams.InputStream
	trailerLock sync.Mutex
	trailers    http.Header
	trailerOnce sync.Once
	dropOnce    sync.Once
}

func (r *inputStreamReader) Close() error {
	r.dropOnce.Do(func() {
		r.stream.ResourceDrop()
	})

	return nil
}

func (r *inputStreamReader) parseTrailers() {
	r.trailerLock.Lock()
	defer r.trailerLock.Unlock()

	// if we got this far, then we release ownership from body, otherwise it is our responsibility to drop it
	r.dropOnce.Do(func() {})

	r.stream.ResourceDrop()
	futureTrailers := types.IncomingBodyFinish(r.body)

	trailersResult := futureTrailers.Get()

	// unroll the future
	if trailersResult.None() {
		return
	}
	if trailersResult.Some().IsErr() {
		return
	}
	if trailersResult.Some().OK().IsErr() {
		return
	}
	maybeWasiTrailers := trailersResult.Some().OK().OK()

	if maybeWasiTrailers.None() {
		return
	}

	wasiTrailers := maybeWasiTrailers.Some()
	for _, kv := range wasiTrailers.Entries().Slice() {
		r.trailers.Add(string(kv.F0), string(kv.F1.Slice()))
	}

	wasiTrailers.ResourceDrop()
}

func (r *inputStreamReader) Read(p []byte) (n int, err error) {
	readResult := r.stream.BlockingRead(uint64(len(p)))
	if readResult.IsErr() {
		readErr := readResult.Err()
		if readErr.Closed() {
			r.trailerOnce.Do(r.parseTrailers)
			return 0, io.EOF
		}
		return 0, fmt.Errorf("failed to read from InputStream %s", readErr.LastOperationFailed().ToDebugString())
	}

	readList := *readResult.OK()
	copy(p, readList.Slice())
	return int(readList.Len()), nil
}

func NewIncomingBodyTrailer(consumer BodyConsumer, trailers http.Header) (io.ReadCloser, error) {
	consumeResult := consumer.Consume()
	if consumeResult.IsErr() {
		return nil, fmt.Errorf("failed to consume incoming request %s", *consumeResult.Err())
	}
	body := consumeResult.OK()
	streamResult := body.Stream()
	if streamResult.IsErr() {
		return nil, fmt.Errorf("failed to consume incoming requests's stream %s", streamResult.Err())
	}
	return &inputStreamReader{
		consumer: consumer,
		trailers: trailers,
		body:     *body,
		stream:   *streamResult.OK(),
	}, nil
}

type outputStreamReader struct {
	body   types.OutgoingBody
	stream streams.OutputStream
}

func NewOutgoingBody(body types.OutgoingBody) (io.WriteCloser, error) {
	stream := body.Write()
	if stream.IsErr() {
		return nil, fmt.Errorf("failed to acquire resource handle to request body: %s", stream.Err())
	}
	return &outputStreamReader{
		body:   body,
		stream: *stream.OK(),
	}, nil
}

func (r *outputStreamReader) Close() error {
	r.stream.ResourceDrop()
	r.body.ResourceDrop()
	return nil
}

func (r *outputStreamReader) Write(p []byte) (n int, err error) {
	contents := cm.ToList(p)
	writeResult := r.stream.BlockingWriteAndFlush(contents)
	if writeResult.IsErr() {
		if writeResult.Err().Closed() {
			return 0, io.EOF
		}

		return 0, fmt.Errorf("failed to write to response body's stream: %s", writeResult.Err().LastOperationFailed().ToDebugString())
	}
	return len(p), nil
}
