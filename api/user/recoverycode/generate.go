package recoverycode

import (
	"context"

	recoverycode1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user/recoverycode"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/recoverycode"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GenerateRecoveryCodes(ctx context.Context, in *npool.GenerateRecoveryCodesRequest) (*npool.GenerateRecoveryCodesResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"GenerateRecoveryCode",
			"Req", req,
		)
		return &npool.GenerateRecoveryCodesResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := recoverycode1.NewHandler(
		ctx,
		recoverycode1.WithAppID(req.AppID, true),
		recoverycode1.WithUserID(req.UserID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GenerateRecoveryCode",
			"Req", req,
			"error", err,
		)
		return &npool.GenerateRecoveryCodesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, err := handler.GenerateRecoveryCodes(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GenerateRecoveryCode",
			"Req", req,
			"error", err,
		)
		return &npool.GenerateRecoveryCodesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GenerateRecoveryCodesResponse{
		Infos: infos,
	}, nil
}
