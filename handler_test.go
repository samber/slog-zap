package slogzap

import (
	"log/slog"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestZapHandler_WithGroupSiblings(t *testing.T) {
	core, logs := observer.New(zapcore.DebugLevel)

	root := slog.New(Option{Logger: zap.New(core)}.NewZapHandler())
	parent := root.WithGroup("a").WithGroup("b").WithGroup("c")

	first := parent.WithGroup("d")
	second := parent.WithGroup("e")

	first.Info("first", "k", 1)
	second.Info("second", "k", 2)

	entries := logs.All()
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}

	for i, group := range []string{"d", "e"} {
		got := entries[i].ContextMap()
		inner, ok := got["a"].(map[string]any)["b"].(map[string]any)["c"].(map[string]any)[group]
		if !ok || inner == nil {
			t.Fatalf("record %d lost group %q: %v", i, group, got)
		}
	}
}
