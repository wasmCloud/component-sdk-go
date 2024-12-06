package types

import (
	"encoding/json"

	"go.opentelemetry.io/otel/trace"
)

type TraceID []byte

func NewTraceID(tid trace.TraceID) *TraceID {
	id := TraceID(tid.String())
	return &id
}

func (tid *TraceID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(*tid))
}

func NewSpanID(sid trace.SpanID) *SpanID {
	id := SpanID(sid.String())
	return &id
}

type SpanID []byte

func (sid *SpanID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(*sid))
}
