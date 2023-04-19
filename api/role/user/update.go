package user

import (
	"context"

	user1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/role/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateUser(ctx context.Context, in *npool.UpdateUserRequest) (*npool.UpdateUserResponse, error) {
	req := in.GetInfo()
	handler, err := user1.NewHandler(
		ctx,
		user1.WithID(req.ID),
		user1.WithRoleID(req.RoleID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateUser",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateUserResponse{
		Info: info,
	}, nil
}
