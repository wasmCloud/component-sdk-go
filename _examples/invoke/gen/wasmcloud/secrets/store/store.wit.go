// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package store represents the imported interface "wasmcloud:secrets/store@0.1.0-draft".
package store

import (
	"go.bytecodealliance.org/cm"
)

// SecretsError represents the variant "wasmcloud:secrets/store@0.1.0-draft#secrets-error".
//
//	variant secrets-error {
//		upstream(string),
//		io(string),
//		not-found,
//	}
type SecretsError cm.Variant[uint8, string, string]

// SecretsErrorUpstream returns a [SecretsError] of case "upstream".
func SecretsErrorUpstream(data string) SecretsError {
	return cm.New[SecretsError](0, data)
}

// Upstream returns a non-nil *[string] if [SecretsError] represents the variant case "upstream".
func (self *SecretsError) Upstream() *string {
	return cm.Case[string](self, 0)
}

// SecretsErrorIO returns a [SecretsError] of case "io".
func SecretsErrorIO(data string) SecretsError {
	return cm.New[SecretsError](1, data)
}

// IO returns a non-nil *[string] if [SecretsError] represents the variant case "io".
func (self *SecretsError) IO() *string {
	return cm.Case[string](self, 1)
}

// SecretsErrorNotFound returns a [SecretsError] of case "not-found".
func SecretsErrorNotFound() SecretsError {
	var data struct{}
	return cm.New[SecretsError](2, data)
}

// NotFound returns true if [SecretsError] represents the variant case "not-found".
func (self *SecretsError) NotFound() bool {
	return self.Tag() == 2
}

var stringsSecretsError = [3]string{
	"upstream",
	"io",
	"not-found",
}

// String implements [fmt.Stringer], returning the variant case name of v.
func (v SecretsError) String() string {
	return stringsSecretsError[v.Tag()]
}

// SecretValue represents the variant "wasmcloud:secrets/store@0.1.0-draft#secret-value".
//
//	variant secret-value {
//		%string(string),
//		bytes(list<u8>),
//	}
type SecretValue cm.Variant[uint8, string, cm.List[uint8]]

// SecretValueString_ returns a [SecretValue] of case "string".
func SecretValueString_(data string) SecretValue {
	return cm.New[SecretValue](0, data)
}

// String_ returns a non-nil *[string] if [SecretValue] represents the variant case "string".
func (self *SecretValue) String_() *string {
	return cm.Case[string](self, 0)
}

// SecretValueBytes returns a [SecretValue] of case "bytes".
func SecretValueBytes(data cm.List[uint8]) SecretValue {
	return cm.New[SecretValue](1, data)
}

// Bytes returns a non-nil *[cm.List[uint8]] if [SecretValue] represents the variant case "bytes".
func (self *SecretValue) Bytes() *cm.List[uint8] {
	return cm.Case[cm.List[uint8]](self, 1)
}

var stringsSecretValue = [2]string{
	"string",
	"bytes",
}

// String implements [fmt.Stringer], returning the variant case name of v.
func (v SecretValue) String() string {
	return stringsSecretValue[v.Tag()]
}

// Secret represents the imported resource "wasmcloud:secrets/store@0.1.0-draft#secret".
//
//	resource secret
type Secret cm.Resource

// ResourceDrop represents the imported resource-drop for resource "secret".
//
// Drops a resource handle.
//
//go:nosplit
func (self Secret) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_SecretResourceDrop((uint32)(self0))
	return
}

// Get represents the imported function "get".
//
//	get: func(key: string) -> result<secret, secrets-error>
//
//go:nosplit
func Get(key string) (result cm.Result[SecretsErrorShape, Secret, SecretsError]) {
	key0, key1 := cm.LowerString(key)
	wasmimport_Get((*uint8)(key0), (uint32)(key1), &result)
	return
}
