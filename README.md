# Go Component SDK

[![Go Reference](https://pkg.go.dev/badge/go.wasmcloud.dev/component.svg)](https://pkg.go.dev/go.wasmcloud.dev/component)

The Go Component SDK provides a set of packages to simplify the development of WebAssembly components targeting the [wasmCloud](https://wasmcloud.com) host runtime.

Writing a wasmCloud Capability Provider? Check out the [Go Provider SDK](https://github.com/wasmCloud/provider-sdk-go).

# Setup

Requires tinygo 0.33 or above.

Import `go.wasmcloud.dev/component` in your Go module.

```bash
go get go.wasmcloud.dev/component@v0.0.5
```

Import the SDK WIT. In `wit/deps.toml`:

```toml

wasmcloud-component = "https://github.com/wasmCloud/component-sdk-go/archive/v0.0.5.tar.gz"

```

Run `wit-deps` to update your wit dependencies.

And in your world definition:

```

include wasmcloud:component-go/imports@0.1.0;

```

# Adapters

## net/wasihttp

The `wasihttp` package provides an implementation of `http.Handler` backed by `wasi:http`, as well as a `http.RoundTripper` backed by `wasi:http`.

### http.Handler

`wasihttp.Handle` registers an `http.Handler` to be served at a given path, converting `wasi:http` requests/responses into standard `http.Request` and `http.ResponseWriter` objects.

```go
package main

import (
	"net/http"
	"go.wasmcloud.dev/component/net/wasihttp"
)

func httpServe(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func init() {
	// request will be fulfilled via wasi:http/incoming-handler
	wasihttp.HandleFunc(httpServe)
}
```

### http.RoundTripper

```go
package main

import (
	"net/http"
	"go.wasmcloud.dev/component/net/wasihttp"
)

var (
	wasiTransport = &wasihttp.Transport{}
	httpClient    = &http.Client{Transport: wasiTransport}
)

func httpClient() {
	// request will be fulfilled via wasi:http/outgoing-handler
	httpClient.Get("http://example.com")
}
```

## log/wasilog

The `wasilog` package provides an implementation of `slog.Handler` backed by `wasi:logging`.

Sample usage:

```go
package main

import (
	"log/slog"
	"go.wasmcloud.dev/component/log/wasilog"
)

func wasilog() {
	logger := slog.New(wasilog.DefaultOptions().NewHandler())

	logger.Info("Hello")
	logger.Info("Hello", "planet", "Earth")
	logger.Info("Hello", slog.String("planet", "Earth"))
}
```

See `wasilog.Options` for log level & other configuration options.

## Community

Similar projects:

- [rajatjindal/wasi-go-sdk](https://github.com/rajatjindal/wasi-go-sdk)
- [dev-wasm/dev-wasm-go](https://github.com/dev-wasm/dev-wasm-go)
- [Mossaka/hello-wasi-http-tinygo](https://github.com/Mossaka/hello-wasi-http-tinygo)
