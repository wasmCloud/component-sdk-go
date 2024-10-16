module github.com/wasmCloud/component-sdk-go/_examples/invoke

go 1.23.0

require (
	github.com/bytecodealliance/wasm-tools-go v0.2.1
	github.com/stretchr/testify v1.9.0
	go.wasmcloud.dev/component v0.0.0-20240910182305-2785f866ff0f
	go.wasmcloud.dev/wadge v0.6.0
)

require (
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/samber/lo v1.47.0 // indirect
	github.com/samber/slog-common v0.17.1 // indirect
	github.com/urfave/cli/v3 v3.0.0-alpha9 // indirect
	github.com/xrash/smetrics v0.0.0-20240521201337-686a1a2994c1 // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	golang.org/x/tools v0.25.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// TODO(lxf): remove this line once the module is published
replace go.wasmcloud.dev/component => ../..
