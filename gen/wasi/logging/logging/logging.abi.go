//go:build !test

package logging

//
//go:wasmimport wasi:logging/logging log
//go:noescape
func wasmimport_Log(level0 uint32, context0 *uint8, context1 uint32, message0 *uint8, message1 uint32)
