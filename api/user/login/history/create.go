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
	handler, err := history1.NewHandler(
		ctx,
		history1.WithID(req.ID),
		history1.WithAppID(req.GetAppID()),
		history1.WithUserID(req.GetUserID()),
		history1.WithClientIP(req.ClientIP),
		history1.WithUserAgent(req.UserAgent),
		history1.WithLocation(req.Location),
		history1.WithLoginType(req.LoginType),
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
