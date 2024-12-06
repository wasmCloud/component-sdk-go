package convert

import (
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.wasmcloud.dev/component/x/wasitel/wasitelmetric/internal/types"
)

func ResourceMetrics(data *metricdata.ResourceMetrics) (*types.ResourceMetrics, error) {
	return nil, nil
}
