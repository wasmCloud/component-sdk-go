package wasilog

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	slogcommon "github.com/samber/slog-common"
	"go.wasmcloud.dev/component/gen/wasi/logging/logging"
)

// DefaultLogger is the default implementation that adapts the [wasi:logging] interface to a [slog.Handler].
//
// [wasi:logging]: https://github.com/WebAssembly/wasi-logging
var DefaultLogger = slog.New(DefaultOptions().NewHandler())

// ContextLogger returns a [DefaultLogger] implementation that has an additional "wasi-context" [slog.Attr] attached to it.
func ContextLogger(wasiContext string) *slog.Logger {
	return DefaultLogger.With(ContextAttr(wasiContext))
}

type contextKey string

func (k contextKey) String() string {
	return string(k)
}

// ContextKey is a predefined key used to track the wasi
const ContextKey = contextKey("wasi-context")

func ContextAttr(name string) slog.Attr {
	return slog.String(string(ContextKey), name)
}

type wasmLoggerFunc func(level logging.Level, context string, message string)

// WasiLoggingOption represents the available options for customizing the [WebassemblyHandler].
type WasiLoggingOption struct {
	// required: log function
	LoggerFunc wasmLoggerFunc
	// log level (default: info)
	Level slog.Leveler

	// optional: fetch attributes from context
	AttrFromContext []func(ctx context.Context) []slog.Attr

	// optional: replace attributes
	ReplaceAttr func(groups []string, a slog.Attr) slog.Attr
}

// WebassemblyHandler implements the [slog.Handler] interface to adapt [slog] to wasi:logging.
type WebassemblyHandler struct {
	option WasiLoggingOption
	attrs  []slog.Attr
	groups []string
}

var _ slog.Handler = (*WebassemblyHandler)(nil)

func wasiLevel(level slog.Level) logging.Level {
	switch level {
	case slog.LevelDebug:
		return logging.LevelDebug
	case slog.LevelInfo:
		return logging.LevelInfo
	case slog.LevelWarn:
		return logging.LevelWarn
	case slog.LevelError:
		return logging.LevelError
	default:
		return logging.LevelDebug
	}
}

func unrollGroups(attrs []slog.Attr) []slog.Attr {
	var output []slog.Attr
	for i := range attrs {
		attr := attrs[i]
		switch attr.Value.Kind() {
		case slog.KindGroup:
			output = append(output, unrollGroups(attr.Value.Group())...)
		case slog.KindLogValuer:
			output = append(output, slog.String(attr.Key, attr.Value.LogValuer().LogValue().String()))
		default:
			output = append(output, attr)
		}
	}
	return output
}

func wasiConverter(replaceAttr func(groups []string, a slog.Attr) slog.Attr, loggerAttr []slog.Attr, groups []string, record *slog.Record) (string, string) {
	attrs := slogcommon.AppendRecordAttrsToAttrs(loggerAttr, groups, record)
	attrs = slogcommon.ReplaceAttrs(replaceAttr, groups, attrs...)
	attrs = slogcommon.RemoveEmptyAttrs(attrs)
	attrs = unrollGroups(attrs)

	extra := slogcommon.AttrsToString(attrs...)
	context, ok := extra[string(ContextKey)]
	if ok {
		// the context key is moved to the 'Context' field in wasi:logging. remove it from the log message.
		delete(extra, string(ContextKey))
	}

	var formattedAttrs []string
	for k, v := range extra {
		formattedAttrs = append(formattedAttrs, fmt.Sprintf("%s=%q", k, v))
	}
	formattedAttrs = append(formattedAttrs, record.Message)

	return strings.Join(formattedAttrs, " "), context
}

// DefaultOptions represents the default set of values used in [WasiLoggingOption] that are used in setting up [DefaultLogger].
func DefaultOptions() WasiLoggingOption {
	return WasiLoggingOption{
		LoggerFunc: logging.Log,
		Level:      slog.LevelInfo,
		AttrFromContext: []func(ctx context.Context) []slog.Attr{
			func(ctx context.Context) []slog.Attr {
				if contextName, ok := ctx.Value(ContextKey).(string); ok {
					return []slog.Attr{slog.String(string(ContextKey), string(contextName))}
				}
				return nil
			},
		},
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Make so groups become a prefix of the key
			// Ex: groups = ["a", "b"], key = "c" => "a.b.c"
			if len(groups) == 0 {
				return a
			}

			a.Key = strings.Join(groups, ".") + "." + a.Key
			return a
		},
	}
}

// NewHandler is used to instantiate a new instance of a [WebassemblyHandler] that implements the [slog.Handler] interface.
func (o WasiLoggingOption) NewHandler() slog.Handler {
	return &WebassemblyHandler{
		option: o,
	}
}

// Enabled reports whether the handler handles records at the given level. The handler ignores records whose level is lower.
func (h *WebassemblyHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.option.Level.Level()
}

// Handle formats its argument [slog.Record] using the provided [context.Context] into a wasi:logging compatible output.
func (h *WebassemblyHandler) Handle(ctx context.Context, record slog.Record) error {
	fromContext := slogcommon.ContextExtractor(ctx, h.option.AttrFromContext)
	message, logContext := wasiConverter(h.option.ReplaceAttr, append(h.attrs, fromContext...), h.groups, &record)

	h.option.LoggerFunc(wasiLevel(record.Level), logContext, message)

	return nil
}

// WithAttrs returns a new [WebAssemblyHandler] whose attributes consists
// of h's attributes followed by attrs.
func (h *WebassemblyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &WebassemblyHandler{
		option: h.option,
		attrs:  slogcommon.AppendAttrsToGroup(h.groups, h.attrs, attrs...),
		groups: h.groups,
	}
}

// WithGroups returns a new [WebassemblyHandler] where the attributes are
// grouped under a common name.
func (h *WebassemblyHandler) WithGroup(name string) slog.Handler {
	// https://cs.opensource.google/go/x/exp/+/46b07846:slog/handler.go;l=247
	if name == "" {
		return h
	}

	return &WebassemblyHandler{
		option: h.option,
		attrs:  h.attrs,
		groups: append(h.groups, name),
	}
}
