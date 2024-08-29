package component

//go:generate wit-bindgen-go generate --world imports --out gen ./wit

import (
	"embed"

	// register wasm imports / exports
	_ "go.wasmcloud.dev/component/gen/wasi/http/incoming-handler"
)

//go:embed wit/*
var Wit embed.FS

func Init() {
	// blank function. exercise the _ imports
}
