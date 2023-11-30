package history

import (
	"context"

	history1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user/login/history"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/login/history"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateHistory(ctx context.Context, in *npool.CreateHistoryRequest) (*npool.CreateHistoryResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateHistory",
			"Req", req,
		)
		return &npool.CreateHistoryResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := history1.NewHandler(
		ctx,
		history1.WithEntID(req.EntID, false),
		history1.WithAppID(req.AppID, true),
		history1.WithUserID(req.UserID, true),
		history1.WithClientIP(req.ClientIP, true),
		history1.WithUserAgent(req.UserAgent, true),
		history1.WithLocation(req.Location, true),
		history1.WithLoginType(req.LoginType, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateHistory",
			"Req", req,
			"error", err,
		)
		return &npool.CreateHistoryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.CreateHistory(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateHistory",
			"Req", req,
			"error", err,
		)
		return &npool.CreateHistoryResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateHistoryResponse{
		Info: info,
	}, nil
}
