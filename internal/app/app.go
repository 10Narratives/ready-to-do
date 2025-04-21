package app

import (
	"log/slog"

	grpcapp "github.com/10Narratives/ddb/internal/app/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, port int) *App {
	grpcApp := grpcapp.New(log, nil, port)
	return &App{
		GRPCServer: grpcApp,
	}
}
