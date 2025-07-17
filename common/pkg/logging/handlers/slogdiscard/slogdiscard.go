// Package slogdiscard provides a slog.Handler implementation that discards all log records.
// Useful for disabling logging in tests or specific environments.
package slogdiscard

import (
	"context"
	"log/slog"
)

// DiscardHandler is a slog.Handler that ignores all log records and attributes.
type DiscardHandler struct{}

// NewDiscardLogger returns a *slog.Logger that discards all log records.
func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHandler())
}

// NewDiscardHandler creates and returns a new DiscardHandler instance.
func NewDiscardHandler() *DiscardHandler {
	return &DiscardHandler{}
}

// Handle implements slog.Handler and discards the log record.
func (h *DiscardHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

// WithAttrs implements slog.Handler and returns the same DiscardHandler (no-op).
func (h *DiscardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

// WithGroup implements slog.Handler and returns the same DiscardHandler (no-op).
func (h *DiscardHandler) WithGroup(_ string) slog.Handler {
	return h
}

// Enabled implements slog.Handler and always returns false, disabling all log levels.
func (h *DiscardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}
