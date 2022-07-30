package app

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	mapp "github.com/NpoolPlatform/appuser-middleware/pkg/app"
	capp "github.com/NpoolPlatform/appuser-middleware/pkg/converter/app"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) GetApp(ctx context.Context, in *npool.GetAppRequest) (*npool.GetAppResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetApp", "error", err)
		return &npool.GetAppResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	info, err := mapp.GetApp(ctx, in.GetAppID())
	if err != nil {
		logger.Sugar().Errorw("GetApp", "error", err)
		return &npool.GetAppResponse{}, status.Error(codes.Internal, "fail create app")
	}

	ginfo, err := capp.Ent2Grpc(info)
	if err != nil {
		logger.Sugar().Errorw("GetApp", "error", err)
		return &npool.GetAppResponse{}, status.Error(codes.Internal, "invalid value")
	}

	return &npool.GetAppResponse{
		Info: ginfo,
	}, nil
}

func (s *Server) GetApps(ctx context.Context, in *npool.GetAppsRequest) (*npool.GetAppsResponse, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (s *Server) GetUserApps(ctx context.Context, in *npool.GetUserAppsRequest) (*npool.GetUserAppsResponse, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (s *Server) GetSignMethods(ctx context.Context, in *npool.GetSignMethodsRequest) (*npool.GetSignMethodsResponse, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (s *Server) GetRecaptchas(ctx context.Context, in *npool.GetRecaptchasRequest) (*npool.GetRecaptchasResponse, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}
