//nolint:dupl
package authing

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing"

	authing1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAuth(ctx context.Context, in *npool.GetAuthRequest) (*npool.GetAuthResponse, error) {
	handler, err := authing1.NewHandler(
		ctx,
		authing1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAuth",
			"In", in,
			"Error", err,
		)
		return &npool.GetAuthResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetAuth(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAuth",
			"In", in,
			"Error", err,
		)
		return &npool.GetAuthResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAuthResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAuths(ctx context.Context, in *npool.GetAuthsRequest) (*npool.GetAuthsResponse, error) {
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
		return &npool.GetAuthsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetAuths(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAuths",
			"In", in,
			"Error", err,
		)
		return &npool.GetAuthsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAuthsResponse{
		Infos: infos,
		Total: total,
	}, nil
}

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
