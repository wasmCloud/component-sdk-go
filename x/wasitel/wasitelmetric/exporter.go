package wasitelmetric

import (
	"context"
	"fmt"
	"sync"

	metric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.wasmcloud.dev/component/x/wasitel/wasitelmetric/internal/convert"
)

func New(opts ...Option) (*Exporter, error) {
	client := newClient(opts...)
	return &Exporter{
		client: client,
	}, nil
}

type Exporter struct {
	client *client

	temporalitySelector metric.TemporalitySelector
	aggregationSelector metric.AggregationSelector

	stopped   bool
	stoppedMu sync.RWMutex
}

var _ metric.Exporter = (*Exporter)(nil)

// Temporality returns the Temporality to use for an instrument kind.
func (e *Exporter) Temporality(k metric.InstrumentKind) metricdata.Temporality {
	return e.temporalitySelector(k)
}

// Aggregation returns the Aggregation to use for an instrument kind.
func (e *Exporter) Aggregation(k metric.InstrumentKind) metric.Aggregation {
	return e.aggregationSelector(k)
}

func (e *Exporter) Export(ctx context.Context, data *metricdata.ResourceMetrics) error {
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

	metrics, err := convert.ResourceMetrics(data)
	if err != nil {
		return err
	}

	err = e.client.UploadMetrics(ctx, metrics)
	if err != nil {
		return err
	}

	return nil
}

// ForceFlush flushes any metric data held by an exporter.
//
// This method returns an error if called after Shutdown.
// This method returns an error if the method is canceled by the passed context.
//
// This method is safe to call concurrently.
func (e *Exporter) ForceFlush(ctx context.Context) error {
	// Check whether the exporter has been told to Shutdown
	e.stoppedMu.RLock()
	stopped := e.stopped
	e.stoppedMu.RUnlock()
	if stopped {
		return errShutdown
	}
	return ctx.Err()
}

var errShutdown = fmt.Errorf("Exporter is shutdown")

// Shutdown flushes all metric data held by an exporter and releases any held
// computational resources.
//
// This method returns an error if called after Shutdown.
// This method returns an error if the method is canceled by the passed context.
//
// This method is safe to call concurrently.
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
		Type: "wasitelmetric",
	}
}
