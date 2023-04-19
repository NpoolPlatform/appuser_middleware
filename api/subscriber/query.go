package subscriber

import (
	"context"

	subscriber1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/subscriber"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetSubscriber(ctx context.Context, in *npool.GetSubscriberRequest) (*npool.GetSubscriberResponse, error) {
	handler, err := subscriber1.NewHandler(
		ctx,
		subscriber1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriber",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriberResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetSubscriber(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriber",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriberResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSubscriberResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSubscriberes(ctx context.Context, in *npool.GetSubscriberesRequest) (*npool.GetSubscriberesResponse, error) {
	handler, err := subscriber1.NewHandler(
		ctx,
		subscriber1.WithConds(in.GetConds()),
		subscriber1.WithOffset(in.GetOffset()),
		subscriber1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriberes",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriberesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetSubscriberes(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSubscriberes",
			"In", in,
			"Error", err,
		)
		return &npool.GetSubscriberesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetSubscriberesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
