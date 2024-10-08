// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package environment represents the imported interface "wasi:cli/environment@0.2.0".
package environment

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// GetEnvironment represents the imported function "get-environment".
//
//	get-environment: func() -> list<tuple<string, string>>
//
//go:nosplit
func GetEnvironment() (result cm.List[[2]string]) {
	wasmimport_GetEnvironment(&result)
	return
}

//go:wasmimport wasi:cli/environment@0.2.0 get-environment
//go:noescape
func wasmimport_GetEnvironment(result *cm.List[[2]string])

// GetArguments represents the imported function "get-arguments".
//
//	get-arguments: func() -> list<string>
//
//go:nosplit
func GetArguments() (result cm.List[string]) {
	wasmimport_GetArguments(&result)
	return
}

//go:wasmimport wasi:cli/environment@0.2.0 get-arguments
//go:noescape
func wasmimport_GetArguments(result *cm.List[string])

// InitialCWD represents the imported function "initial-cwd".
//
//	initial-cwd: func() -> option<string>
//
//go:nosplit
func InitialCWD() (result cm.Option[string]) {
	wasmimport_InitialCWD(&result)
	return
}

//go:wasmimport wasi:cli/environment@0.2.0 initial-cwd
//go:noescape
func wasmimport_InitialCWD(result *cm.Option[string])
