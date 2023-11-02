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
	if req == nil {
		logger.Sugar().Errorw(
			"CreateUser",
			"Req", req,
		)
		return &npool.CreateUserResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := user1.NewHandler(
		ctx,
		user1.WithEntID(req.EntID, false),
		user1.WithAppID(req.AppID, true),
		user1.WithPhoneNO(req.PhoneNO, false),
		user1.WithEmailAddress(req.EmailAddress, false),
		user1.WithImportFromAppID(req.ImportedFromAppID, false),
		user1.WithPasswordHash(req.PasswordHash, true),
		user1.WithRoleIDs(req.GetRoleIDs(), true),
		user1.WithThirdPartyID(req.ThirdPartyID, false),
		user1.WithThirdPartyUserID(req.ThirdPartyUserID, false),
		user1.WithThirdPartyUsername(req.ThirdPartyUsername, false),
		user1.WithThirdPartyAvatar(req.ThirdPartyAvatar, false),
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
		logger.Sugar().Errorw(
			"CreateUser",
			"Req", req,
			"error", err,
		)
		return &npool.CreateUserResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateUserResponse{
		Info: info,
	}, nil
}
