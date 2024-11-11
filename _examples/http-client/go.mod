module github.com/wasmCloud/component-sdk-go/_examples/http-client

go 1.23.0

require (
	go.bytecodealliance.org v0.4.0
	go.wasmcloud.dev/component v0.0.5
)

require (
	github.com/coreos/go-semver v0.3.1 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/regclient/regclient v0.7.2 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/ulikunitz/xz v0.5.12 // indirect
	github.com/urfave/cli/v3 v3.0.0-alpha9.2 // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
)

// NOTE(lxf): Remove this line if running outside of component-sdk-go repository
replace go.wasmcloud.dev/component => ../..
