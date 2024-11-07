// Code generated by component-sdk-go-codegen. DO NOT EDIT.

// Package terminalstdin represents the imported interface "wasi:cli/terminal-stdin@0.2.0".
//
// An interface providing an optional `terminal-input` for stdin as a
// link-time authority.
package terminalstdin

import (
	"go.bytecodealliance.org/cm"
	terminalinput "go.wasmcloud.dev/component/gen/wasi/cli/terminal-input"
)

// TerminalInput represents the imported type alias "wasi:cli/terminal-stdin@0.2.0#terminal-input".
//
// See [terminalinput.TerminalInput] for more information.
type TerminalInput = terminalinput.TerminalInput

// GetTerminalStdin represents the imported function "get-terminal-stdin".
//
// If stdin is connected to a terminal, return a `terminal-input` handle
// allowing further interaction with it.
//
//	get-terminal-stdin: func() -> option<terminal-input>
//
//go:nosplit
func GetTerminalStdin() (result cm.Option[TerminalInput]) {
	wasmimport_GetTerminalStdin(&result)
	return
}
