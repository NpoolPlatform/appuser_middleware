package user

import (
	"context"

	user1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(ctx context.Context, in *npool.CreateUserRequest) (*npool.CreateUserResponse, error) {
	req := in.GetInfo()
	handler, err := user1.NewHandler(
		ctx,
		user1.WithID(req.ID),
		user1.WithAppID(req.GetAppID()),
		user1.WithPhoneNO(req.PhoneNO),
		user1.WithEmailAddress(req.EmailAddress),
		user1.WithImportedFromAppID(req.ImportedFromAppID),
		user1.WithPasswordHash(req.PasswordHash),
		user1.WithRoleIDs(req.GetRoleIDs()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateUser",
			"Req", req,
			"error", err,
		)
		return &npool.CreateUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.CreateUser(ctx)
	if err != nil {
		return &npool.CreateUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateUserResponse{
		Info: info,
	}, nil
}
