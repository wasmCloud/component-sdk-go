// Original source: https://github.com/open-telemetry/opentelemetry-go/blob/v1.31.0/exporters/otlp/otlptrace/internal/tracetransform/attribute.go
package convert

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"

	"go.wasmcloud.dev/component/x/wasitel/wasiteltrace/internal/types"
)

// Iterator transforms an attribute iterator into OTLP key-values.
func Iterator(iter attribute.Iterator) []*types.KeyValue {
	l := iter.Len()
	if l == 0 {
		return nil
	}

	out := make([]*types.KeyValue, 0, l)
	for iter.Next() {
		out = append(out, KeyValue(iter.Attribute()))
	}
	return out
}

// ResourceAttributes transforms a Resource OTLP key-values.
func ResourceAttributes(res *resource.Resource) []*types.KeyValue {
	return Iterator(res.Iter())
}

// KeyValues transforms a slice of attribute KeyValues into OTLP key-values.
func KeyValues(attrs []attribute.KeyValue) []*types.KeyValue {
	if len(attrs) == 0 {
		return nil
	}

	out := make([]*types.KeyValue, 0, len(attrs))
	for _, kv := range attrs {
		out = append(out, KeyValue(kv))
	}
	return out
}

// KeyValue transforms an attribute KeyValue into an OTLP key-value.
func KeyValue(kv attribute.KeyValue) *types.KeyValue {
	return &types.KeyValue{Key: string(kv.Key), Value: Value(kv.Value)}
}

// Value transforms an attribute Value into an OTLP AnyValue.
func Value(v attribute.Value) *types.AnyValue {
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
