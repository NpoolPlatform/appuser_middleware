package app

import (
	"context"

	app1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
)

func (s *Server) DeleteApp(ctx context.Context, in *npool.DeleteAppRequest) (*npool.DeleteAppResponse, error) {
	id := in.GetInfo().GetID()
	handler, err := app1.NewHandler(
		ctx,
		app1.WithID(&id),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteApp",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAppResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteApp(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteApp",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAppResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteAppResponse{
		Info: info,
	}, nil
}
