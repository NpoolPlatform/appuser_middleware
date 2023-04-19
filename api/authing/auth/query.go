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

func (s *Server) GetAuth(ctx context.Context, in *npool.GetAuthRequest) (*npool.GetAuthResponse, error) {
	_handler, err := auth1.NewHandler(
		ctx,
		handler.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAuth",
			"In", in,
			"Error", err,
		)
		return &npool.GetAuthResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := _handler.GetAuth(ctx)
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
	_handler, err := auth1.NewHandler(
		ctx,
		auth1.WithConds(in.GetConds()),
		handler.WithOffset(in.GetOffset()),
		handler.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAuths",
			"In", in,
			"Error", err,
		)
		return &npool.GetAuthsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := _handler.GetAuths(ctx)
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
