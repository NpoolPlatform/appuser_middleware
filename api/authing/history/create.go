package history

/*
import (
	"context"

	history1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/history"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/history"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateHistory(ctx context.Context, in *npool.CreateHistoryRequest) (*npool.CreateHistoryResponse, error) {
	handler, err := history1.NewHandler(
		ctx,
		history1.WithConds(in.GetConds()),
		history1.WithOffset(in.GetOffset()),
		history1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAuths",
			"In", in,
			"Error", err,
		)
		return &npool.CreateHistoryResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateHistory(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAuths",
			"In", in,
			"Error", err,
		)
		return &npool.CreateHistoryResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateHistoryResponse{
		Info: info,
	}, nil
}
*/
