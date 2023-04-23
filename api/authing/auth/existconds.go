package auth

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	auth1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/auth"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistAuthConds(ctx context.Context, in *npool.ExistAuthCondsRequest) (*npool.ExistAuthCondsResponse, error) {
	handler, err := auth1.NewHandler(
		ctx,
		auth1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAuthConds",
			"In", in,
			"error", err,
		)
		return &npool.ExistAuthCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	exist, err := handler.ExistAuthConds(ctx)
	if err != nil {
		return &npool.ExistAuthCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAuthCondsResponse{
		Info: exist,
	}, nil
}
