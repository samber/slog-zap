package slogzap

import (
	"context"

	"log/slog"

	"go.uber.org/zap"
)

type Option struct {
	// log level (default: debug)
	Level slog.Leveler

	// optional: zap logger (default: zap.L())
	Logger *zap.Logger

	// optional: customize json payload builder
	Converter Converter
}

func (o Option) NewZapHandler() slog.Handler {
	if o.Level == nil {
		o.Level = slog.LevelDebug
	}

	if o.Logger == nil {
		// should be selected lazily ?
		o.Logger = zap.L()
	}

	return &ZapHandler{
		option: o,
		attrs:  []slog.Attr{},
		groups: []string{},
	}
}

var _ slog.Handler = (*ZapHandler)(nil)

type ZapHandler struct {
	option Option
	attrs  []slog.Attr
	groups []string
}

func (h *ZapHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.option.Level.Level()
}

func (h *ZapHandler) Handle(ctx context.Context, record slog.Record) error {
	converter := DefaultConverter
	if h.option.Converter != nil {
		converter = h.option.Converter
	}

	level := levelMap[record.Level]
	fields := converter(h.attrs, &record)

	h.option.Logger.Log(level, record.Message, fields...)

	return nil
}

func (h *ZapHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ZapHandler{
		option: h.option,
		attrs:  appendAttrsToGroup(h.groups, h.attrs, attrs),
		groups: h.groups,
	}
}

func (h *ZapHandler) WithGroup(name string) slog.Handler {
	return &ZapHandler{
		option: h.option,
		attrs:  h.attrs,
		groups: append(h.groups, name),
	}
}
