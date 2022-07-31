package user

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
)

func (s *Server) UpdateUser(ctx context.Context, in *npool.UpdateUserRequest) (*npool.UpdateUserResponse, error) {
	return &npool.UpdateUserResponse{}, status.Error(codes.Internal, err.Error())
}
