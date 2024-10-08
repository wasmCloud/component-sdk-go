// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package tcpcreatesocket represents the imported interface "wasi:sockets/tcp-create-socket@0.2.0".
package tcpcreatesocket

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"github.com/wasmCloud/component-sdk-go/_examples/http-client/gen/wasi/sockets/network"
	"github.com/wasmCloud/component-sdk-go/_examples/http-client/gen/wasi/sockets/tcp"
)

// CreateTCPSocket represents the imported function "create-tcp-socket".
//
//	create-tcp-socket: func(address-family: ip-address-family) -> result<tcp-socket,
//	error-code>
//
//go:nosplit
func CreateTCPSocket(addressFamily network.IPAddressFamily) (result cm.Result[tcp.TCPSocket, tcp.TCPSocket, network.ErrorCode]) {
	addressFamily0 := (uint32)(addressFamily)
	wasmimport_CreateTCPSocket((uint32)(addressFamily0), &result)
	return
}

//go:wasmimport wasi:sockets/tcp-create-socket@0.2.0 create-tcp-socket
//go:noescape
func wasmimport_CreateTCPSocket(addressFamily0 uint32, result *cm.Result[tcp.TCPSocket, tcp.TCPSocket, network.ErrorCode])
