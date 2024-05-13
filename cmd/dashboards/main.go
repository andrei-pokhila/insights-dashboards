package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/andrei-pokhila/insights-dashboards/src/app"
	"github.com/andrei-pokhila/insights-dashboards/src/config"
)

func main() {
	cfg := config.NewConfig()

	logger := setupLogger(cfg)

	logger.Info("started")

	application := app.NewApp(logger, cfg.GrpcPort)
	go application.GRPCServer.MustStart()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	s := <-stop
	logger.Info("stopping application", slog.String("signal", s.String()))

	application.GRPCServer.Stop()
	logger.Info("stopped")
}

func setupLogger(cfg *config.Config) *slog.Logger {
	var logger *slog.Logger

	switch cfg.Debug {
	case true:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	case false:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	return logger
}
