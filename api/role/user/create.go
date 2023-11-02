package user

import (
	"context"

	user1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/role/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(ctx context.Context, in *npool.CreateUserRequest) (*npool.CreateUserResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateUser",
			"In", in,
		)
		return &npool.CreateUserResponse{}, status.Error(codes.Aborted, "Info is empty")
	}
	handler, err := user1.NewHandler(
		ctx,
		user1.WithEntID(req.EntID, false),
		user1.WithAppID(req.AppID, true),
		user1.WithRoleID(req.RoleID, true),
		user1.WithUserID(req.UserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUser",
			"In", in,
			"Error", err,
		)
		return &npool.CreateUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateUserResponse{
		Info: info,
	}, nil
}
