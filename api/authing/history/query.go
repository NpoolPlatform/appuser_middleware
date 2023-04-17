//nolint:dupl
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

func (s *Server) GetHistories(ctx context.Context, in *npool.GetHistoriesRequest) (*npool.GetHistoriesResponse, error) {
	handler, err := handler.NewHandler(
		ctx,
		handler.WithConds(in.GetConds()),
		handler.WithOffset(in.GetOffset()),
		handler.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetHistories",
			"In", in,
			"Error", err,
		)
		return &npool.GetHistoriesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	_handler := &history1.Handler{
		Handler: handler,
	}
	infos, total, err := _handler.GetHistories(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetHistories",
			"In", in,
			"Error", err,
		)
		return &npool.GetHistoriesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.GetHistoriesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
