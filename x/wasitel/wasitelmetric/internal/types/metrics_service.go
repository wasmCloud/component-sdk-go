// Original source: https://github.com/open-telemetry/opentelemetry-proto-go/blob/v1.3.1/slim/otlp/collector/metrics/v1/metrics_service.pb.go
package types

type ExportMetricsServiceRequest struct {
	// An array of ResourceMetrics.
	// For data coming from a single resource this array will typically contain one
	// element. Intermediary nodes (such as OpenTelemetry Collector) that receive
	// data from multiple origins typically batch the data before forwarding further and
	// in that case this array will contain multiple elements.
	ResourceMetrics []*ResourceMetrics `json:"resource_metrics,omitempty"`
}
