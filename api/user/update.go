package user

import (
	"context"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
)

func (s *Server) UpdateUser(ctx context.Context, in *npool.UpdateUserRequest) (*npool.UpdateUserResponse, error) {
	return &npool.UpdateUserResponse{}, nil
}
