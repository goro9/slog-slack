package main

import (
	"fmt"
	"time"

	"log/slog"

	slogslack "github.com/goro9/slog-slack/v2"
)

func main() {
	webhook := "https://hooks.slack.com/services/xxx/yyy/zzz"
	channel := "alerts"

	logger := slog.New(slogslack.Option{Level: slog.LevelDebug, WebhookURL: webhook, Channel: channel}.NewSlackHandler())
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
