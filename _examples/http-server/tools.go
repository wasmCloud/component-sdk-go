//go:build tools

package main

import (
	_ "go.wasmcloud.dev/component/wit-bindgen-go"
	_ "go.wasmcloud.dev/wadge/cmd/wadge-bindgen-go"
)
