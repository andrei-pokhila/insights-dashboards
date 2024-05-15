package basic_diffs

import (
	"context"
	"log/slog"

	dashboards "github.com/andrei-pokhila/insights-dashboards/gen/go/dashboards"
	"github.com/andrei-pokhila/insights-dashboards/src/dashboards/basic_diffs"
	grpc "google.golang.org/grpc"
)

type serverAPI struct {
	dashboards.UnimplementedBasicDiffsServer
	logger *slog.Logger
}

func Register(gRPC *grpc.Server, logger *slog.Logger) {
	dashboards.RegisterBasicDiffsServer(gRPC, &serverAPI{logger: logger})
}

func (s *serverAPI) GetFundingRate(
	ctx context.Context,
	in *dashboards.BasicRequest,
) (*dashboards.FundingResponse, error) {
	candles := basic_diffs.GetCandles(in)

	s.logger.Info("GetFundingRate")

	// return nil, status.Errorf(codes.Unimplemented, "method GetFundingRate not implemented")

	return candles, nil
}
