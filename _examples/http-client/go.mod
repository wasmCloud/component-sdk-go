module github.com/wasmCloud/component-sdk-go/_examples/http-server

go 1.22.5

require (
	github.com/bytecodealliance/wasm-tools-go v0.2.0
	go.wasmcloud.dev/component v0.0.0-20240910182305-2785f866ff0f
)

require (
	github.com/samber/lo v1.44.0 // indirect
	github.com/samber/slog-common v0.17.1 // indirect
	golang.org/x/text v0.16.0 // indirect
)

// TODO(lxf): remove this line once the module is published
replace go.wasmcloud.dev/component => ../..
