//nolint:dupl
package history

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/history"

	authing1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/history"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetHistories(ctx context.Context, in *npool.GetHistoriesRequest) (*npool.GetHistoriesResponse, error) {
	handler, err := authing1.NewHandler(
		ctx,
		authing1.WithAppID(in.GetAppID()),
		authing1.WithOffset(in.GetOffset()),
		authing1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAuths",
			"In", in,
			"Error", err,
		)
		return &npool.GetHistoriesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetHistories(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAuths",
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
