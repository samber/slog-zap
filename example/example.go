package main

import (
	"fmt"
	"time"

	"log/slog"

	slogzap "github.com/samber/slog-zap"
	"go.uber.org/zap"
)

func main() {
	zapLogger, _ := zap.NewProduction()

	logger := slog.New(slogzap.Option{Level: slog.LevelDebug, Logger: zapLogger}.NewZapHandler())
	logger = logger.With("release", "v1.0.0")

	logger.
		With(
			slog.Group("user",
				slog.String("id", "user-123"),
				slog.Time("created_at", time.Now().AddDate(0, 0, -1)),
			),
		).
		With("environment", "dev").
		With("error", fmt.Errorf("an error")).
		Error("A message")
}
