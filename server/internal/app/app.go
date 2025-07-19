package app

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/10Narratives/ready-to-do/common/pkg/logging/sl"
	grpcapp "github.com/10Narratives/ready-to-do/server/internal/app/grpc"
	pgapp "github.com/10Narratives/ready-to-do/server/internal/app/postgres"
	"github.com/10Narratives/ready-to-do/server/internal/config"
)

type App struct {
	GRPCApp *grpcapp.App
	PGApp   *pgapp.App

	Logger *slog.Logger
}

func New(cfg *config.Config) (*App, error) {
	logger, err := sl.New(
		sl.WithLevel(cfg.Logging.Level),
		sl.WithFormat(cfg.Logging.Format),
		sl.WithOutput(cfg.Logging.Output),
	)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize logger: %s", err.Error())
	}

	grpcApp, err := grpcapp.New(&cfg.Transport.GRPC)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize gRPC component: %s", err.Error())
	}

	pgApp, err := pgapp.New(&cfg.Database)
	if err != nil {
		return nil, fmt.Errorf("cannot initalize postgres component: %s", err.Error())
	}

	return &App{
		GRPCApp: grpcApp,
		PGApp:   pgApp,
		Logger:  logger,
	}, nil
}

func (a *App) Stop(ctx context.Context) error {
	return nil
}
