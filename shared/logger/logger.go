package logger

import (
	"log/slog"
	"os"
)

// New()
func New(service string) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	return slog.New(handler).With(
		slog.String("service", service),
	)
}

// Get()

// WithService()
