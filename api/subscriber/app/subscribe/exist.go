package appsubscribe

import (
	"context"

	appsubscribe1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/subscriber/app/subscribe"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber/app/subscribe"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistAppSubscribeConds(ctx context.Context, in *npool.ExistAppSubscribeCondsRequest) (*npool.ExistAppSubscribeCondsResponse, error) {
	handler, err := appsubscribe1.NewHandler(
		ctx,
		appsubscribe1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAppSubscribeConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAppSubscribeCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	exist, err := handler.ExistAppSubscribeConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAppSubscribeConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAppSubscribeCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistAppSubscribeCondsResponse{
		Info: exist,
	}, nil
}
