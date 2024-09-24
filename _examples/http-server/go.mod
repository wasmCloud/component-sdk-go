module github.com/wasmCloud/component-sdk-go/_examples/http-server

go 1.23.0

require (
	github.com/bytecodealliance/wasm-tools-go v0.2.0
	github.com/stretchr/testify v1.9.0
	github.com/wasmCloud/west v0.2.0
	go.wasmcloud.dev/component v0.0.0-20240910182305-2785f866ff0f
)

require (
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/samber/lo v1.44.0 // indirect
	github.com/samber/slog-common v0.17.1 // indirect
	github.com/urfave/cli/v3 v3.0.0-alpha9 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/tools v0.24.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// TODO(lxf): remove this line once the module is published
replace go.wasmcloud.dev/component => ../..
