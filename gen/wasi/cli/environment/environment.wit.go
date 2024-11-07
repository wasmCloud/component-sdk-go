// Code generated by component-sdk-go-codegen. DO NOT EDIT.

// Package environment represents the imported interface "wasi:cli/environment@0.2.0".
package environment

import (
	"go.bytecodealliance.org/cm"
)

// GetEnvironment represents the imported function "get-environment".
//
// Get the POSIX-style environment variables.
//
// Each environment variable is provided as a pair of string variable names
// and string value.
//
// Morally, these are a value import, but until value imports are available
// in the component model, this import function should return the same
// values each time it is called.
//
//	get-environment: func() -> list<tuple<string, string>>
//
//go:nosplit
func GetEnvironment() (result cm.List[[2]string]) {
	wasmimport_GetEnvironment(&result)
	return
}

// GetArguments represents the imported function "get-arguments".
//
// Get the POSIX-style arguments to the program.
//
//	get-arguments: func() -> list<string>
//
//go:nosplit
func GetArguments() (result cm.List[string]) {
	wasmimport_GetArguments(&result)
	return
}

// InitialCWD represents the imported function "initial-cwd".
//
// Return a path that programs should use as their initial current working
// directory, interpreting `.` as shorthand for this.
//
//	initial-cwd: func() -> option<string>
//
//go:nosplit
func InitialCWD() (result cm.Option[string]) {
	wasmimport_InitialCWD(&result)
	return
}
