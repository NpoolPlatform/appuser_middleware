//nolint:dupl
package auth

import (
	"context"

	auth1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/auth"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistAuth(ctx context.Context, in *npool.ExistAuthRequest) (*npool.ExistAuthResponse, error) {
	handler, err := auth1.NewHandler(
		ctx,
		auth1.WithAppID(in.GetAppID()),
		auth1.WithUserID(in.UserID),
		auth1.WithResource(in.GetResource()),
		auth1.WithMethod(in.GetMethod()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAuth",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAuthResponse{}, status.Error(codes.Aborted, err.Error())
	}
	exist, err := handler.ExistAuth(ctx)
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
