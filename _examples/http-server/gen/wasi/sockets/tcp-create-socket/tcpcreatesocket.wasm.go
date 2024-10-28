// Code generated by wit-bindgen-go. DO NOT EDIT.

package tcpcreatesocket

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasi:sockets@0.2.0".

//go:wasmimport wasi:sockets/tcp-create-socket@0.2.0 create-tcp-socket
//go:noescape
func wasmimport_CreateTCPSocket(addressFamily0 uint32, result *cm.Result[TCPSocket, TCPSocket, ErrorCode])
