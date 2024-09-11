# Go Component SDK

Warning: API stabilization in progress. Once ready it will be published as `v0.0.1`.

The Go Component SDK provides a set of packages to simplify the development of WebAssembly components targeting the [wasmCloud](https://wasmcloud.com) host runtime.

Writing a wasmCloud Capability Provider? Check out the [Go Provider SDK](https://github.com/wasmCloud/provider-sdk-go).

# Setup

Requires tinygo 0.33 or above.

Import `go.wasmcloud.dev/component` in your Go module.

Import the SDK WIT. In `wit/deps.toml`:

```
wasmcloud-component = "https://github.com/wasmCloud/component-sdk-go/archive/<TBD>.tar.gz"
```

Run `wit-deps` to update your wit dependencies.

And in your world definition:

```
  include wasmcloud:component/imports;
```

## net/wasihttp

The `wasihttp` package provides an implementation of `http.Handler` backed by `wasi:http`, as well as a `http.RoundTripper` backed by `wasi:http`.

### http.Handler

`wasihttp.Handle` registers an `http.Handler` to be served at a given path, converting `wasi:http` requests/responses into standard `http.Request` and `http.ResponseWriter` objects.

```go
import (
  "net/http"
  "go.wasmcloud.dev/component/net/wasihttp"
)

func httpServe(w http.ResponseWriter, *r http.Request) {
  w.Write([]byte("Hello, world!"))
}

func init() {
// request will be fulfilled via wasi:http/incoming-handler
  wasihttp.Handle("/", httpServe)
}
```

### http.RoundTripper

```go
import (
  "net/http"
  "go.wasmcloud.dev/component/net/wasihttp"
)

var wasiTransport = &wasihttp.Transport{}
var httpClient = &http.Client{Transport: wasiTransport}

// request will be fulfilled via wasi:http/outgoing-handler
httpClient.Get("http://example.com")
```

## log/wasilog

The `wasilog` package provides an implementation of `slog.Handler` backed by `wasi:logging`.

Sample usage:

```go
import (
  "log/slog"
  "go.wasmcloud.dev/component/log/wasilog"
)

logger := slog.New(wasilog.DefaultOptions().NewHandler())

logger.Info("Hello")
logger.Info("Hello", "planet", "Earth")
logger.Info("Hello", slog.String("planet", "Earth"))
```

See `wasilog.Options` for log level & other configuration options.
