package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/10Narratives/ready-to-do/server/internal/app"
	"github.com/10Narratives/ready-to-do/server/internal/config"
)

func main() {
	application := app.New(config.MustLoad())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errChan := make(chan error, 3)

	go func() {
		if err := application.PGApp.Run(); err != nil {
			errChan <- fmt.Errorf("postgres error: %w", err)
		}
	}()

	go func() {
		if err := application.GRPCApp.Run(); err != nil {
			errChan <- fmt.Errorf("gRPC error: %w", err)
		}
	}()

	go func() {
		if err := application.GWApp.Run(); err != nil {
			errChan <- fmt.Errorf("gateway error: %w", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-stop:
		application.Logger.Info("shutting down gracefully...")

		shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 30*time.Second)
		defer shutdownCancel()

		var wg sync.WaitGroup
		wg.Add(3)

		go func() {
			defer wg.Done()
			if err := application.GWApp.Stop(shutdownCtx); err != nil {
				application.Logger.Error("gateway shutdown error", slog.String("error", err.Error()))
			}
		}()

		go func() {
			defer wg.Done()
			if err := application.GRPCApp.Stop(shutdownCtx); err != nil {
				application.Logger.Error("gRPC shutdown error", slog.String("error", err.Error()))
			}
		}()

		go func() {
			defer wg.Done()
			if err := application.PGApp.Stop(shutdownCtx); err != nil {
				application.Logger.Error("postgres shutdown error", slog.String("error", err.Error()))
			}
		}()

		wg.Wait()
		application.Logger.Info("shutdown completed")

	case err := <-errChan:
		application.Logger.Error("critical component failed", slog.String("error", err.Error()))
		cancel()

		if shutdownErr := application.Stop(context.Background()); shutdownErr != nil {
			application.Logger.Error("shutdown error", slog.String("error", shutdownErr.Error()))
		}
		os.Exit(1)
	}
}
