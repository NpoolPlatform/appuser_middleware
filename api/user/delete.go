package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	user1 "github.com/NpoolPlatform/appuser-middleware/pkg/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteUser(ctx context.Context, in *npool.DeleteUserRequest) (*npool.DeleteUserResponse, error) {
	req := in.GetInfo()
	if req == nil {
		return &npool.DeleteUserResponse{}, status.Error(codes.InvalidArgument, "invalid userinfo")
	}

	handler, err := user1.NewHandler(
		ctx,
		user1.WithAppID(req.GetAppID()),
		user1.WithID(req.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteUser",
			"Req", req,
			"error", err,
		)
		return &npool.DeleteUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.DeleteUser(ctx)
	if err != nil {
		return &npool.DeleteUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteUserResponse{
		Info: info,
	}, nil
}
