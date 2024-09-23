package component

//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world sdk --out gen ./wit

import (
	"embed"
)

//go:embed wit/*
var Wit embed.FS
