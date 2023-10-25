package slogzap

import (
	"log/slog"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LogLevels = map[slog.Level]zapcore.Level{
	slog.LevelDebug: zap.DebugLevel,
	slog.LevelInfo:  zap.InfoLevel,
	slog.LevelWarn:  zap.WarnLevel,
	slog.LevelError: zap.ErrorLevel,
}
