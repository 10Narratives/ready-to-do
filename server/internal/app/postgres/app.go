package pgapp

import (
	"context"

	databasecfg "github.com/10Narratives/ready-to-do/server/internal/config/database"
)

type App struct {
}

func New(cfg *databasecfg.Database) (*App, error) {
	return &App{}, nil
}

func (a *App) Run() error {
	return nil
}

func (a *App) Stop(ctx context.Context) error {
	return nil
}
