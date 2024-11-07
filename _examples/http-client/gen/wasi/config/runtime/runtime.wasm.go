// Code generated by component-sdk-go-codegen. DO NOT EDIT.

package runtime

import (
	"go.bytecodealliance.org/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasi:config@0.2.0-draft".

//go:wasmimport wasi:config/runtime@0.2.0-draft get
//go:noescape
func wasmimport_Get(key0 *uint8, key1 uint32, result *cm.Result[OptionStringShape, cm.Option[string], ConfigError])

//go:wasmimport wasi:config/runtime@0.2.0-draft get-all
//go:noescape
func wasmimport_GetAll(result *cm.Result[ConfigErrorShape, cm.List[[2]string], ConfigError])
