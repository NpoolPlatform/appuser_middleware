package auth

import (
	"context"

	auth1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/auth"
	handler "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/handler"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateAuth(ctx context.Context, in *npool.CreateAuthRequest) (*npool.CreateAuthResponse, error) {
	req := in.GetInfo()
	_handler, err := auth1.NewHandler(
		ctx,
		handler.WithID(req.ID),
		handler.WithAppID(req.GetAppID()),
		handler.WithRoleID(req.RoleID),
		handler.WithUserID(req.UserID),
		handler.WithResource(req.Resource),
		handler.WithMethod(req.Method),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAuth",
			"In", in,
			"Error", err,
		)
		return &npool.CreateAuthResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := _handler.CreateAuth(ctx)
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
