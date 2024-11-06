// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package streams represents the imported interface "wasi:io/streams@0.2.0".
package streams

import (
	ioerror "github.com/wasmCloud/component-sdk-go/_examples/invoke/gen/wasi/io/error"
	"github.com/wasmCloud/component-sdk-go/_examples/invoke/gen/wasi/io/poll"
	"go.bytecodealliance.org/cm"
)

// Error represents the imported type alias "wasi:io/streams@0.2.0#error".
//
// See [ioerror.Error] for more information.
type Error = ioerror.Error

// Pollable represents the imported type alias "wasi:io/streams@0.2.0#pollable".
//
// See [poll.Pollable] for more information.
type Pollable = poll.Pollable

// StreamError represents the imported variant "wasi:io/streams@0.2.0#stream-error".
//
//	variant stream-error {
//		last-operation-failed(error),
//		closed,
//	}
type StreamError cm.Variant[uint8, Error, Error]

// StreamErrorLastOperationFailed returns a [StreamError] of case "last-operation-failed".
func StreamErrorLastOperationFailed(data Error) StreamError {
	return cm.New[StreamError](0, data)
}

// LastOperationFailed returns a non-nil *[Error] if [StreamError] represents the variant case "last-operation-failed".
func (self *StreamError) LastOperationFailed() *Error {
	return cm.Case[Error](self, 0)
}

// StreamErrorClosed returns a [StreamError] of case "closed".
func StreamErrorClosed() StreamError {
	var data struct{}
	return cm.New[StreamError](1, data)
}

// Closed returns true if [StreamError] represents the variant case "closed".
func (self *StreamError) Closed() bool {
	return self.Tag() == 1
}

var stringsStreamError = [2]string{
	"last-operation-failed",
	"closed",
}

// String implements [fmt.Stringer], returning the variant case name of v.
func (v StreamError) String() string {
	return stringsStreamError[v.Tag()]
}

// InputStream represents the imported resource "wasi:io/streams@0.2.0#input-stream".
//
//	resource input-stream
type InputStream cm.Resource

// ResourceDrop represents the imported resource-drop for resource "input-stream".
//
// Drops a resource handle.
//
//go:nosplit
func (self InputStream) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_InputStreamResourceDrop((uint32)(self0))
	return
}

// BlockingRead represents the imported method "blocking-read".
//
//	blocking-read: func(len: u64) -> result<list<u8>, stream-error>
//
//go:nosplit
func (self InputStream) BlockingRead(len_ uint64) (result cm.Result[cm.List[uint8], cm.List[uint8], StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_InputStreamBlockingRead((uint32)(self0), (uint64)(len0), &result)
	return
}

// BlockingSkip represents the imported method "blocking-skip".
//
//	blocking-skip: func(len: u64) -> result<u64, stream-error>
//
//go:nosplit
func (self InputStream) BlockingSkip(len_ uint64) (result cm.Result[uint64, uint64, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_InputStreamBlockingSkip((uint32)(self0), (uint64)(len0), &result)
	return
}

// Read represents the imported method "read".
//
//	read: func(len: u64) -> result<list<u8>, stream-error>
//
//go:nosplit
func (self InputStream) Read(len_ uint64) (result cm.Result[cm.List[uint8], cm.List[uint8], StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_InputStreamRead((uint32)(self0), (uint64)(len0), &result)
	return
}

// Skip represents the imported method "skip".
//
//	skip: func(len: u64) -> result<u64, stream-error>
//
//go:nosplit
func (self InputStream) Skip(len_ uint64) (result cm.Result[uint64, uint64, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_InputStreamSkip((uint32)(self0), (uint64)(len0), &result)
	return
}

// Subscribe represents the imported method "subscribe".
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self InputStream) Subscribe() (result Pollable) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_InputStreamSubscribe((uint32)(self0))
	result = cm.Reinterpret[Pollable]((uint32)(result0))
	return
}

// OutputStream represents the imported resource "wasi:io/streams@0.2.0#output-stream".
//
//	resource output-stream
type OutputStream cm.Resource

// ResourceDrop represents the imported resource-drop for resource "output-stream".
//
// Drops a resource handle.
//
//go:nosplit
func (self OutputStream) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutputStreamResourceDrop((uint32)(self0))
	return
}

// BlockingFlush represents the imported method "blocking-flush".
//
//	blocking-flush: func() -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) BlockingFlush() (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutputStreamBlockingFlush((uint32)(self0), &result)
	return
}

// BlockingSplice represents the imported method "blocking-splice".
//
//	blocking-splice: func(src: borrow<input-stream>, len: u64) -> result<u64, stream-error>
//
//go:nosplit
func (self OutputStream) BlockingSplice(src InputStream, len_ uint64) (result cm.Result[uint64, uint64, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	src0 := cm.Reinterpret[uint32](src)
	len0 := (uint64)(len_)
	wasmimport_OutputStreamBlockingSplice((uint32)(self0), (uint32)(src0), (uint64)(len0), &result)
	return
}

// BlockingWriteAndFlush represents the imported method "blocking-write-and-flush".
//
//	blocking-write-and-flush: func(contents: list<u8>) -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) BlockingWriteAndFlush(contents cm.List[uint8]) (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	contents0, contents1 := cm.LowerList(contents)
	wasmimport_OutputStreamBlockingWriteAndFlush((uint32)(self0), (*uint8)(contents0), (uint32)(contents1), &result)
	return
}

// BlockingWriteZeroesAndFlush represents the imported method "blocking-write-zeroes-and-flush".
//
//	blocking-write-zeroes-and-flush: func(len: u64) -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) BlockingWriteZeroesAndFlush(len_ uint64) (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_OutputStreamBlockingWriteZeroesAndFlush((uint32)(self0), (uint64)(len0), &result)
	return
}

// CheckWrite represents the imported method "check-write".
//
//	check-write: func() -> result<u64, stream-error>
//
//go:nosplit
func (self OutputStream) CheckWrite() (result cm.Result[uint64, uint64, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutputStreamCheckWrite((uint32)(self0), &result)
	return
}

// Flush represents the imported method "flush".
//
//	flush: func() -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) Flush() (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutputStreamFlush((uint32)(self0), &result)
	return
}

// Splice represents the imported method "splice".
//
//	splice: func(src: borrow<input-stream>, len: u64) -> result<u64, stream-error>
//
//go:nosplit
func (self OutputStream) Splice(src InputStream, len_ uint64) (result cm.Result[uint64, uint64, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	src0 := cm.Reinterpret[uint32](src)
	len0 := (uint64)(len_)
	wasmimport_OutputStreamSplice((uint32)(self0), (uint32)(src0), (uint64)(len0), &result)
	return
}

// Subscribe represents the imported method "subscribe".
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self OutputStream) Subscribe() (result Pollable) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_OutputStreamSubscribe((uint32)(self0))
	result = cm.Reinterpret[Pollable]((uint32)(result0))
	return
}

// Write represents the imported method "write".
//
//	write: func(contents: list<u8>) -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) Write(contents cm.List[uint8]) (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	contents0, contents1 := cm.LowerList(contents)
	wasmimport_OutputStreamWrite((uint32)(self0), (*uint8)(contents0), (uint32)(contents1), &result)
	return
}

// WriteZeroes represents the imported method "write-zeroes".
//
//	write-zeroes: func(len: u64) -> result<_, stream-error>
//
//go:nosplit
func (self OutputStream) WriteZeroes(len_ uint64) (result cm.Result[StreamError, struct{}, StreamError]) {
	self0 := cm.Reinterpret[uint32](self)
	len0 := (uint64)(len_)
	wasmimport_OutputStreamWriteZeroes((uint32)(self0), (uint64)(len0), &result)
	return
}
