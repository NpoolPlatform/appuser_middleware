package subscriber

import (
	"context"

	subscriber1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/subscriber"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteSubscriber(ctx context.Context, in *npool.DeleteSubscriberRequest) (*npool.DeleteSubscriberResponse, error) {
	req := in.GetInfo()
	handler, err := subscriber1.NewHandler(
		ctx,
		subscriber1.WithID(req.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteSubscriber",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteSubscriberResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteSubscriber(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteSubscriber",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteSubscriberResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteSubscriberResponse{
		Info: info,
	}, nil
}
