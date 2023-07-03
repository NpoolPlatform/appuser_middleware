package appsubscribe

import (
	"context"

	appsubscribe1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/subscriber/app/subscribe"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber/app/subscribe"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteAppSubscribe(ctx context.Context, in *npool.DeleteAppSubscribeRequest) (*npool.DeleteAppSubscribeResponse, error) {
	req := in.GetInfo()
	handler, err := appsubscribe1.NewHandler(
		ctx,
		appsubscribe1.WithID(req.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAppSubscribe",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAppSubscribeResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteAppSubscribe(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAppSubscribe",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAppSubscribeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteAppSubscribeResponse{
		Info: info,
	}, nil
}
