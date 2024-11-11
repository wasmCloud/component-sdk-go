package wasiteltrace

import (
	"context"
	"fmt"
	"sync"

	"go.opentelemetry.io/otel/sdk/trace"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.wasmcloud.dev/component/x/wasitel/wasiteltrace/internal/convert"
)

func New(opts ...Option) (*Exporter, error) {
	client := newClient(opts...)
	return &Exporter{
		client: client,
	}, nil
}

var _ tracesdk.SpanExporter = (*Exporter)(nil)

type Exporter struct {
	client    *client
	stopped   bool
	stoppedMu sync.RWMutex
}

func (e *Exporter) ExportSpans(ctx context.Context, spans []trace.ReadOnlySpan) error {
	err := ctx.Err()
	if err != nil {
		return err
	}

	// Check whether the exporter has been told to Shutdown
	e.stoppedMu.RLock()
	stopped := e.stopped
	e.stoppedMu.RUnlock()
	if stopped {
		return nil
	}

	// Check if there's nothing to export
	converted := convert.ResourceSpans(spans)
	if len(converted) == 0 {
		return nil
	}

	err = e.client.UploadTraces(ctx, converted)
	if err != nil {
		return fmt.Errorf("failed to export spans: %w", err)
	}
	return nil
}

// Shutdown is called to stop the exporter, it performs no action.
func (e *Exporter) Shutdown(ctx context.Context) error {
	e.stoppedMu.Lock()
	e.stopped = true
	e.stoppedMu.Unlock()

	return nil
}

// MarshalLog is the marshaling function used by the logging system to represent this Exporter.
func (e *Exporter) MarshalLog() interface{} {
	return struct {
		Type string
	}{
		Type: "wasiteltrace",
	}
}
