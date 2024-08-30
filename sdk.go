package component

//go:generate wit-bindgen-go generate --world imports --out gen ./wit

import (
	"embed"
	"log/slog"

	"go.wasmcloud.dev/component/log/wasilog"
	"go.wasmcloud.dev/component/net"
)

//go:embed wit/*
var Wit embed.FS

var DefaultLogger = slog.New(wasilog.DefaultOptions().NewHandler())

// exposed as public variable so components can mock calls
var EnableSockets = net.EnableSockets
