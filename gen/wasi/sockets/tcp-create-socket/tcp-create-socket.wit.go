// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package tcpcreatesocket represents the imported interface "wasi:sockets/tcp-create-socket@0.2.0".
package tcpcreatesocket

import (
	"go.bytecodealliance.org/cm"
	"go.wasmcloud.dev/component/gen/wasi/sockets/network"
	"go.wasmcloud.dev/component/gen/wasi/sockets/tcp"
)

// Network represents the imported type alias "wasi:sockets/tcp-create-socket@0.2.0#network".
//
// See [network.Network] for more information.
type Network = network.Network

// ErrorCode represents the type alias "wasi:sockets/tcp-create-socket@0.2.0#error-code".
//
// See [network.ErrorCode] for more information.
type ErrorCode = network.ErrorCode

// IPAddressFamily represents the type alias "wasi:sockets/tcp-create-socket@0.2.0#ip-address-family".
//
// See [network.IPAddressFamily] for more information.
type IPAddressFamily = network.IPAddressFamily

// TCPSocket represents the imported type alias "wasi:sockets/tcp-create-socket@0.2.0#tcp-socket".
//
// See [tcp.TCPSocket] for more information.
type TCPSocket = tcp.TCPSocket

// CreateTCPSocket represents the imported function "create-tcp-socket".
//
//	create-tcp-socket: func(address-family: ip-address-family) -> result<tcp-socket,
//	error-code>
//
//go:nosplit
func CreateTCPSocket(addressFamily IPAddressFamily) (result cm.Result[TCPSocket, TCPSocket, ErrorCode]) {
	addressFamily0 := (uint32)(addressFamily)
	wasmimport_CreateTCPSocket((uint32)(addressFamily0), &result)
	return
}
