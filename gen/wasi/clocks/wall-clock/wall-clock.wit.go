// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package wallclock represents the imported interface "wasi:clocks/wall-clock@0.2.0".
package wallclock

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// DateTime represents the record "wasi:clocks/wall-clock@0.2.0#datetime".
//
//	record datetime {
//		seconds: u64,
//		nanoseconds: u32,
//	}
type DateTime struct {
	_           cm.HostLayout
	Seconds     uint64
	Nanoseconds uint32
}

// Now represents the imported function "now".
//
//	now: func() -> datetime
//
//go:nosplit
func Now() (result DateTime) {
	wasmimport_Now(&result)
	return
}

// Resolution represents the imported function "resolution".
//
//	resolution: func() -> datetime
//
//go:nosplit
func Resolution() (result DateTime) {
	wasmimport_Resolution(&result)
	return
}
