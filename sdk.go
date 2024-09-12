package component

//go:generate wit-bindgen-go generate --world sdk --out gen ./wit

import (
	"embed"
	"log/slog"

	"go.wasmcloud.dev/component/log/wasilog"
	"go.wasmcloud.dev/component/net"
)

//go:embed wit/*
var Wit embed.FS

var DefaultLogger = slog.New(wasilog.DefaultOptions().NewHandler())

func ContextLogger(wasiContext string) *slog.Logger {
	return DefaultLogger.With(wasilog.ContextAttr(wasiContext))
}

// exposed as public variable so components can mock calls
var EnableSockets = net.EnableSockets
