// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package exit represents the imported interface "wasi:cli/exit@0.2.0".
package exit

import (
	"go.bytecodealliance.org/cm"
)

// Exit represents the imported function "exit".
//
// Exit the current instance and any linked instances.
//
//	exit: func(status: result)
//
//go:nosplit
func Exit(status cm.BoolResult) {
	status0 := cm.BoolToU32(status)
	wasmimport_Exit((uint32)(status0))
	return
}
