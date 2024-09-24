module github.com/wasmCloud/component-sdk-go/_examples/http-client

go 1.23.0

require (
	github.com/bytecodealliance/wasm-tools-go v0.2.0
	go.wasmcloud.dev/component v0.0.0-20240910182305-2785f866ff0f
)

// TODO(lxf): remove this line once the module is published
replace go.wasmcloud.dev/component => ../..
