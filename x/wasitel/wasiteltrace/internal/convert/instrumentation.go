// Original source: https://github.com/open-telemetry/opentelemetry-go/blob/v1.31.0/exporters/otlp/otlptrace/internal/tracetransform/instrumentation.go
package convert

import (
	"go.opentelemetry.io/otel/sdk/instrumentation"

	"go.wasmcloud.dev/component/x/wasitel/wasiteltrace/internal/types"
)

func convertInstrumentationScope(il instrumentation.Scope) *types.InstrumentationScope {
	if il == (instrumentation.Scope{}) {
		return nil
	}
	return &types.InstrumentationScope{
		Name:    il.Name,
		Version: il.Version,
	}
}
