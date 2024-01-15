package recoverycode

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	recoverycode1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user/recoverycode"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/recoverycode"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistRecoveryCodeConds(ctx context.Context, in *npool.ExistRecoveryCodeCondsRequest) (*npool.ExistRecoveryCodeCondsResponse, error) {
	handler, err := recoverycode1.NewHandler(
		ctx,
		recoverycode1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistRecoveryCodeConds",
			"In", in,
			"error", err,
		)
		return &npool.ExistRecoveryCodeCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	exist, err := handler.ExistRecoveryCodeConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistRecoveryCodeConds",
			"In", in,
			"error", err,
		)
		return &npool.ExistRecoveryCodeCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistRecoveryCodeCondsResponse{
		Info: exist,
	}, nil
}
