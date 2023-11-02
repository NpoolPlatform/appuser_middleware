package appsubscribe

import (
	"context"

	appsubscribe1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/subscriber/app/subscribe"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber/app/subscribe"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAppSubscribe(ctx context.Context, in *npool.CreateAppSubscribeRequest) (*npool.CreateAppSubscribeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateAppSubscribe",
			"In", in,
		)
		return &npool.CreateAppSubscribeResponse{}, status.Error(codes.Aborted, "invalid req")
	}
	handler, err := appsubscribe1.NewHandler(
		ctx,
		appsubscribe1.WithEntID(req.EntID, false),
		appsubscribe1.WithAppID(req.AppID, true),
		appsubscribe1.WithSubscribeAppID(req.SubscribeAppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppSubscribe",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAppSubscribeResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateAppSubscribe(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppSubscribe",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAppSubscribeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateAppSubscribeResponse{
		Info: info,
	}, nil
}
