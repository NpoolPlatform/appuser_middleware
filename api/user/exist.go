//nolint:dupl
package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	user1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistUser(ctx context.Context, in *npool.ExistUserRequest) (*npool.ExistUserResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithAppID(&in.AppID, true),
		user1.WithEntID(&in.UserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistUser",
			"In", in,
			"error", err,
		)
		return &npool.ExistUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.ExistUser(ctx)
	if err != nil {
		return &npool.ExistUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistUserResponse{
		Info: info,
	}, nil
}

func (s *Server) ExistUserConds(ctx context.Context, in *npool.ExistUserCondsRequest) (*npool.ExistUserCondsResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistUserConds",
			"In", in,
			"error", err,
		)
		return &npool.ExistUserCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	exist, err := handler.ExistUserConds(ctx)
	if err != nil {
		return &npool.ExistUserCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistUserCondsResponse{
		Info: exist,
	}, nil
}
