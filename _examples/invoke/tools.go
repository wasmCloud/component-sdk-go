//go:build tools

package main

import (
	_ "go.wasmcloud.dev/component/codegen"
	_ "go.wasmcloud.dev/wadge/cmd/wadge-bindgen-go"
)
