//nolint:dupl
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

func (s *Server) DeleteAuth(ctx context.Context, in *npool.DeleteAuthRequest) (*npool.DeleteAuthResponse, error) {
	req := in.GetInfo()
	handler, err := handler.NewHandler(
		ctx,
		handler.WithID(req.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAuth",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAuthResponse{}, status.Error(codes.Aborted, err.Error())
	}
	_handler := &auth1.Handler{
		Handler: handler,
	}
	info, err := _handler.DeleteAuth(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAuth",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAuthResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteAuthResponse{
		Info: info,
	}, nil
}
