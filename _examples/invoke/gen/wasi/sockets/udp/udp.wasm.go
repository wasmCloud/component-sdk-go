// Code generated by wit-bindgen-go. DO NOT EDIT.

package udp

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasi:sockets@0.2.0".

//go:wasmimport wasi:sockets/udp@0.2.0 [resource-drop]udp-socket
//go:noescape
func wasmimport_UDPSocketResourceDrop(self0 uint32)

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.address-family
//go:noescape
func wasmimport_UDPSocketAddressFamily(self0 uint32) (result0 uint32)

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.finish-bind
//go:noescape
func wasmimport_UDPSocketFinishBind(self0 uint32, result *cm.Result[ErrorCode, struct{}, ErrorCode])

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.local-address
//go:noescape
func wasmimport_UDPSocketLocalAddress(self0 uint32, result *cm.Result[IPSocketAddressShape, IPSocketAddress, ErrorCode])

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.receive-buffer-size
//go:noescape
func wasmimport_UDPSocketReceiveBufferSize(self0 uint32, result *cm.Result[uint64, uint64, ErrorCode])

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.remote-address
//go:noescape
func wasmimport_UDPSocketRemoteAddress(self0 uint32, result *cm.Result[IPSocketAddressShape, IPSocketAddress, ErrorCode])

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.send-buffer-size
//go:noescape
func wasmimport_UDPSocketSendBufferSize(self0 uint32, result *cm.Result[uint64, uint64, ErrorCode])

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.set-receive-buffer-size
//go:noescape
func wasmimport_UDPSocketSetReceiveBufferSize(self0 uint32, value0 uint64, result *cm.Result[ErrorCode, struct{}, ErrorCode])

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.set-send-buffer-size
//go:noescape
func wasmimport_UDPSocketSetSendBufferSize(self0 uint32, value0 uint64, result *cm.Result[ErrorCode, struct{}, ErrorCode])

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.set-unicast-hop-limit
//go:noescape
func wasmimport_UDPSocketSetUnicastHopLimit(self0 uint32, value0 uint32, result *cm.Result[ErrorCode, struct{}, ErrorCode])

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.start-bind
//go:noescape
func wasmimport_UDPSocketStartBind(self0 uint32, network0 uint32, localAddress0 uint32, localAddress1 uint32, localAddress2 uint32, localAddress3 uint32, localAddress4 uint32, localAddress5 uint32, localAddress6 uint32, localAddress7 uint32, localAddress8 uint32, localAddress9 uint32, localAddress10 uint32, localAddress11 uint32, result *cm.Result[ErrorCode, struct{}, ErrorCode])

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.stream
//go:noescape
func wasmimport_UDPSocketStream(self0 uint32, remoteAddress0 uint32, remoteAddress1 uint32, remoteAddress2 uint32, remoteAddress3 uint32, remoteAddress4 uint32, remoteAddress5 uint32, remoteAddress6 uint32, remoteAddress7 uint32, remoteAddress8 uint32, remoteAddress9 uint32, remoteAddress10 uint32, remoteAddress11 uint32, remoteAddress12 uint32, result *cm.Result[TupleIncomingDatagramStreamOutgoingDatagramStreamShape, cm.Tuple[IncomingDatagramStream, OutgoingDatagramStream], ErrorCode])

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.subscribe
//go:noescape
func wasmimport_UDPSocketSubscribe(self0 uint32) (result0 uint32)

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.unicast-hop-limit
//go:noescape
func wasmimport_UDPSocketUnicastHopLimit(self0 uint32, result *cm.Result[uint8, uint8, ErrorCode])

//go:wasmimport wasi:sockets/udp@0.2.0 [resource-drop]incoming-datagram-stream
//go:noescape
func wasmimport_IncomingDatagramStreamResourceDrop(self0 uint32)

//go:wasmimport wasi:sockets/udp@0.2.0 [method]incoming-datagram-stream.receive
//go:noescape
func wasmimport_IncomingDatagramStreamReceive(self0 uint32, maxResults0 uint64, result *cm.Result[cm.List[IncomingDatagram], cm.List[IncomingDatagram], ErrorCode])

//go:wasmimport wasi:sockets/udp@0.2.0 [method]incoming-datagram-stream.subscribe
//go:noescape
func wasmimport_IncomingDatagramStreamSubscribe(self0 uint32) (result0 uint32)

//go:wasmimport wasi:sockets/udp@0.2.0 [resource-drop]outgoing-datagram-stream
//go:noescape
func wasmimport_OutgoingDatagramStreamResourceDrop(self0 uint32)

//go:wasmimport wasi:sockets/udp@0.2.0 [method]outgoing-datagram-stream.check-send
//go:noescape
func wasmimport_OutgoingDatagramStreamCheckSend(self0 uint32, result *cm.Result[uint64, uint64, ErrorCode])

//go:wasmimport wasi:sockets/udp@0.2.0 [method]outgoing-datagram-stream.send
//go:noescape
func wasmimport_OutgoingDatagramStreamSend(self0 uint32, datagrams0 *OutgoingDatagram, datagrams1 uint32, result *cm.Result[uint64, uint64, ErrorCode])

//go:wasmimport wasi:sockets/udp@0.2.0 [method]outgoing-datagram-stream.subscribe
//go:noescape
func wasmimport_OutgoingDatagramStreamSubscribe(self0 uint32) (result0 uint32)
