// Original source: https://github.com/open-telemetry/opentelemetry-proto-go/blob/v1.3.1/slim/otlp/common/v1/common.pb.go
package types

// AnyValue is used to represent any type of attribute value. AnyValue may contain a
// primitive value such as a string or integer or it may contain an arbitrary nested
// object containing arrays, key-value lists and primitives.
type AnyValue struct {
	// The value is one of the listed fields. It is valid for all values to be unspecified
	// in which case this AnyValue is considered to be "empty".
	//
	// Types that are assignable to Value:
	//	*AnyValue_StringValue
	//	*AnyValue_BoolValue
	//	*AnyValue_IntValue
	//	*AnyValue_DoubleValue
	//	*AnyValue_ArrayValue
	//	*AnyValue_KvlistValue
	//	*AnyValue_BytesValue
	StringValue string        `json:"stringValue,omitempty"`
	BoolValue   bool          `json:"boolValue,omitempty"`
	IntValue    int64         `json:"intValue,omitempty"`
	DoubleValue float64       `json:"doubleValue,omitempty"`
	ArrayValue  *ArrayValue   `json:"arrayValue,omitempty"`
	KvlistValue *KeyValueList `json:"kvlistValue,omitempty"`
	BytesValue  []byte        `json:"bytesValue,omitempty"`
}

// ArrayValue is a list of AnyValue messages. We need ArrayValue as a message
// since oneof in AnyValue does not allow repeated fields.
type ArrayValue struct {
	// Array of values. The array may be empty (contain 0 elements).
	Values []*AnyValue `json:"values,omitempty"`
}

// KeyValueList is a list of KeyValue messages. We need KeyValueList as a message
// since `oneof` in AnyValue does not allow repeated fields. Everywhere else where we need
// a list of KeyValue messages (e.g. in Span) we use `repeated KeyValue` directly to
// avoid unnecessary extra wrapping (which slows down the protocol). The 2 approaches
// are semantically equivalent.
type KeyValueList struct {
	// A collection of key/value pairs of key-value pairs. The list may be empty (may
	// contain 0 elements).
	// The keys MUST be unique (it is not allowed to have more than one
	// value with the same key).
	Values []*KeyValue `json:"values,omitempty"`
}

// KeyValue is a key-value pair that is used to store Span attributes, Link
// attributes, etc.
type KeyValue struct {
	Key   string    `json:"key,omitempty"`
	Value *AnyValue `json:"value,omitempty"`
}

// InstrumentationScope is a message representing the instrumentation scope information
// such as the fully qualified name and version.
type InstrumentationScope struct {
	// An empty instrumentation scope name means the name is unknown.
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	// Additional attributes that describe the scope. [Optional].
	// Attribute keys MUST be unique (it is not allowed to have more than one
	// attribute with the same key).
	Attributes             []*KeyValue `json:"attributes,omitempty"`
	DroppedAttributesCount uint32      `json:"dropped_attributes_count,omitempty"`
}
