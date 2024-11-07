// Original source: https://github.com/open-telemetry/opentelemetry-go/blob/v1.31.0/exporters/otlp/otlptrace/internal/tracetransform/span.go
package convert

import (
	"math"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/sdk/instrumentation"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"

	"go.wasmcloud.dev/component/x/wasitel/wasiteltrace/internal/types"
)

// ResourceSpans transforms a slice of OpenTelemetry spans into a slice of
// OTLP ResourceSpans.
func ResourceSpans(spans []tracesdk.ReadOnlySpan) []*types.ResourceSpans {
	if len(spans) == 0 {
		return nil
	}

	rsm := make(map[attribute.Distinct]*types.ResourceSpans)

	type key struct {
		r  attribute.Distinct
		is instrumentation.Scope
	}
	ssm := make(map[key]*types.ScopeSpans)

	var resources int
	for _, sdkspan := range spans {
		if sdkspan == nil {
			continue
		}

		rKey := sdkspan.Resource().Equivalent()
		k := key{
			r:  rKey,
			is: sdkspan.InstrumentationScope(),
		}
		scopeSpan, iOk := ssm[k]
		if !iOk {
			// Either the resource or instrumentation scope were unknown.
			scopeSpan = &types.ScopeSpans{
				Scope:     convertInstrumentationScope(sdkspan.InstrumentationScope()),
				Spans:     []*types.Span{},
				SchemaUrl: sdkspan.InstrumentationScope().SchemaURL,
			}
		}
		scopeSpan.Spans = append(scopeSpan.Spans, span(sdkspan))
		ssm[k] = scopeSpan

		rs, rOk := rsm[rKey]
		if !rOk {
			resources++
			// The resource was unknown.
			rs = &types.ResourceSpans{
				Resource:   Resource(sdkspan.Resource()),
				ScopeSpans: []*types.ScopeSpans{scopeSpan},
				SchemaUrl:  sdkspan.Resource().SchemaURL(),
			}
			rsm[rKey] = rs
			continue
		}

		// The resource has been seen before. Check if the instrumentation
		// library lookup was unknown because if so we need to add it to the
		// ResourceSpanss. Otherwise, the instrumentation library has already
		// been seen and the append we did above will be included it in the
		// ScopeSpans reference.
		if !iOk {
			rs.ScopeSpans = append(rs.ScopeSpans, scopeSpan)
		}
	}

	// Transform the categorized map into a slice
	rss := make([]*types.ResourceSpans, 0, resources)
	for _, rs := range rsm {
		rss = append(rss, rs)
	}
	return rss
}

// span transforms an OpenTelemetry Span into an OTLP span.
func span(sd tracesdk.ReadOnlySpan) *types.Span {
	if sd == nil {
		return nil
	}

	tid := sd.SpanContext().TraceID()
	sid := sd.SpanContext().SpanID()

	s := &types.Span{
		TraceId:                types.NewTraceID(tid),
		SpanId:                 types.NewSpanID(sid),
		TraceState:             sd.SpanContext().TraceState().String(),
		Status:                 status(sd.Status().Code, sd.Status().Description),
		StartTimeUnixNano:      uint64(max(0, sd.StartTime().UnixNano())), // nolint:gosec // Overflow checked.
		EndTimeUnixNano:        uint64(max(0, sd.EndTime().UnixNano())),   // nolint:gosec // Overflow checked.
		Links:                  links(sd.Links()),
		Kind:                   spanKind(sd.SpanKind()),
		Name:                   sd.Name(),
		Attributes:             KeyValues(sd.Attributes()),
		Events:                 spanEvents(sd.Events()),
		DroppedAttributesCount: clampUint32(sd.DroppedAttributes()),
		DroppedEventsCount:     clampUint32(sd.DroppedEvents()),
		DroppedLinksCount:      clampUint32(sd.DroppedLinks()),
	}

	if psid := sd.Parent().SpanID(); psid.IsValid() {
		s.ParentSpanId = types.NewSpanID(psid)
	}
	s.Flags = buildSpanFlags(sd.Parent())

	return s
}

func clampUint32(v int) uint32 {
	if v < 0 {
		return 0
	}
	if int64(v) > math.MaxUint32 {
		return math.MaxUint32
	}
	return uint32(v) // nolint: gosec  // Overflow/Underflow checked.
}

// status transform a span code and message into an OTLP span status.
func status(status codes.Code, message string) *types.Status {
	var c types.Status_StatusCode
	switch status {
	case codes.Ok:
		c = types.Status_STATUS_CODE_OK
	case codes.Error:
		c = types.Status_STATUS_CODE_ERROR
	default:
		c = types.Status_STATUS_CODE_UNSET
	}
	return &types.Status{
		Code:    c,
		Message: message,
	}
}

// links transforms span Links to OTLP span links.
func links(links []tracesdk.Link) []*types.Span_Link {
	if len(links) == 0 {
		return nil
	}

	sls := make([]*types.Span_Link, 0, len(links))
	for _, otLink := range links {
		// This redefinition is necessary to prevent otLink.*ID[:] copies
		// being reused -- in short we need a new otLink per iteration.
		otLink := otLink

		tid := otLink.SpanContext.TraceID()
		sid := otLink.SpanContext.SpanID()

		flags := buildSpanFlags(otLink.SpanContext)

		sl := &types.Span_Link{
			TraceId:                types.NewTraceID(tid),
			SpanId:                 types.NewSpanID(sid),
			Attributes:             KeyValues(otLink.Attributes),
			DroppedAttributesCount: clampUint32(otLink.DroppedAttributeCount),
			Flags:                  flags,
		}
		sls = append(sls, sl)
	}
	return sls
}

func buildSpanFlags(sc trace.SpanContext) uint32 {
	flags := types.SpanFlags_SPAN_FLAGS_CONTEXT_HAS_IS_REMOTE_MASK
	if sc.IsRemote() {
		flags |= types.SpanFlags_SPAN_FLAGS_CONTEXT_IS_REMOTE_MASK
	}

	return uint32(flags) // nolint:gosec // Flags is a bitmask and can't be negative
}

// spanEvents transforms span Events to an OTLP span events.
func spanEvents(es []tracesdk.Event) []*types.Span_Event {
	if len(es) == 0 {
		return nil
	}

	events := make([]*types.Span_Event, len(es))
	// Transform message events
	for i := 0; i < len(es); i++ {
		events[i] = &types.Span_Event{
			Name:                   es[i].Name,
			TimeUnixNano:           uint64(max(0, es[i].Time.UnixNano())), // nolint:gosec // Overflow checked.
			Attributes:             KeyValues(es[i].Attributes),
			DroppedAttributesCount: clampUint32(es[i].DroppedAttributeCount),
		}
	}
	return events
}

// spanKind transforms a SpanKind to an OTLP span kind.
func spanKind(kind trace.SpanKind) types.Span_SpanKind {
	switch kind {
	case trace.SpanKindInternal:
		return types.Span_SPAN_KIND_INTERNAL
	case trace.SpanKindClient:
		return types.Span_SPAN_KIND_CLIENT
	case trace.SpanKindServer:
		return types.Span_SPAN_KIND_SERVER
	case trace.SpanKindProducer:
		return types.Span_SPAN_KIND_PRODUCER
	case trace.SpanKindConsumer:
		return types.Span_SPAN_KIND_CONSUMER
	default:
		return types.Span_SPAN_KIND_UNSPECIFIED
	}
}
