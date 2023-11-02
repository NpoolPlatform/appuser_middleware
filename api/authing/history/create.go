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
	if req == nil {
		logger.Sugar().Errorw(
			"CreateHistory",
			"In", in,
		)
		return &npool.CreateHistoryResponse{}, status.Error(codes.Aborted, "Info is empty")
	}
	_handler, err := history1.NewHandler(
		ctx,
		handler.WithEntID(req.EntID, false),
		handler.WithAppID(req.AppID, true),
		handler.WithUserID(req.UserID, false),
		handler.WithResource(req.Resource, true),
		handler.WithMethod(req.Method, true),
		history1.WithAllowed(req.Allowed, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateHistory",
			"In", in,
			"Error", err,
		)
		return &npool.CreateHistoryResponse{}, status.Error(codes.Aborted, err.Error())
	}
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
