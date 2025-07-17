package slogpretty

import (
	"context"
	"encoding/json"
	"io"
	stdLog "log"
	"log/slog"
	"os"

	"github.com/fatih/color"
)

// PrettyHandlerOptions holds options for configuring the PrettyHandler.
type PrettyHandlerOptions struct {
	SlogOpts *slog.HandlerOptions // Options for the underlying slog handler.
}

// PrettyHandler is a slog.Handler that outputs logs in a human-friendly, colorized format.
type PrettyHandler struct {
	slog.Handler                // Embedded slog.Handler for base functionality.
	l            *stdLog.Logger // Standard logger used for output.
	attrs        []slog.Attr    // Additional attributes to include in every log record.
}

// NewPrettyHandler creates a new PrettyHandler that writes pretty, colorized logs to the given writer.
func (opts PrettyHandlerOptions) NewPrettyHandler(
	out io.Writer,
) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, opts.SlogOpts),
		l:       stdLog.New(out, "", 0),
	}

	return h
}

// PrettyHandler implements slog.Handler to provide human-friendly, colorized log output.
//
// It embeds a base slog.Handler (typically a JSON handler for attribute processing), and uses a standard
// library logger for output. PrettyHandler formats log records with colored level indicators, timestamps,
// and pretty-printed attributes for improved readability in terminal environments.
//
// Use PrettyHandlerOptions.NewPrettyHandler to construct a PrettyHandler with custom slog.HandlerOptions
// and output destination.
//
// Example usage:
//
//	ph := slogpretty.PrettyHandlerOptions{SlogOpts: &slog.HandlerOptions{Level: slog.LevelInfo}}.
//		NewPrettyHandler(os.Stdout)
//	logger := slog.New(ph)
//
//	logger.Info("Hello, world!", slog.String("foo", "bar"))
func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())

	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	for _, a := range h.attrs {
		fields[a.Key] = a.Value.Any()
	}

	var b []byte
	var err error

	if len(fields) > 0 {
		b, err = json.MarshalIndent(fields, "", "  ")
		if err != nil {
			return err
		}
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(
		timeStr,
		level,
		msg,
		color.WhiteString(string(b)),
	)

	return nil
}

// PrettyHandler is a slog.Handler implementation that outputs log records in a human-friendly,
// colorized format suitable for development and local debugging. It prints log level, timestamp,
// message, and structured fields in a readable way.
//
// Example usage:
//
//	logger := NewPrettyLogger(nil)
//	logger.Info("Hello, world!", slog.String("foo", "bar"))
//
// Handle implements slog.Handler. It formats and prints the log record to the configured output.
// Log levels are colorized, and structured fields are pretty-printed as indented JSON.
func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &PrettyHandler{
		Handler: h.Handler,
		l:       h.l,
		attrs:   attrs,
	}
}

// WithAttrs returns a new PrettyHandler with the provided attributes.
// It implements slog.Handler.
func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	return &PrettyHandler{
		Handler: h.Handler.WithGroup(name),
		l:       h.l,
	}
}

// PrettyHandler is a slog.Handler implementation that outputs log records in a human-friendly,
// colorized format suitable for development and local debugging. It prints log level, timestamp,
// message, and structured fields in a readable way.
//
// Example usage:
//
//	logger := NewPrettyLogger(nil)
//	logger.Info("Hello, world!", slog.String("foo", "bar"))
//
// Handle implements slog.Handler. It formats and prints the log record to the configured output.
// Log levels are colorized, and structured fields are pretty-printed as indented JSON.
//
// WithAttrs returns a new PrettyHandler with the provided attributes. It implements slog.Handler.
//
// WithGroup returns a new PrettyHandler with the provided group name. It implements slog.Handler.
func NewPrettyLogger(opts *PrettyHandlerOptions, out ...io.Writer) *slog.Logger {
	output := io.Writer(os.Stdout)
	if len(out) > 0 {
		output = out[0]
	}

	if opts == nil {
		opts = &PrettyHandlerOptions{
			SlogOpts: &slog.HandlerOptions{
				Level: slog.LevelInfo,
			},
		}
	}

	handler := opts.NewPrettyHandler(output)
	return slog.New(handler)
}
