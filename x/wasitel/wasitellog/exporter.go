package wasitellog

import (
	"context"
	"fmt"
	"sync"

	logsdk "go.opentelemetry.io/otel/sdk/log"
	"go.wasmcloud.dev/component/x/wasitel/wasitellog/internal/convert"
)

func New(opts ...Option) (*Exporter, error) {
	client := newClient(opts...)
	return &Exporter{
		client: client,
	}, nil
}

var _ logsdk.Exporter = (*Exporter)(nil)

type Exporter struct {
	client    *client
	stopped   bool
	stoppedMu sync.RWMutex
}

func (e *Exporter) Export(ctx context.Context, logs []logsdk.Record) error {
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

	// Check whether there's anything to export
	converted := convert.ResourceLogs(logs)
	if len(converted) == 0 {
		return nil
	}

	err = e.client.UploadLogs(ctx, converted)
	if err != nil {
		return fmt.Errorf("failed to export spans: %w", err)
	}
	return nil
}

func (e *Exporter) Shutdown(ctx context.Context) error {
	e.stoppedMu.Lock()
	e.stopped = true
	e.stoppedMu.Unlock()

	return nil
}

func (e *Exporter) ForceFlush(ctx context.Context) error {
	return nil
}
