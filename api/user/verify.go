package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	user1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) VerifyAccount(ctx context.Context, in *npool.VerifyAccountRequest) (*npool.VerifyAccountResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithAppID(in.GetAppID()),
		user1.WithAccount(in.GetAccount(), in.GetAccountType()),
		user1.WithPasswordHash(&in.PasswordHash),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"VerifyAccount",
			"In", in,
			"error", err,
		)
		return &npool.VerifyAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.VerifyAccount(ctx)
	if err != nil {
		return &npool.VerifyAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.VerifyAccountResponse{
		Info: info,
	}, nil
}

func (s *Server) VerifyUser(ctx context.Context, in *npool.VerifyUserRequest) (*npool.VerifyUserResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithAppID(in.GetAppID()),
		user1.WithID(&in.UserID),
		user1.WithPasswordHash(&in.PasswordHash),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"VerifyUser",
			"In", in,
			"error", err,
		)
		return &npool.VerifyUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.VerifyUser(ctx)
	if err != nil {
		return &npool.VerifyUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.VerifyUserResponse{
		Info: info,
	}, nil
}
