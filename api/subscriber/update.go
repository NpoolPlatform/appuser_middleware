package subscriber

import (
	"context"

	subscriber1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/subscriber"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateSubscriber(ctx context.Context, in *npool.UpdateSubscriberRequest) (*npool.UpdateSubscriberResponse, error) {
	req := in.GetInfo()
	handler, err := subscriber1.NewHandler(
		ctx,
		subscriber1.WithID(req.ID),
		subscriber1.WithRegistered(req.Registered),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSubscriber",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSubscriberResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateSubscriber(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSubscriber",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSubscriberResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateSubscriberResponse{
		Info: info,
	}, nil
}
