package loggerSwitch

import (
	"log/slog"
	"os"
)

// NewLogger создаёт новый slog.Logger с уровнем, заданным строкой levelStr
func NewLogger(levelStr string) *slog.Logger {
	var logLevel slog.Level
	switch levelStr {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	return slog.New(slog.NewTextHandler(os.Stderr, opts))
}
