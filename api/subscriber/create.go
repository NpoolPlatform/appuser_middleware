package subscriber

import (
	"context"

	subscriber1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/subscriber"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateSubscriber(ctx context.Context, in *npool.CreateSubscriberRequest) (*npool.CreateSubscriberResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateSubscriber",
			"In", in,
		)
		return &npool.CreateSubscriberResponse{}, status.Error(codes.Aborted, "Info is empty")
	}
	handler, err := subscriber1.NewHandler(
		ctx,
		subscriber1.WithEntID(req.EntID, false),
		subscriber1.WithAppID(req.AppID, true),
		subscriber1.WithEmailAddress(req.EmailAddress, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSubscriber",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSubscriberResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateSubscriber(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSubscriber",
			"In", in,
			"Error", err,
		)
		return &npool.CreateSubscriberResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateSubscriberResponse{
		Info: info,
	}, nil
}
