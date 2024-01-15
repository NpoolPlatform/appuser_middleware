package recoverycode

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	recoverycode1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user/recoverycode"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/recoverycode"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetRecoveryCodess(ctx context.Context, in *npool.GetRecoveryCodesRequest) (*npool.GetRecoveryCodesResponse, error) {
	handler, err := recoverycode1.NewHandler(
		ctx,
		recoverycode1.WithConds(in.GetConds()),
		recoverycode1.WithOffset(in.GetOffset()),
		recoverycode1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRecoveryCodes",
			"In", in,
			"error", err,
		)
		return &npool.GetRecoveryCodesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	infos, total, err := handler.GetRecoveryCodes(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRecoveryCodes",
			"In", in,
			"error", err,
		)
		return &npool.GetRecoveryCodesResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetRecoveryCodesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
