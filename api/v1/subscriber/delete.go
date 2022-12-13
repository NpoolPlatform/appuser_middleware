//nolint:dupl
package subscriber

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"

	subscriber1 "github.com/NpoolPlatform/appuser-middleware/pkg/subscriber"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteSubscriber(ctx context.Context, in *npool.DeleteSubscriberRequest) (*npool.DeleteSubscriberResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		return &npool.DeleteSubscriberResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := subscriber1.DeleteSubscriber(ctx, in.GetID())
	if err != nil {
		return &npool.DeleteSubscriberResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteSubscriberResponse{
		Info: info,
	}, nil
}
