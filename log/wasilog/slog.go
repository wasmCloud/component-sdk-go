package wasilog

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	slogcommon "github.com/samber/slog-common"
	"go.wasmcloud.dev/component/gen/wasi/logging/logging"
)

type WasiLoggingOption struct {
	// log level (default: debug)
	Level slog.Leveler

	// optional: fetch attributes from context
	AttrFromContext []func(ctx context.Context) []slog.Attr

	ReplaceAttr func(groups []string, a slog.Attr) slog.Attr
}

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

func wasiConverter(replaceAttr func(groups []string, a slog.Attr) slog.Attr, loggerAttr []slog.Attr, groups []string, record *slog.Record) (string, string) {
	attrs := slogcommon.AppendRecordAttrsToAttrs(loggerAttr, groups, record)

	attrs = slogcommon.ReplaceAttrs(replaceAttr, []string{}, attrs...)
	attrs = slogcommon.RemoveEmptyAttrs(attrs)

	extra := slogcommon.AttrsToString(attrs...)

	context, ok := extra["context"]
	if ok {
		delete(extra, "context")
	}

	var formattedAttrs []string
	for k, v := range extra {
		formattedAttrs = append(formattedAttrs, fmt.Sprintf("%s=%q", k, v))
	}
	formattedAttrs = append(formattedAttrs, record.Message)

	return strings.Join(formattedAttrs, " "), context
}

func DefaultOptions() WasiLoggingOption {
	return WasiLoggingOption{
		Level:           slog.LevelDebug,
		AttrFromContext: []func(ctx context.Context) []slog.Attr{},
	}
}

func (o WasiLoggingOption) NewHandler() slog.Handler {
	return &WebassemblyHandler{
		option: o,
	}
}

func (h *WebassemblyHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.option.Level.Level()
}

func (h *WebassemblyHandler) Handle(ctx context.Context, record slog.Record) error {
	fromContext := slogcommon.ContextExtractor(ctx, h.option.AttrFromContext)
	message, logContext := wasiConverter(h.option.ReplaceAttr, append(h.attrs, fromContext...), h.groups, &record)

	logging.Log(wasiLevel(record.Level), logContext, message)

	return nil
}

func (h *WebassemblyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &WebassemblyHandler{
		option: h.option,
		attrs:  slogcommon.AppendAttrsToGroup(h.groups, h.attrs, attrs...),
		groups: h.groups,
	}
}

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
