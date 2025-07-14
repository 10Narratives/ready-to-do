package app

import (
	"context"
	"log/slog"

	gwapp "github.com/10Narratives/ready-to-do/server/internal/app/gateway"
	grpcapp "github.com/10Narratives/ready-to-do/server/internal/app/grpc"
	pgapp "github.com/10Narratives/ready-to-do/server/internal/app/postgres"
	"github.com/10Narratives/ready-to-do/server/internal/config"
	"github.com/10Narratives/ready-to-do/server/internal/lib/logging/sl"
)

type App struct {
	GRPCApp *grpcapp.App
	GWApp   *gwapp.App
	PGApp   *pgapp.App

	Logger *slog.Logger
}

func New(cfg *config.Config) *App {
	log := sl.MustLogger(
		sl.WithFormat(cfg.Logging.Format),
		sl.WithOutput(cfg.Logging.Output),
		sl.WithLevel(cfg.Logging.Level),
	)

	var transportCfg = &cfg.Transport
	gRPCApp := grpcapp.New(transportCfg)
	gwApp := gwapp.New(transportCfg)

	var pgAppCfg = &cfg.Database
	pgApp := pgapp.New(pgAppCfg)

	return &App{
		GRPCApp: gRPCApp,
		GWApp:   gwApp,
		PGApp:   pgApp,
		Logger:  log,
	}
}

func (a *App) Stop(ctx context.Context) error {
	return nil
}
