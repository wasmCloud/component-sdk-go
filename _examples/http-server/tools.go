//go:build tools

package main

import (
	_ "github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go"
	_ "github.com/wasmCloud/west/cmd/west-bindgen-go"
)
