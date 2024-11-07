// Original source: https://github.com/open-telemetry/opentelemetry-proto-go/blob/v1.3.1/slim/otlp/collector/trace/v1/trace_service.pb.go
package types

type ExportTraceServiceRequest struct {
	ResourceSpans []*ResourceSpans `json:"resource_spans,omitempty"`
}
