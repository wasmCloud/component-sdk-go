module go.wasmcloud.dev/component/x/wasitel

go 1.23.2

require (
	go.opentelemetry.io/otel v1.31.0
	go.opentelemetry.io/otel/sdk v1.31.0
	go.opentelemetry.io/otel/trace v1.31.0
	go.wasmcloud.dev/component v0.0.4
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	go.bytecodealliance.org v0.4.0 // indirect
	go.opentelemetry.io/otel/metric v1.31.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
)

replace go.wasmcloud.dev/component => ../../
