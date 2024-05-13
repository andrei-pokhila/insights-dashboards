package basic_diffs

import (
	"context"

	dashboards "github.com/andrei-pokhila/insights-dashboards/gen/go/dashboards"
	"github.com/andrei-pokhila/insights-dashboards/src/dashboards/basic_diffs"
	grpc "google.golang.org/grpc"
)

type serverAPI struct {
	dashboards.UnimplementedBasicDiffsServer
}

func Register(gRPC *grpc.Server) {
	dashboards.RegisterBasicDiffsServer(gRPC, &serverAPI{})
}

func (s *serverAPI) GetFundingRate(
	ctx context.Context,
	in *dashboards.BasicRequest,
) (*dashboards.FundingResponse, error) {
	candles := basic_diffs.GetCandles(in)

	// return nil, status.Errorf(codes.Unimplemented, "method GetFundingRate not implemented")

	return candles, nil
}
