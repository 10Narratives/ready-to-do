package grpcapp

import (
	"context"
	"log/slog"

	transportcfg "github.com/10Narratives/ready-to-do/server/internal/config/transport"
	"github.com/10Narratives/ready-to-do/server/internal/lib/logging/sl"
)

type App struct {
	cfg *transportcfg.Config
	log *slog.Logger
}

func New(cfg *transportcfg.Config) *App {
	log := sl.MustLogger(
		sl.WithFormat(cfg.GRPC.Logging.Format),
		sl.WithOutput(cfg.GRPC.Logging.Output),
		sl.WithLevel(cfg.GRPC.Logging.Level),
	)

	return &App{
		cfg: cfg,
		log: log,
	}
}

func (a *App) Run() error {
	return nil
}

func (a *App) Stop(ctx context.Context) error {
	return nil
}
