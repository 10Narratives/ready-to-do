package postgresapp

import (
	"context"
	"log/slog"

	databasecfg "github.com/10Narratives/ready-to-do/server/internal/config/database"
	"github.com/10Narratives/ready-to-do/server/internal/lib/logging/sl"
)

type App struct {
	cfg *databasecfg.Config
	log *slog.Logger
}

func New(cfg *databasecfg.Config) *App {
	log := sl.MustLogger(
		sl.WithFormat(cfg.Logging.Format),
		sl.WithOutput(cfg.Logging.Output),
		sl.WithLevel(cfg.Logging.Level),
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
