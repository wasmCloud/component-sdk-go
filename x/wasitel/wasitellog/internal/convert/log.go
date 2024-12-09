// Original source: https://github.com/open-telemetry/opentelemetry-go/blob/v1.31.0/exporters/otlp/otlplog/otlploghttp/internal/transform/log.go

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package convert

import (
	"time"

	"go.wasmcloud.dev/component/x/wasitel/wasitellog/internal/types"

	"go.opentelemetry.io/otel/attribute"
	api "go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	logsdk "go.opentelemetry.io/otel/sdk/log"
)

// ResourceLogs returns an slice of OTLP ResourceLogs generated from records.
func ResourceLogs(records []logsdk.Record) []*types.ResourceLogs {
	if len(records) == 0 {
		return nil
	}

	resMap := make(map[attribute.Distinct]*types.ResourceLogs)

	type key struct {
		r  attribute.Distinct
		is instrumentation.Scope
	}
	scopeMap := make(map[key]*types.ScopeLogs)

	var resources int
	for _, r := range records {
		res := r.Resource()
		rKey := res.Equivalent()
		scope := r.InstrumentationScope()
		k := key{
			r:  rKey,
			is: scope,
		}
		sl, iOk := scopeMap[k]
		if !iOk {
			sl = new(types.ScopeLogs)
			var emptyScope instrumentation.Scope
			if scope != emptyScope {
				sl.Scope = &types.InstrumentationScope{
					Name:       scope.Name,
					Version:    scope.Version,
					Attributes: AttrIter(scope.Attributes.Iter()),
				}
				sl.SchemaUrl = scope.SchemaURL
			}
			scopeMap[k] = sl
		}

		sl.LogRecords = append(sl.LogRecords, LogRecord(r))
		rl, rOk := resMap[rKey]
		if !rOk {
			resources++
			rl = new(types.ResourceLogs)
			if res.Len() > 0 {
				rl.Resource = &types.Resource{
					Attributes: AttrIter(res.Iter()),
				}
			}
			rl.SchemaUrl = res.SchemaURL()
			resMap[rKey] = rl
		}
		if !iOk {
			rl.ScopeLogs = append(rl.ScopeLogs, sl)
		}
	}

	// Transform the categorized map into a slice
	resLogs := make([]*types.ResourceLogs, 0, resources)
	for _, rl := range resMap {
		resLogs = append(resLogs, rl)
	}

	return resLogs
}

// LogRecord returns an OTLP LogRecord generated from record.
func LogRecord(record logsdk.Record) *types.LogRecord {
	r := &types.LogRecord{
		TimeUnixNano:         timeUnixNano(record.Timestamp()),
		ObservedTimeUnixNano: timeUnixNano(record.ObservedTimestamp()),
		SeverityNumber:       SeverityNumber(record.Severity()),
		SeverityText:         record.SeverityText(),
		Body:                 LogAttrValue(record.Body()),
		Attributes:           make([]*types.KeyValue, 0, record.AttributesLen()),
		Flags:                uint32(record.TraceFlags()),
		// TODO: DroppedAttributesCount: /* ... */,
	}
	record.WalkAttributes(func(kv api.KeyValue) bool {
		r.Attributes = append(r.Attributes, LogAttr(kv))
		return true
	})
	if tID := record.TraceID(); tID.IsValid() {
		r.TraceId = types.NewTraceID(tID)
	}
	if sID := record.SpanID(); sID.IsValid() {
		r.SpanId = types.NewSpanID(sID)
	}
	return r
}

// timeUnixNano returns t as a Unix time, the number of nanoseconds elapsed
// since January 1, 1970 UTC as uint64. The result is undefined if the Unix
// time in nanoseconds cannot be represented by an int64 (a date before the
// year 1678 or after 2262). timeUnixNano on the zero Time returns 0. The
// result does not depend on the location associated with t.
func timeUnixNano(t time.Time) uint64 {
	nano := t.UnixNano()
	if nano < 0 {
		return 0
	}
	return uint64(nano) // nolint:gosec // Overflow checked.
}

// AttrIter transforms an [attribute.Iterator] into OTLP key-values.
func AttrIter(iter attribute.Iterator) []*types.KeyValue {
	l := iter.Len()
	if l == 0 {
		return nil
	}

	out := make([]*types.KeyValue, 0, l)
	for iter.Next() {
		out = append(out, Attr(iter.Attribute()))
	}
	return out
}

// Attrs transforms a slice of [attribute.KeyValue] into OTLP key-values.
func Attrs(attrs []attribute.KeyValue) []*types.KeyValue {
	if len(attrs) == 0 {
		return nil
	}

	out := make([]*types.KeyValue, 0, len(attrs))
	for _, kv := range attrs {
		out = append(out, Attr(kv))
	}
	return out
}

// Attr transforms an [attribute.KeyValue] into an OTLP key-value.
func Attr(kv attribute.KeyValue) *types.KeyValue {
	return &types.KeyValue{Key: string(kv.Key), Value: AttrValue(kv.Value)}
}

// AttrValue transforms an [attribute.Value] into an OTLP AnyValue.
func AttrValue(v attribute.Value) *types.AnyValue {
	av := new(types.AnyValue)
	switch v.Type() {
	case attribute.BOOL:
		av.BoolValue = v.AsBool()
	case attribute.BOOLSLICE:
		av.ArrayValue = &types.ArrayValue{
			Values: boolSliceValues(v.AsBoolSlice()),
		}
	case attribute.INT64:
		av.IntValue = v.AsInt64()
	case attribute.INT64SLICE:
		av.ArrayValue = &types.ArrayValue{
			Values: int64SliceValues(v.AsInt64Slice()),
		}
	case attribute.FLOAT64:
		av.DoubleValue = v.AsFloat64()
	case attribute.FLOAT64SLICE:
		av.ArrayValue = &types.ArrayValue{
			Values: float64SliceValues(v.AsFloat64Slice()),
		}
	case attribute.STRING:
		av.StringValue = v.AsString()
	case attribute.STRINGSLICE:
		av.ArrayValue = &types.ArrayValue{
			Values: stringSliceValues(v.AsStringSlice()),
		}
	default:
		av.StringValue = "INVALID"
	}
	return av
}

func boolSliceValues(vals []bool) []*types.AnyValue {
	converted := make([]*types.AnyValue, len(vals))
	for i, v := range vals {
		converted[i] = &types.AnyValue{
			BoolValue: v,
		}
	}
	return converted
}

func int64SliceValues(vals []int64) []*types.AnyValue {
	converted := make([]*types.AnyValue, len(vals))
	for i, v := range vals {
		converted[i] = &types.AnyValue{
			IntValue: v,
		}
	}
	return converted
}

func float64SliceValues(vals []float64) []*types.AnyValue {
	converted := make([]*types.AnyValue, len(vals))
	for i, v := range vals {
		converted[i] = &types.AnyValue{
			DoubleValue: v,
		}
	}
	return converted
}

func stringSliceValues(vals []string) []*types.AnyValue {
	converted := make([]*types.AnyValue, len(vals))
	for i, v := range vals {
		converted[i] = &types.AnyValue{
			StringValue: v,
		}
	}
	return converted
}

// Attrs transforms a slice of [api.KeyValue] into OTLP key-values.
func LogAttrs(attrs []api.KeyValue) []*types.KeyValue {
	if len(attrs) == 0 {
		return nil
	}

	out := make([]*types.KeyValue, 0, len(attrs))
	for _, kv := range attrs {
		out = append(out, LogAttr(kv))
	}
	return out
}

// LogAttr transforms an [api.KeyValue] into an OTLP key-value.
func LogAttr(attr api.KeyValue) *types.KeyValue {
	return &types.KeyValue{
		Key:   attr.Key,
		Value: LogAttrValue(attr.Value),
	}
}

// LogAttrValues transforms a slice of [api.Value] into an OTLP []AnyValue.
func LogAttrValues(vals []api.Value) []*types.AnyValue {
	if len(vals) == 0 {
		return nil
	}

	out := make([]*types.AnyValue, 0, len(vals))
	for _, v := range vals {
		out = append(out, LogAttrValue(v))
	}
	return out
}

// LogAttrValue transforms an [api.Value] into an OTLP AnyValue.
func LogAttrValue(v api.Value) *types.AnyValue {
	av := new(types.AnyValue)
	switch v.Kind() {
	case api.KindBool:
		av.BoolValue = v.AsBool()
	case api.KindInt64:
		av.IntValue = v.AsInt64()
	case api.KindFloat64:
		av.DoubleValue = v.AsFloat64()
	case api.KindString:
		av.StringValue = v.AsString()
	case api.KindBytes:
		av.BytesValue = v.AsBytes()
	case api.KindSlice:
		av.ArrayValue = &types.ArrayValue{
			Values: LogAttrValues(v.AsSlice()),
		}
	case api.KindMap:
		av.KvlistValue = &types.KeyValueList{
			Values: LogAttrs(v.AsMap()),
		}
	default:
		av.StringValue = "INVALID"
	}
	return av
}

// SeverityNumber transforms a [log.Severity] into an OTLP SeverityNumber.
func SeverityNumber(s api.Severity) types.SeverityNumber {
	switch s {
	case api.SeverityTrace:
		return types.SeverityNumber_SEVERITY_NUMBER_TRACE
	case api.SeverityTrace2:
		return types.SeverityNumber_SEVERITY_NUMBER_TRACE2
	case api.SeverityTrace3:
		return types.SeverityNumber_SEVERITY_NUMBER_TRACE3
	case api.SeverityTrace4:
		return types.SeverityNumber_SEVERITY_NUMBER_TRACE4
	case api.SeverityDebug:
		return types.SeverityNumber_SEVERITY_NUMBER_DEBUG
	case api.SeverityDebug2:
		return types.SeverityNumber_SEVERITY_NUMBER_DEBUG2
	case api.SeverityDebug3:
		return types.SeverityNumber_SEVERITY_NUMBER_DEBUG3
	case api.SeverityDebug4:
		return types.SeverityNumber_SEVERITY_NUMBER_DEBUG4
	case api.SeverityInfo:
		return types.SeverityNumber_SEVERITY_NUMBER_INFO
	case api.SeverityInfo2:
		return types.SeverityNumber_SEVERITY_NUMBER_INFO2
	case api.SeverityInfo3:
		return types.SeverityNumber_SEVERITY_NUMBER_INFO3
	case api.SeverityInfo4:
		return types.SeverityNumber_SEVERITY_NUMBER_INFO4
	case api.SeverityWarn:
		return types.SeverityNumber_SEVERITY_NUMBER_WARN
	case api.SeverityWarn2:
		return types.SeverityNumber_SEVERITY_NUMBER_WARN2
	case api.SeverityWarn3:
		return types.SeverityNumber_SEVERITY_NUMBER_WARN3
	case api.SeverityWarn4:
		return types.SeverityNumber_SEVERITY_NUMBER_WARN4
	case api.SeverityError:
		return types.SeverityNumber_SEVERITY_NUMBER_ERROR
	case api.SeverityError2:
		return types.SeverityNumber_SEVERITY_NUMBER_ERROR2
	case api.SeverityError3:
		return types.SeverityNumber_SEVERITY_NUMBER_ERROR3
	case api.SeverityError4:
		return types.SeverityNumber_SEVERITY_NUMBER_ERROR4
	case api.SeverityFatal:
		return types.SeverityNumber_SEVERITY_NUMBER_FATAL
	case api.SeverityFatal2:
		return types.SeverityNumber_SEVERITY_NUMBER_FATAL2
	case api.SeverityFatal3:
		return types.SeverityNumber_SEVERITY_NUMBER_FATAL3
	case api.SeverityFatal4:
		return types.SeverityNumber_SEVERITY_NUMBER_FATAL4
	}
	return types.SeverityNumber_SEVERITY_NUMBER_UNSPECIFIED
}
