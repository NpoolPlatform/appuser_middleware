package app

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	app1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistApp(ctx context.Context, in *npool.ExistAppRequest) (*npool.ExistAppResponse, error) {
	handler, err := app1.NewHandler(
		ctx,
		app1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistApp",
			"In", in,
			"error", err,
		)
		return &npool.ExistAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.ExistApp(ctx)
	if err != nil {
		return &npool.ExistAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppResponse{
		Info: info,
	}, nil
}

func (s *Server) ExistAppConds(ctx context.Context, in *npool.ExistAppCondsRequest) (*npool.ExistAppCondsResponse, error) {
	handler, err := app1.NewHandler(
		ctx,
		app1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAppConds",
			"In", in,
			"error", err,
		)
		return &npool.ExistAppCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	exist, err := handler.ExistAppConds(ctx)
	if err != nil {
		return &npool.ExistAppCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAppCondsResponse{
		Info: exist,
	}, nil
}
