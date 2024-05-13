package app

import (
	"log/slog"

	appGrpc "github.com/andrei-pokhila/insights-dashboards/src/app/grpc"
)

type App struct {
	GRPCServer *appGrpc.App
}

func NewApp(logger *slog.Logger, port int) *App {
	grpcApp := appGrpc.NewApp(logger, port)

	return &App{
		GRPCServer: grpcApp,
	}
}
