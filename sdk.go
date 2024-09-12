package component

//go:generate wit-bindgen-go generate --world imports --out gen ./wit

import (
	"embed"
	"log/slog"

	"go.wasmcloud.dev/component/log/wasilog"
	"go.wasmcloud.dev/component/net"

	// TODO(lxf): Investigate if it is better to remove this import and let callers go directly to wasihttp
	// Would allow callers to avoid importing wasmcloud:component world.
	_ "go.wasmcloud.dev/component/net/wasihttp"
)

//go:embed wit/*
var Wit embed.FS

var DefaultLogger = slog.New(wasilog.DefaultOptions().NewHandler())

func ContextLogger(wasiContext string) *slog.Logger {
	return DefaultLogger.With(wasilog.ContextAttr(wasiContext))
}

// exposed as public variable so components can mock calls
var EnableSockets = net.EnableSockets
