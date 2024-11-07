// Code generated by component-sdk-go-codegen. DO NOT EDIT.

// Package terminaloutput represents the imported interface "wasi:cli/terminal-output@0.2.0".
//
// Terminal output.
//
// In the future, this may include functions for querying the terminal
// size, being notified of terminal size changes, querying supported
// features, and so on.
package terminaloutput

import (
	"go.bytecodealliance.org/cm"
)

// TerminalOutput represents the imported resource "wasi:cli/terminal-output@0.2.0#terminal-output".
//
// The output side of a terminal.
//
//	resource terminal-output
type TerminalOutput cm.Resource

// ResourceDrop represents the imported resource-drop for resource "terminal-output".
//
// Drops a resource handle.
//
//go:nosplit
func (self TerminalOutput) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_TerminalOutputResourceDrop((uint32)(self0))
	return
}
