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

func (s *Server) ExistAuth(ctx context.Context, in *npool.ExistAuthRequest) (*npool.ExistAuthResponse, error) {
	h, err := handler.NewHandler(
		ctx,
		handler.WithAppID(in.GetAppID()),
		handler.WithUserID(in.UserID),
		handler.WithResource(in.GetResource()),
		handler.WithMethod(in.GetMethod()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAuth",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAuthResponse{}, status.Error(codes.Aborted, err.Error())
	}
	_handler := &auth1.Handler{
		Handler: h,
	}
	exist, err := _handler.ExistAuth(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAuth",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAuthResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistAuthResponse{
		Info: exist,
	}, nil
}
