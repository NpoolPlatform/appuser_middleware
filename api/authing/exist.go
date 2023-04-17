//nolint:dupl
package authing

import (
	"context"

	authing1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistAuth(ctx context.Context, in *npool.ExistAuthRequest) (*npool.ExistAuthResponse, error) {
	handler, err := authing1.NewHandler(
		ctx,
		authing1.WithAppID(in.GetAppID()),
		authing1.WithUserID(in.UserID),
		authing1.WithResource(in.GetResource()),
		authing1.WithMethod(in.GetMethod()),
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
