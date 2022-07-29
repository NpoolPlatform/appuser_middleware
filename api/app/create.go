package app

import (
	"context"
	capp "github.com/NpoolPlatform/appuser-middleware/pkg/converter/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	mw "github.com/NpoolPlatform/appuser-middleware/pkg/app"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
)

func (s *Server) CreateApp(ctx context.Context, in *npool.CreateAppRequest) (*npool.CreateAppResponse, error) {
	if err := validate(in.GetInfo()); err != nil {
		logger.Sugar().Errorw("CreateApp", "error", err)
		return &npool.CreateAppResponse{}, err
	}

	info, err := mw.CreateApp(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateApp", "error", err)
		return &npool.CreateAppResponse{}, err
	}

	ginfo, err := capp.CreateEnt2Grpc(info)
	if err != nil {
		logger.Sugar().Errorw("CreateApp", "error", err)
		return &npool.CreateAppResponse{}, status.Error(codes.Internal, "invalid value")
	}

	return &npool.CreateAppResponse{
		Info: ginfo,
	}, nil

}
