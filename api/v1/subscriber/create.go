//nolint:dupl
package subscriber

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"

	mgrapi "github.com/NpoolPlatform/appuser-manager/api/v2/subscriber"

	subscriber1 "github.com/NpoolPlatform/appuser-middleware/pkg/subscriber"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateSubscriber(ctx context.Context, in *npool.CreateSubscriberRequest) (*npool.CreateSubscriberResponse, error) {
	if err := mgrapi.Validate(in.GetInfo()); err != nil {
		return &npool.CreateSubscriberResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := subscriber1.CreateSubscriber(ctx, in.GetInfo())
	if err != nil {
		return &npool.CreateSubscriberResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSubscriberResponse{
		Info: info,
	}, nil
}
