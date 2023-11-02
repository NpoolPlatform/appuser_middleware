package user

import (
	"context"

	user1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/role/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetUser(ctx context.Context, in *npool.GetUserRequest) (*npool.GetUserResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUser",
			"In", in,
			"Error", err,
		)
		return &npool.GetUserResponse{}, status.Error(codes.Aborted, err.Error())
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
			"Error", err,
		)
		return &npool.GetUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetUsers(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetUsers",
			"In", in,
			"Error", err,
		)
		return &npool.GetUsersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetUsersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
