package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	user1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteUser(ctx context.Context, in *npool.DeleteUserRequest) (*npool.DeleteUserResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteUser",
			"Req", req,
		)
		return &npool.DeleteUserResponse{}, status.Error(codes.InvalidArgument, "invalid userinfo")
	}

	handler, err := user1.NewHandler(
		ctx,
		user1.WithAppID(req.AppID, true),
		user1.WithID(req.ID, false),
		user1.WithEntID(req.EntID, false),
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

func (s *Server) DeleteThirdUser(ctx context.Context, in *npool.DeleteThirdUserRequest) (*npool.DeleteThirdUserResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteThirdUser",
			"Req", req,
		)
		return &npool.DeleteThirdUserResponse{}, status.Error(codes.InvalidArgument, "invalid userinfo")
	}

	handler, err := user1.NewHandler(
		ctx,
		user1.WithAppID(req.AppID, true),
		user1.WithID(req.ID, true),
		user1.WithThirdPartyID(req.ThirdPartyID, true),
		user1.WithThirdPartyUserID(req.ThirdPartyUserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteThirdUser",
			"Req", req,
			"error", err,
		)
		return &npool.DeleteThirdUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.DeleteThirdUser(ctx)
	if err != nil {
		return &npool.DeleteThirdUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteThirdUserResponse{
		Info: info,
	}, nil
}
