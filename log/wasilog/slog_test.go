package wasilog

import (
	"context"
	"fmt"
	"log/slog"
	"testing"
	_ "unsafe"

	"go.wasmcloud.dev/component/gen/wasi/logging/logging"
)

func TestLogLevelMapping(t *testing.T) {
	tt := map[string]struct {
		wasiLevel logging.Level
		slogLevel slog.Level
	}{
		"debug": {
			wasiLevel: logging.LevelDebug,
			slogLevel: slog.LevelDebug,
		},
		"info": {
			wasiLevel: logging.LevelInfo,
			slogLevel: slog.LevelInfo,
		},
		"warn": {
			wasiLevel: logging.LevelWarn,
			slogLevel: slog.LevelWarn,
		},
		"error": {
			wasiLevel: logging.LevelError,
			slogLevel: slog.LevelError,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			output := func(level logging.Level, _ string, _ string) {
				if level != tc.wasiLevel {
					t.Errorf("expected: %v, got: %v", tc.wasiLevel, level)
				}
			}
			options := DefaultOptions()
			options.LoggerFunc = output
			logger := slog.New(options.NewHandler())
			logger.Log(context.TODO(), tc.slogLevel, "test")
		})
	}
}

func TestDefaultLevel(t *testing.T) {
	// default level is info
	// debug -> info -> warn -> error
	// not all wasi levels are mapped (ex: trace & fatal)

	allLevels := []slog.Level{
		slog.LevelDebug,
		slog.LevelInfo,
		slog.LevelWarn,
		slog.LevelError,
	}

	for i, defaultLevel := range allLevels {
		for j, emitLevel := range allLevels {
			t.Run(fmt.Sprintf("%s_%s", defaultLevel, emitLevel), func(t *testing.T) {
				output := func(level logging.Level, _ string, _ string) {
					if i > j {
						t.Errorf("Emitted log level %v is lower than default log level %v", emitLevel, defaultLevel)
					}
				}
				options := DefaultOptions()
				options.LoggerFunc = output
				options.Level = defaultLevel
				logger := slog.New(options.NewHandler())
				logger.Log(context.TODO(), emitLevel, "test")
			})
		}
	}
}

func TestContextLift(t *testing.T) {
	// preferred
	t.Run("With", func(t *testing.T) {
		output := func(_ logging.Level, context string, _ string) {
			if context != "from_with" {
				t.Errorf("expected: %v, got: %v", "from_context", context)
			}
		}
		options := DefaultOptions()
		options.LoggerFunc = output
		logger := slog.New(options.NewHandler())
		logger = logger.With(ContextAttr("from_with"))
		logger.Log(context.TODO(), slog.LevelInfo, "test")
	})
	t.Run("attribute", func(t *testing.T) {
		output := func(_ logging.Level, context string, _ string) {
			if context != "from_attribute" {
				t.Errorf("expected: %v, got: %v", "from_attribute", context)
			}
		}
		options := DefaultOptions()
		options.LoggerFunc = output
		logger := slog.New(options.NewHandler())
		logger.Log(context.Background(), slog.LevelInfo, "test", ContextAttr("from_attribute"))
		logger.Log(context.Background(), slog.LevelInfo, "test", ContextKey.String(), "from_attribute")
		logger.Log(context.Background(), slog.LevelInfo, "test", slog.String(ContextKey.String(), "from_attribute"))
		logger.Log(context.Background(), slog.LevelInfo, "test", slog.Any(ContextKey.String(), "from_attribute"))
	})

	t.Run("context", func(t *testing.T) {
		output := func(_ logging.Level, context string, _ string) {
			if context != "from_context" {
				t.Errorf("expected: %v, got: %v", "from_context", context)
			}
		}
		options := DefaultOptions()
		options.LoggerFunc = output
		logger := slog.New(options.NewHandler())
		ctx := context.WithValue(context.Background(), ContextKey, "from_context")
		logger.Log(ctx, slog.LevelInfo, "test")
	})
}

func TestGroups(t *testing.T) {
	tt := map[string]struct {
		cb       func(logger *slog.Logger)
		expected string
	}{
		"inline": {
			cb: func(logger *slog.Logger) {
				logger.Info("test", slog.Group("top", slog.String("bottom", "simple")))
			},
			expected: `top.bottom="simple" test`,
		},
		"nested": {
			cb: func(logger *slog.Logger) {
				g := logger.With("bottom", "nested").WithGroup("top")
				g.Info("test")
			},
			expected: `top.bottom="nested" test`,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			options := DefaultOptions()
			options.LoggerFunc = func(_ logging.Level, _ string, msg string) {
				if want, got := tc.expected, msg; got != want {
					t.Errorf("expected: %v, got: %v", want, got)
				}
			}

			logger := slog.New(options.NewHandler())
			tc.cb(logger)
		})
	}

	logger := slog.Default()
	logger.Info("test", slog.Group("top", slog.String("bottom", "simple")))
	g := logger.With("bottom", "nested").WithGroup("top")
	g.Info("test")
	logger = logger.With("id", "123")
	parserLogger := logger.WithGroup("parser")
	parserLogger.Info("parsing", "file", "test.txt")
}

type Token string

// LogValue implements slog.LogValuer.
// It avoids revealing the token.
func (Token) LogValue() slog.Value {
	return slog.StringValue("REDACTED_TOKEN")
}

func TestLogValueMask(t *testing.T) {
	output := func(_ logging.Level, _ string, msg string) {
		if want, got := `token="REDACTED_TOKEN" test`, msg; got != want {
			t.Errorf("expected: %v, got: %v", want, got)
		}
	}
	options := DefaultOptions()
	options.LoggerFunc = output
	logger := slog.New(options.NewHandler())
	logger.Info("test", "token", Token("launch-the-nukes-code"))
}
