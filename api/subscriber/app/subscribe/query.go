package appsubscribe

import (
	"context"

	appsubscribe1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/subscriber/app/subscribe"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber/app/subscribe"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAppSubscribe(ctx context.Context, in *npool.GetAppSubscribeRequest) (*npool.GetAppSubscribeResponse, error) {
	handler, err := appsubscribe1.NewHandler(
		ctx,
		appsubscribe1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppSubscribe",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppSubscribeResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetAppSubscribe(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppSubscribe",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppSubscribeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAppSubscribeResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppSubscribes(ctx context.Context, in *npool.GetAppSubscribesRequest) (*npool.GetAppSubscribesResponse, error) {
	handler, err := appsubscribe1.NewHandler(
		ctx,
		appsubscribe1.WithConds(in.GetConds()),
		appsubscribe1.WithOffset(in.GetOffset()),
		appsubscribe1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppSubscribes",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppSubscribesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetAppSubscribes(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppSubscribes",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppSubscribesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAppSubscribesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
