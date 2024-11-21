package types

// Possible values for LogRecord.SeverityNumber.
type SeverityNumber int32

const (
	// UNSPECIFIED is the default SeverityNumber, it MUST NOT be used.
	SeverityNumber_SEVERITY_NUMBER_UNSPECIFIED SeverityNumber = 0
	SeverityNumber_SEVERITY_NUMBER_TRACE       SeverityNumber = 1
	SeverityNumber_SEVERITY_NUMBER_TRACE2      SeverityNumber = 2
	SeverityNumber_SEVERITY_NUMBER_TRACE3      SeverityNumber = 3
	SeverityNumber_SEVERITY_NUMBER_TRACE4      SeverityNumber = 4
	SeverityNumber_SEVERITY_NUMBER_DEBUG       SeverityNumber = 5
	SeverityNumber_SEVERITY_NUMBER_DEBUG2      SeverityNumber = 6
	SeverityNumber_SEVERITY_NUMBER_DEBUG3      SeverityNumber = 7
	SeverityNumber_SEVERITY_NUMBER_DEBUG4      SeverityNumber = 8
	SeverityNumber_SEVERITY_NUMBER_INFO        SeverityNumber = 9
	SeverityNumber_SEVERITY_NUMBER_INFO2       SeverityNumber = 10
	SeverityNumber_SEVERITY_NUMBER_INFO3       SeverityNumber = 11
	SeverityNumber_SEVERITY_NUMBER_INFO4       SeverityNumber = 12
	SeverityNumber_SEVERITY_NUMBER_WARN        SeverityNumber = 13
	SeverityNumber_SEVERITY_NUMBER_WARN2       SeverityNumber = 14
	SeverityNumber_SEVERITY_NUMBER_WARN3       SeverityNumber = 15
	SeverityNumber_SEVERITY_NUMBER_WARN4       SeverityNumber = 16
	SeverityNumber_SEVERITY_NUMBER_ERROR       SeverityNumber = 17
	SeverityNumber_SEVERITY_NUMBER_ERROR2      SeverityNumber = 18
	SeverityNumber_SEVERITY_NUMBER_ERROR3      SeverityNumber = 19
	SeverityNumber_SEVERITY_NUMBER_ERROR4      SeverityNumber = 20
	SeverityNumber_SEVERITY_NUMBER_FATAL       SeverityNumber = 21
	SeverityNumber_SEVERITY_NUMBER_FATAL2      SeverityNumber = 22
	SeverityNumber_SEVERITY_NUMBER_FATAL3      SeverityNumber = 23
	SeverityNumber_SEVERITY_NUMBER_FATAL4      SeverityNumber = 24
)

// Enum value maps for SeverityNumber.
var (
	SeverityNumber_name = map[int32]string{
		0:  "SEVERITY_NUMBER_UNSPECIFIED",
		1:  "SEVERITY_NUMBER_TRACE",
		2:  "SEVERITY_NUMBER_TRACE2",
		3:  "SEVERITY_NUMBER_TRACE3",
		4:  "SEVERITY_NUMBER_TRACE4",
		5:  "SEVERITY_NUMBER_DEBUG",
		6:  "SEVERITY_NUMBER_DEBUG2",
		7:  "SEVERITY_NUMBER_DEBUG3",
		8:  "SEVERITY_NUMBER_DEBUG4",
		9:  "SEVERITY_NUMBER_INFO",
		10: "SEVERITY_NUMBER_INFO2",
		11: "SEVERITY_NUMBER_INFO3",
		12: "SEVERITY_NUMBER_INFO4",
		13: "SEVERITY_NUMBER_WARN",
		14: "SEVERITY_NUMBER_WARN2",
		15: "SEVERITY_NUMBER_WARN3",
		16: "SEVERITY_NUMBER_WARN4",
		17: "SEVERITY_NUMBER_ERROR",
		18: "SEVERITY_NUMBER_ERROR2",
		19: "SEVERITY_NUMBER_ERROR3",
		20: "SEVERITY_NUMBER_ERROR4",
		21: "SEVERITY_NUMBER_FATAL",
		22: "SEVERITY_NUMBER_FATAL2",
		23: "SEVERITY_NUMBER_FATAL3",
		24: "SEVERITY_NUMBER_FATAL4",
	}
	SeverityNumber_value = map[string]int32{
		"SEVERITY_NUMBER_UNSPECIFIED": 0,
		"SEVERITY_NUMBER_TRACE":       1,
		"SEVERITY_NUMBER_TRACE2":      2,
		"SEVERITY_NUMBER_TRACE3":      3,
		"SEVERITY_NUMBER_TRACE4":      4,
		"SEVERITY_NUMBER_DEBUG":       5,
		"SEVERITY_NUMBER_DEBUG2":      6,
		"SEVERITY_NUMBER_DEBUG3":      7,
		"SEVERITY_NUMBER_DEBUG4":      8,
		"SEVERITY_NUMBER_INFO":        9,
		"SEVERITY_NUMBER_INFO2":       10,
		"SEVERITY_NUMBER_INFO3":       11,
		"SEVERITY_NUMBER_INFO4":       12,
		"SEVERITY_NUMBER_WARN":        13,
		"SEVERITY_NUMBER_WARN2":       14,
		"SEVERITY_NUMBER_WARN3":       15,
		"SEVERITY_NUMBER_WARN4":       16,
		"SEVERITY_NUMBER_ERROR":       17,
		"SEVERITY_NUMBER_ERROR2":      18,
		"SEVERITY_NUMBER_ERROR3":      19,
		"SEVERITY_NUMBER_ERROR4":      20,
		"SEVERITY_NUMBER_FATAL":       21,
		"SEVERITY_NUMBER_FATAL2":      22,
		"SEVERITY_NUMBER_FATAL3":      23,
		"SEVERITY_NUMBER_FATAL4":      24,
	}
)

// A collection of ScopeLogs from a Resource.
type ResourceLogs struct {
	// The resource for the logs in this message.
	// If this field is not set then resource info is unknown.
	Resource *Resource `json:"resource,omitempty"`
	// A list of ScopeLogs that originate from a resource.
	ScopeLogs []*ScopeLogs `json:"scope_logs,omitempty"`
	// The Schema URL, if known. This is the identifier of the Schema that the resource data
	// is recorded in. To learn more about Schema URL see
	// https://opentelemetry.io/docs/specs/otel/schemas/#schema-url
	// This schema_url applies to the data in the "resource" field. It does not apply
	// to the data in the "scope_logs" field which have their own schema_url field.
	SchemaUrl string `json:"schema_url,omitempty"`
}

// A collection of Logs produced by a Scope.
type ScopeLogs struct {
	// The instrumentation scope information for the logs in this message.
	// Semantically when InstrumentationScope isn't set, it is equivalent with
	// an empty instrumentation scope name (unknown).
	Scope *InstrumentationScope `json:"scope,omitempty"`
	// A list of log records.
	LogRecords []*LogRecord `json:"log_records,omitempty"`
	// The Schema URL, if known. This is the identifier of the Schema that the log data
	// is recorded in. To learn more about Schema URL see
	// https://opentelemetry.io/docs/specs/otel/schemas/#schema-url
	// This schema_url applies to all logs in the "logs" field.
	SchemaUrl string `json:"schema_url,omitempty"`
}

// A log record according to OpenTelemetry Log Data Model:
// https://github.com/open-telemetry/oteps/blob/main/text/logs/0097-log-data-model.md
type LogRecord struct {
	// time_unix_nano is the time when the event occurred.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	// Value of 0 indicates unknown or missing timestamp.
	TimeUnixNano uint64 `json:"time_unix_nano,omitempty"`
	// Time when the event was observed by the collection system.
	// For events that originate in OpenTelemetry (e.g. using OpenTelemetry Logging SDK)
	// this timestamp is typically set at the generation time and is equal to Timestamp.
	// For events originating externally and collected by OpenTelemetry (e.g. using
	// Collector) this is the time when OpenTelemetry's code observed the event measured
	// by the clock of the OpenTelemetry code. This field MUST be set once the event is
	// observed by OpenTelemetry.
	//
	// For converting OpenTelemetry log data to formats that support only one timestamp or
	// when receiving OpenTelemetry log data by recipients that support only one timestamp
	// internally the following logic is recommended:
	//   - Use time_unix_nano if it is present, otherwise use observed_time_unix_nano.
	//
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	// Value of 0 indicates unknown or missing timestamp.
	ObservedTimeUnixNano uint64 `json:"observed_time_unix_nano,omitempty"`
	// Numerical value of the severity, normalized to values described in Log Data Model.
	// [Optional].
	SeverityNumber SeverityNumber `json:"severity_number,omitempty"`
	// The severity text (also known as log level). The original string representation as
	// it is known at the source. [Optional].
	SeverityText string `json:"severity_text,omitempty"`
	// A value containing the body of the log record. Can be for example a human-readable
	// string message (including multi-line) describing the event in a free form or it can
	// be a structured data composed of arrays and maps of other values. [Optional].
	Body *AnyValue `json:"body,omitempty"`
	// Additional attributes that describe the specific event occurrence. [Optional].
	// Attribute keys MUST be unique (it is not allowed to have more than one
	// attribute with the same key).
	Attributes             []*KeyValue `json:"attributes,omitempty"`
	DroppedAttributesCount uint32      `json:"dropped_attributes_count,omitempty"`
	// Flags, a bit field. 8 least significant bits are the trace flags as
	// defined in W3C Trace Context specification. 24 most significant bits are reserved
	// and must be set to 0. Readers must not assume that 24 most significant bits
	// will be zero and must correctly mask the bits when reading 8-bit trace flag (use
	// flags & LOG_RECORD_FLAGS_TRACE_FLAGS_MASK). [Optional].
	Flags uint32 `json:"flags,omitempty"`
	// A unique identifier for a trace. All logs from the same trace share
	// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes OR
	// of length other than 16 bytes is considered invalid (empty string in OTLP/JSON
	// is zero-length and thus is also invalid).
	//
	// This field is optional.
	//
	// The receivers SHOULD assume that the log record is not associated with a
	// trace if any of the following is true:
	//   - the field is not present,
	//   - the field contains an invalid value.
	TraceId *TraceID `json:"trace_id,omitempty"`
	// A unique identifier for a span within a trace, assigned when the span
	// is created. The ID is an 8-byte array. An ID with all zeroes OR of length
	// other than 8 bytes is considered invalid (empty string in OTLP/JSON
	// is zero-length and thus is also invalid).
	//
	// This field is optional. If the sender specifies a valid span_id then it SHOULD also
	// specify a valid trace_id.
	//
	// The receivers SHOULD assume that the log record is not associated with a
	// span if any of the following is true:
	//   - the field is not present,
	//   - the field contains an invalid value.
	SpanId *SpanID `json:"span_id,omitempty"`
}
