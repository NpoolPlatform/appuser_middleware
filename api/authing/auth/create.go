package auth

import (
	"context"

	auth1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/auth"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAuth(ctx context.Context, in *npool.CreateAuthRequest) (*npool.CreateAuthResponse, error) {
	req := in.GetInfo()
	handler, err := auth1.NewHandler(
		ctx,
		auth1.WithID(req.ID),
		auth1.WithAppID(req.GetAppID()),
		auth1.WithRoleID(req.RoleID),
		auth1.WithUserID(req.UserID),
		auth1.WithResource(req.GetResource()),
		auth1.WithMethod(req.GetMethod()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAuth",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAuthResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateAuth(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAuth",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAuthResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateAuthResponse{
		Info: info,
	}, nil
}
