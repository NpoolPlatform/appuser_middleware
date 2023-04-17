//nolint:nolintlint,dupl
package app

import (
	"context"

	app1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetApp(ctx context.Context, in *npool.GetAppRequest) (*npool.GetAppResponse, error) {
	handler, err := app1.NewHandler(
		ctx,
		app1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetApp",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetApp(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetApp",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAppResponse{
		Info: info,
	}, nil
}

func (s *Server) GetApps(ctx context.Context, in *npool.GetAppsRequest) (*npool.GetAppsResponse, error) {
	handler, err := app1.NewHandler(
		ctx,
		app1.WithConds(in.Conds),
		app1.WithOffset(in.GetOffset()),
		app1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetApps",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetApps(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetApps",
			"In", in,
			"Error", err,
		)
		return &npool.GetAppsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetAppsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
