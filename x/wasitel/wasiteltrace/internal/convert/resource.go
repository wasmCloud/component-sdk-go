// Original source: https://github.com/open-telemetry/opentelemetry-go/blob/v1.31.0/exporters/otlp/otlptrace/internal/tracetransform/resource.go
package convert

import (
	"go.opentelemetry.io/otel/sdk/resource"

	"go.wasmcloud.dev/component/x/wasitel/wasiteltrace/internal/types"
)

// Resource transforms a Resource into an OTLP Resource.
func Resource(r *resource.Resource) *types.Resource {
	if r == nil {
		return nil
	}
	return &types.Resource{Attributes: ResourceAttributes(r)}
}
