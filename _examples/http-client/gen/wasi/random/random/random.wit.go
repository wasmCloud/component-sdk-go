// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package random represents the imported interface "wasi:random/random@0.2.0".
package random

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// GetRandomBytes represents the imported function "get-random-bytes".
//
//	get-random-bytes: func(len: u64) -> list<u8>
//
//go:nosplit
func GetRandomBytes(len_ uint64) (result cm.List[uint8]) {
	len0 := (uint64)(len_)
	wasmimport_GetRandomBytes((uint64)(len0), &result)
	return
}

// GetRandomU64 represents the imported function "get-random-u64".
//
//	get-random-u64: func() -> u64
//
//go:nosplit
func GetRandomU64() (result uint64) {
	result0 := wasmimport_GetRandomU64()
	result = (uint64)((uint64)(result0))
	return
}
