package authing

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing"
)

func (s *Server) GetAuthOnly(ctx context.Context, in *npool.GetAuthOnlyRequest) (*npool.GetAuthOnlyResponse, error) {
	return nil, nil
}

func (s *Server) GetAuths(ctx context.Context, in *npool.GetAuthsRequest) (*npool.GetAuthsResponse, error) {
	return nil, nil
}

func (s *Server) GetHistories(ctx context.Context, in *npool.GetHistoriesRequest) (*npool.GetHistoriesResponse, error) {
	return nil, nil
}
