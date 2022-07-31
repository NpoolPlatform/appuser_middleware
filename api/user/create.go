package user

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
)

func (s *Server) CreateUser(ctx context.Context, in *npool.CreateUserRequest) (*npool.CreateUserResponse, error) {
	return &npool.CreateUserResponse{}, status.Error(codes.Internal, err.Error())
}
