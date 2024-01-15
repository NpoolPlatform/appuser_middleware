package recoverycode

import (
	"context"

	recoverycode1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user/recoverycode"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/recoverycode"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateRecoveryCode(ctx context.Context, in *npool.UpdateRecoveryCodeRequest) (*npool.UpdateRecoveryCodeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateRecoveryCode",
			"In", in,
		)
		return &npool.UpdateRecoveryCodeResponse{}, status.Error(codes.Aborted, "Info is empty")
	}
	handler, err := recoverycode1.NewHandler(
		ctx,
		recoverycode1.WithID(req.ID, true),
		recoverycode1.WithEntID(req.EntID, false),
		recoverycode1.WithUsed(req.Used, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRecoveryCode",
			"In", in,
			"error", err,
		)
		return &npool.UpdateRecoveryCodeResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateRecoveryCode(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRecoveryCode",
			"In", in,
			"error", err,
		)
		return &npool.UpdateRecoveryCodeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateRecoveryCodeResponse{
		Info: info,
	}, nil
}
