// Original source: https://github.com/open-telemetry/opentelemetry-proto-go/blob/v1.3.1/slim/otlp/trace/v1/trace.pb.go
package types

// SpanFlags represents constants used to interpret the
// Span.flags field, which is protobuf 'fixed32' type and is to
// be used as bit-fields. Each non-zero value defined in this enum is
// a bit-mask.  To extract the bit-field, for example, use an
// expression like:
//
//	(span.flags & SPAN_FLAGS_TRACE_FLAGS_MASK)
//
// See https://www.w3.org/TR/trace-context-2/#trace-flags for the flag definitions.
//
// Note that Span flags were introduced in version 1.1 of the
// OpenTelemetry protocol.  Older Span producers do not set this
// field, consequently consumers should not rely on the absence of a
// particular flag bit to indicate the presence of a particular feature.
type SpanFlags int32

const (
	// The zero value for the enum. Should not be used for comparisons.
	// Instead use bitwise "and" with the appropriate mask as shown above.
	SpanFlags_SPAN_FLAGS_DO_NOT_USE SpanFlags = 0
	// Bits 0-7 are used for trace flags.
	SpanFlags_SPAN_FLAGS_TRACE_FLAGS_MASK SpanFlags = 255
	// Bits 8 and 9 are used to indicate that the parent span or link span is remote.
	// Bit 8 (`HAS_IS_REMOTE`) indicates whether the value is known.
	// Bit 9 (`IS_REMOTE`) indicates whether the span or link is remote.
	SpanFlags_SPAN_FLAGS_CONTEXT_HAS_IS_REMOTE_MASK SpanFlags = 256
	SpanFlags_SPAN_FLAGS_CONTEXT_IS_REMOTE_MASK     SpanFlags = 512
)

var (
	SpanFlags_name = map[int32]string{
		0:   "SPAN_FLAGS_DO_NOT_USE",
		255: "SPAN_FLAGS_TRACE_FLAGS_MASK",
		256: "SPAN_FLAGS_CONTEXT_HAS_IS_REMOTE_MASK",
		512: "SPAN_FLAGS_CONTEXT_IS_REMOTE_MASK",
	}
	SpanFlags_value = map[string]int32{
		"SPAN_FLAGS_DO_NOT_USE":                 0,
		"SPAN_FLAGS_TRACE_FLAGS_MASK":           255,
		"SPAN_FLAGS_CONTEXT_HAS_IS_REMOTE_MASK": 256,
		"SPAN_FLAGS_CONTEXT_IS_REMOTE_MASK":     512,
	}
)

// End SpanFlags

// SpanKind is the type of span. Can be used to specify additional relationships between spans
// in addition to a parent/child relationship.
type Span_SpanKind int32

const (
	// Unspecified. Do NOT use as default.
	// Implementations MAY assume SpanKind to be INTERNAL when receiving UNSPECIFIED.
	Span_SPAN_KIND_UNSPECIFIED Span_SpanKind = 0
	// Indicates that the span represents an internal operation within an application,
	// as opposed to an operation happening at the boundaries. Default value.
	Span_SPAN_KIND_INTERNAL Span_SpanKind = 1
	// Indicates that the span covers server-side handling of an RPC or other
	// remote network request.
	Span_SPAN_KIND_SERVER Span_SpanKind = 2
	// Indicates that the span describes a request to some remote service.
	Span_SPAN_KIND_CLIENT Span_SpanKind = 3
	// Indicates that the span describes a producer sending a message to a broker.
	// Unlike CLIENT and SERVER, there is often no direct critical path latency relationship
	// between producer and consumer spans. A PRODUCER span ends when the message was accepted
	// by the broker while the logical processing of the message might span a much longer time.
	Span_SPAN_KIND_PRODUCER Span_SpanKind = 4
	// Indicates that the span describes consumer receiving a message from a broker.
	// Like the PRODUCER kind, there is often no direct critical path latency relationship
	// between producer and consumer spans.
	Span_SPAN_KIND_CONSUMER Span_SpanKind = 5
)

// Enum value maps for Span_SpanKind.
var (
	Span_SpanKind_name = map[int32]string{
		0: "SPAN_KIND_UNSPECIFIED",
		1: "SPAN_KIND_INTERNAL",
		2: "SPAN_KIND_SERVER",
		3: "SPAN_KIND_CLIENT",
		4: "SPAN_KIND_PRODUCER",
		5: "SPAN_KIND_CONSUMER",
	}
	Span_SpanKind_value = map[string]int32{
		"SPAN_KIND_UNSPECIFIED": 0,
		"SPAN_KIND_INTERNAL":    1,
		"SPAN_KIND_SERVER":      2,
		"SPAN_KIND_CLIENT":      3,
		"SPAN_KIND_PRODUCER":    4,
		"SPAN_KIND_CONSUMER":    5,
	}
)

// End Span_SpanKind

// For the semantics of status codes see
// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/api.md#set-status
type Status_StatusCode int32

const (
	// The default status.
	Status_STATUS_CODE_UNSET Status_StatusCode = 0
	// The Span has been validated by an Application developer or Operator to
	// have completed successfully.
	Status_STATUS_CODE_OK Status_StatusCode = 1
	// The Span contains an error.
	Status_STATUS_CODE_ERROR Status_StatusCode = 2
)

// Enum value maps for Status_StatusCode.
var (
	Status_StatusCode_name = map[int32]string{
		0: "STATUS_CODE_UNSET",
		1: "STATUS_CODE_OK",
		2: "STATUS_CODE_ERROR",
	}
	Status_StatusCode_value = map[string]int32{
		"STATUS_CODE_UNSET": 0,
		"STATUS_CODE_OK":    1,
		"STATUS_CODE_ERROR": 2,
	}
)

// End Status_StatusCode

// A collection of ScopeSpans from a Resource.
type ResourceSpans struct {
	// The resource for the spans in this message.
	// If this field is not set then no resource info is known.
	Resource *Resource `json:"resource,omitempty"`
	// A list of ScopeSpans that originate from a resource.
	ScopeSpans []*ScopeSpans `json:"scope_spans,omitempty"`
	// The Schema URL, if known. This is the identifier of the Schema that the resource data
	// is recorded in. To learn more about Schema URL see
	// https://opentelemetry.io/docs/specs/otel/schemas/#schema-url
	// This schema_url applies to the data in the "resource" field. It does not apply
	// to the data in the "scope_spans" field which have their own schema_url field.
	SchemaUrl string `json:"schema_url,omitempty"`
}

// A collection of Spans produced by an InstrumentationScope.
type ScopeSpans struct {
	// The instrumentation scope information for the spans in this message.
	// Semantically when InstrumentationScope isn't set, it is equivalent with
	// an empty instrumentation scope name (unknown).
	Scope *InstrumentationScope `json:"scope,omitempty"`
	// A list of Spans that originate from an instrumentation scope.
	Spans []*Span `json:"spans,omitempty"`
	// The Schema URL, if known. This is the identifier of the Schema that the span data
	// is recorded in. To learn more about Schema URL see
	// https://opentelemetry.io/docs/specs/otel/schemas/#schema-url
	// This schema_url applies to all spans and span events in the "spans" field.
	SchemaUrl string `json:"schema_url,omitempty"`
}

// A Span represents a single operation performed by a single component of the system.
type Span struct {
	// A unique identifier for a trace. All spans from the same trace share
	// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes OR
	// of length other than 16 bytes is considered invalid (empty string in OTLP/JSON
	// is zero-length and thus is also invalid).
	//
	// This field is required.
	TraceId *TraceID `json:"trace_id,omitempty"`
	// A unique identifier for a span within a trace, assigned when the span
	// is created. The ID is an 8-byte array. An ID with all zeroes OR of length
	// other than 8 bytes is considered invalid (empty string in OTLP/JSON
	// is zero-length and thus is also invalid).
	//
	// This field is required.
	SpanId *SpanID `json:"span_id,omitempty"`
	// trace_state conveys information about request position in multiple distributed tracing graphs.
	// It is a trace_state in w3c-trace-context format: https://www.w3.org/TR/trace-context/#tracestate-header
	// See also https://github.com/w3c/distributed-tracing for more details about this field.
	TraceState string `json:"trace_state,omitempty"`
	// The `span_id` of this span's parent span. If this is a root span, then this
	// field must be empty. The ID is an 8-byte array.
	ParentSpanId *SpanID `json:"parent_span_id,omitempty"`
	// Flags, a bit field.
	//
	// Bits 0-7 (8 least significant bits) are the trace flags as defined in W3C Trace
	// Context specification. To read the 8-bit W3C trace flag, use
	// `flags & SPAN_FLAGS_TRACE_FLAGS_MASK`.
	//
	// See https://www.w3.org/TR/trace-context-2/#trace-flags for the flag definitions.
	//
	// Bits 8 and 9 represent the 3 states of whether a span's parent
	// is remote. The states are (unknown, is not remote, is remote).
	// To read whether the value is known, use `(flags & SPAN_FLAGS_CONTEXT_HAS_IS_REMOTE_MASK) != 0`.
	// To read whether the span is remote, use `(flags & SPAN_FLAGS_CONTEXT_IS_REMOTE_MASK) != 0`.
	//
	// When creating span messages, if the message is logically forwarded from another source
	// with an equivalent flags fields (i.e., usually another OTLP span message), the field SHOULD
	// be copied as-is. If creating from a source that does not have an equivalent flags field
	// (such as a runtime representation of an OpenTelemetry span), the high 22 bits MUST
	// be set to zero.
	// Readers MUST NOT assume that bits 10-31 (22 most significant bits) will be zero.
	//
	// [Optional].
	Flags uint32 `json:"flags,omitempty"`
	// A description of the span's operation.
	//
	// For example, the name can be a qualified method name or a file name
	// and a line number where the operation is called. A best practice is to use
	// the same display name at the same call point in an application.
	// This makes it easier to correlate spans in different traces.
	//
	// This field is semantically required to be set to non-empty string.
	// Empty value is equivalent to an unknown span name.
	//
	// This field is required.
	Name string `json:"name,omitempty"`
	// Distinguishes between spans generated in a particular context. For example,
	// two spans with the same name may be distinguished using `CLIENT` (caller)
	// and `SERVER` (callee) to identify queueing latency associated with the span.
	Kind Span_SpanKind `json:"kind,omitempty"`
	// start_time_unix_nano is the start time of the span. On the client side, this is the time
	// kept by the local machine where the span execution starts. On the server side, this
	// is the time when the server's application handler starts running.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	//
	// This field is semantically required and it is expected that end_time >= start_time.
	StartTimeUnixNano uint64 `json:"start_time_unix_nano,omitempty"`
	// end_time_unix_nano is the end time of the span. On the client side, this is the time
	// kept by the local machine where the span execution ends. On the server side, this
	// is the time when the server application handler stops running.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	//
	// This field is semantically required and it is expected that end_time >= start_time.
	EndTimeUnixNano uint64 `json:"end_time_unix_nano,omitempty"`
	// attributes is a collection of key/value pairs. Note, global attributes
	// like server name can be set using the resource API. Examples of attributes:
	//
	//     "/http/user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
	//     "/http/server_latency": 300
	//     "example.com/myattribute": true
	//     "example.com/score": 10.239
	//
	// The OpenTelemetry API specification further restricts the allowed value types:
	// https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/common/README.md#attribute
	// Attribute keys MUST be unique (it is not allowed to have more than one
	// attribute with the same key).
	Attributes []*KeyValue `json:"attributes,omitempty"`
	// dropped_attributes_count is the number of attributes that were discarded. Attributes
	// can be discarded because their keys are too long or because there are too many
	// attributes. If this value is 0, then no attributes were dropped.
	DroppedAttributesCount uint32 `json:"dropped_attributes_count,omitempty"`
	// events is a collection of Event items.
	Events []*Span_Event `json:"events,omitempty"`
	// dropped_events_count is the number of dropped events. If the value is 0, then no
	// events were dropped.
	DroppedEventsCount uint32 `json:"dropped_events_count,omitempty"`
	// links is a collection of Links, which are references from this span to a span
	// in the same or different trace.
	Links []*Span_Link `json:"links,omitempty"`
	// dropped_links_count is the number of dropped links after the maximum size was
	// enforced. If this value is 0, then no links were dropped.
	DroppedLinksCount uint32 `json:"dropped_links_count,omitempty"`
	// An optional final status for this span. Semantically when Status isn't set, it means
	// span's status code is unset, i.e. assume STATUS_CODE_UNSET (code = 0).
	Status *Status `json:"status,omitempty"`
}

// The Status type defines a logical error model that is suitable for different
// programming environments, including REST APIs and RPC APIs.
type Status struct {
	// A developer-facing human readable error message.
	Message string `json:"message,omitempty"`
	// The status code.
	Code Status_StatusCode `json:"code,omitempty"`
}

// Event is a time-stamped annotation of the span, consisting of user-supplied
// text description and key-value pairs.
type Span_Event struct {
	// time_unix_nano is the time the event occurred.
	TimeUnixNano uint64 `json:"time_unix_nano,omitempty"`
	// name of the event.
	// This field is semantically required to be set to non-empty string.
	Name string `json:"name,omitempty"`
	// attributes is a collection of attribute key/value pairs on the event.
	// Attribute keys MUST be unique (it is not allowed to have more than one
	// attribute with the same key).
	Attributes []*KeyValue `json:"attributes,omitempty"`
	// dropped_attributes_count is the number of dropped attributes. If the value is 0,
	// then no attributes were dropped.
	DroppedAttributesCount uint32 `json:"dropped_attributes_count,omitempty"`
}

// A pointer from the current span to another span in the same trace or in a
// different trace. For example, this can be used in batching operations,
// where a single batch handler processes multiple requests from different
// traces or when the handler receives a request from a different project.
type Span_Link struct {
	// A unique identifier of a trace that this linked span is part of. The ID is a
	// 16-byte array.
	TraceId *TraceID `json:"trace_id,omitempty"`
	// A unique identifier for the linked span. The ID is an 8-byte array.
	SpanId *SpanID `json:"span_id,omitempty"`
	// The trace_state associated with the link.
	TraceState string `json:"trace_state,omitempty"`
	// attributes is a collection of attribute key/value pairs on the link.
	// Attribute keys MUST be unique (it is not allowed to have more than one
	// attribute with the same key).
	Attributes []*KeyValue `json:"attributes,omitempty"`
	// dropped_attributes_count is the number of dropped attributes. If the value is 0,
	// then no attributes were dropped.
	DroppedAttributesCount uint32 `json:"dropped_attributes_count,omitempty"`
	// Flags, a bit field.
	//
	// Bits 0-7 (8 least significant bits) are the trace flags as defined in W3C Trace
	// Context specification. To read the 8-bit W3C trace flag, use
	// `flags & SPAN_FLAGS_TRACE_FLAGS_MASK`.
	//
	// See https://www.w3.org/TR/trace-context-2/#trace-flags for the flag definitions.
	//
	// Bits 8 and 9 represent the 3 states of whether the link is remote.
	// The states are (unknown, is not remote, is remote).
	// To read whether the value is known, use `(flags & SPAN_FLAGS_CONTEXT_HAS_IS_REMOTE_MASK) != 0`.
	// To read whether the link is remote, use `(flags & SPAN_FLAGS_CONTEXT_IS_REMOTE_MASK) != 0`.
	//
	// Readers MUST NOT assume that bits 10-31 (22 most significant bits) will be zero.
	// When creating new spans, bits 10-31 (most-significant 22-bits) MUST be zero.
	//
	// [Optional].
	Flags uint32 `json:"flags,omitempty"`
}
