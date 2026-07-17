package logger

import "log/slog"

// service

// correlation_id
func WithCorrelationID(log *slog.Logger, correaltionID string) *slog.Logger {
	return log.With(slog.String("correlation_id", correaltionID))
}

// request_id
func WithRequestID(log *slog.Logger, requestID string) *slog.Logger {
	return log.With(slog.String("request_id", requestID))
}
