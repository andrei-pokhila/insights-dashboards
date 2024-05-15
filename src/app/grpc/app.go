package grpc

import (
	"fmt"
	"log/slog"
	"net"

	basicDiffs "github.com/andrei-pokhila/insights-dashboards/src/grpc/basic_diffs"
	"google.golang.org/grpc"
)

type App struct {
	logger     *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewApp(logger *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()
	basicDiffs.Register(gRPCServer, logger)

	return &App{
		logger:     logger,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustStart() {
	if err := a.Start(); err != nil {
		panic(err)
	}
}

func (a *App) Start() error {
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", a.port))
	if err != nil {
		a.logger.Error(fmt.Sprint(err))

		return fmt.Errorf("failed to listen: %v", err)
	}

	a.logger.Info("gRPC server started", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		a.logger.Error(fmt.Sprint(err))

		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
