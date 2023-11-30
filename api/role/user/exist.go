package user

import (
	"context"

	user1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/role/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistUserConds(ctx context.Context, in *npool.ExistUserCondsRequest) (*npool.ExistUserCondsResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistUserConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistUserCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	exist, err := handler.ExistUserConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistUserConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistUserCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistUserCondsResponse{
		Info: exist,
	}, nil
}
