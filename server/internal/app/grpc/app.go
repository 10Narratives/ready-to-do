package grpcapp

import (
	"context"

	transportcfg "github.com/10Narratives/ready-to-do/server/internal/config/transport"
)

type App struct{}

func New(cfg *transportcfg.GRPC) (*App, error) {
	return &App{}, nil
}

func (a *App) Run() error {
	return nil
}

func (a *App) Stop(ctx context.Context) error {
	return nil
}
