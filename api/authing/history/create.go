package history

import (
	"context"

	handler "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/handler"
	history1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/history"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/history"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateHistory(ctx context.Context, in *npool.CreateHistoryRequest) (*npool.CreateHistoryResponse, error) {
	req := in.GetInfo()
	h, err := handler.NewHandler(
		ctx,
		handler.WithID(req.ID),
		handler.WithAppID(req.GetAppID()),
		handler.WithUserID(req.UserID),
		handler.WithResource(req.GetResource()),
		handler.WithMethod(req.GetMethod()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateHistory",
			"In", in,
			"Error", err,
		)
		return &npool.CreateHistoryResponse{}, status.Error(codes.Aborted, err.Error())
	}
	_handler := (&history1.Handler{
		Handler: h,
	}).WithAllowed(req.GetAllowed())
	info, err := _handler.CreateHistory(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateHistory",
			"In", in,
			"Error", err,
		)
		return &npool.CreateHistoryResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateHistoryResponse{
		Info: info,
	}, nil
}
