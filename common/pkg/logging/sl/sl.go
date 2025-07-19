package sl

import (
	"errors"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/10Narratives/ready-to-do/common/pkg/logging/handlers/slogdiscard"
	"github.com/10Narratives/ready-to-do/common/pkg/logging/handlers/slogpretty"
	"github.com/natefinch/lumberjack"
)

type LoggerOptions struct {
	level  slog.Level
	format string
	output string
}

func defaultOptions() *LoggerOptions {
	return &LoggerOptions{
		level:  slog.LevelError,
		format: "json",
		output: "stdout",
	}
}

type LoggerOption func(*LoggerOptions)

func WithLevel(level string) LoggerOption {
	return func(lo *LoggerOptions) {
		lo.level = parseLevel(level)
	}
}

func parseLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		panic("unsupported log level: " + level)
	}
}

func WithFormat(format string) LoggerOption {
	return func(lo *LoggerOptions) {
		lo.format = format
	}
}

func WithOutput(output string) LoggerOption {
	return func(lo *LoggerOptions) {
		lo.output = output
	}
}

func New(opts ...LoggerOption) (*slog.Logger, error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	output, err := createOutput(options)
	if err != nil {
		return nil, err
	}

	handler, err := createHandler(options.format, output, options.level)
	if err != nil {
		return nil, err
	}

	return slog.New(handler), nil
}

func createOutput(opts *LoggerOptions) (io.Writer, error) {
	if opts.output == "stdout" {
		return os.Stdout, nil
	}

	dir := filepath.Dir(opts.output)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	return &lumberjack.Logger{
		Filename:  opts.output,
		LocalTime: true,
	}, nil
}

func createHandler(format string, output io.Writer, level slog.Level) (slog.Handler, error) {
	opts := &slog.HandlerOptions{
		Level: level,
	}

	switch format {
	case "json":
		return slog.NewJSONHandler(output, opts), nil
	case "pretty":
		return slogpretty.NewPrettyLogger(&slogpretty.PrettyHandlerOptions{
			SlogOpts: opts,
		}, output).Handler(), nil
	case "plain":
		return slog.NewTextHandler(output, opts), nil
	case "discard":
		return slogdiscard.NewDiscardLogger().Handler(), nil
	default:
		return nil, errors.New("unsupported format")
	}
}
