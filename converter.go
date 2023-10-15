package slogzap

import (
	"log/slog"

	slogcommon "github.com/samber/slog-common"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SourceKey = "source"
var ErrorKeys = []string{"error", "err"}

type Converter func(addSource bool, replaceAttr func(groups []string, a slog.Attr) slog.Attr, loggerAttr []slog.Attr, groups []string, record *slog.Record) []zapcore.Field

func DefaultConverter(addSource bool, replaceAttr func(groups []string, a slog.Attr) slog.Attr, loggerAttr []slog.Attr, groups []string, record *slog.Record) []zapcore.Field {
	// aggregate all attributes
	attrs := slogcommon.AppendRecordAttrsToAttrs(loggerAttr, groups, record)

	// developer formatters
	attrs = slogcommon.ReplaceError(attrs, ErrorKeys...)
	if addSource {
		attrs = append(attrs, slogcommon.Source(SourceKey, record))
	}
	attrs = slogcommon.ReplaceAttrs(replaceAttr, []string{}, attrs...)

	// handler formatter
	fields := slogcommon.AttrsToMap(attrs...)

	output := []zapcore.Field{}
	for k, v := range fields {
		output = append(output, zap.Any(k, v))
	}

	return output
}
