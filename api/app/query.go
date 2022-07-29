package app

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
)

func (s *Server) GetApp(ctx context.Context, in *npool.GetAppRequest) (*npool.GetAppResponse, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
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
