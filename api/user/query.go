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

func (s *Server) GetUser(ctx context.Context, in *npool.GetUserRequest) (*npool.GetUserResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithAppID(in.AppID),
		user1.WithID(&in.UserID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUser",
			"In", in,
			"error", err,
		)
		return &npool.GetUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.GetUser(ctx)
	if err != nil {
		return &npool.GetUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetUserResponse{
		Info: info,
	}, nil
}

func (s *Server) GetUsers(ctx context.Context, in *npool.GetUsersRequest) (*npool.GetUsersResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithConds(in.GetConds()),
		user1.WithOffset(in.GetOffset()),
		user1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUsers",
			"In", in,
			"error", err,
		)
		return &npool.GetUsersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, total, err := handler.GetUsers(ctx)
	if err != nil {
		return &npool.GetUsersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetUsersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
