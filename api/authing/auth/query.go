//nolint:dupl
package auth

import (
	"context"

	auth1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/auth"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAuth(ctx context.Context, in *npool.GetAuthRequest) (*npool.GetAuthResponse, error) {
	handler, err := auth1.NewHandler(
		ctx,
		auth1.WithID(&in.ID),
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
	handler, err := auth1.NewHandler(
		ctx,
		auth1.WithAppID(in.GetAppID()),
		auth1.WithOffset(in.GetOffset()),
		auth1.WithLimit(in.GetLimit()),
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
