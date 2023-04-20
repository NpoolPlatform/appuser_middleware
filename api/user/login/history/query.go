package history

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	history1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user/login/history"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/login/history"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetHistory(ctx context.Context, in *npool.GetHistoryRequest) (*npool.GetHistoryResponse, error) {
	handler, err := history1.NewHandler(
		ctx,
		history1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetHistory",
			"In", in,
			"error", err,
		)
		return &npool.GetHistoryResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.GetHistory(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetHistory",
			"In", in,
			"error", err,
		)
		return &npool.GetHistoryResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetHistoryResponse{
		Info: info,
	}, nil
}

func (s *Server) GetHistories(ctx context.Context, in *npool.GetHistoriesRequest) (*npool.GetHistoriesResponse, error) {
	handler, err := history1.NewHandler(
		ctx,
		history1.WithConds(in.GetConds()),
		history1.WithOffset(in.GetOffset()),
		history1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetHistories",
			"In", in,
			"error", err,
		)
		return &npool.GetHistoriesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, total, err := handler.GetHistories(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetHistories",
			"In", in,
			"error", err,
		)
		return &npool.GetHistoriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetHistoriesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
