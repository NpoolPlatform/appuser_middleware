package user

import (
	"context"

	user1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/role/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteUser(ctx context.Context, in *npool.DeleteUserRequest) (*npool.DeleteUserResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteUser",
			"In", in,
		)
		return &npool.DeleteUserResponse{}, status.Error(codes.Aborted, "Info is empty")
	}
	handler, err := user1.NewHandler(
		ctx,
		user1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteUser",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteUserResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteUserResponse{
		Info: info,
	}, nil
}
