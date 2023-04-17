package subscriber

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"

	mgrapi "github.com/NpoolPlatform/appuser-manager/api/subscriber"

	subscriber1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/subscriber"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateSubscriber(ctx context.Context, in *npool.UpdateSubscriberRequest) (*npool.UpdateSubscriberResponse, error) {
	if err := mgrapi.Validate(in.GetInfo()); err != nil {
		return &npool.UpdateSubscriberResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := subscriber1.UpdateSubscriber(ctx, in.GetInfo())
	if err != nil {
		return &npool.UpdateSubscriberResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateSubscriberResponse{
		Info: info,
	}, nil
}
