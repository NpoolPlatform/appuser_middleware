package kyc

import (
	"context"

	kyc1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/kyc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteKyc(ctx context.Context, in *npool.DeleteKycRequest) (*npool.DeleteKycResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteKyc",
			"In", in,
		)
		return &npool.DeleteKycResponse{}, status.Error(codes.Aborted, "Info is empty")
	}
	handler, err := kyc1.NewHandler(
		ctx,
		kyc1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteKyc",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteKycResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteKyc(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteKyc",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteKycResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteKycResponse{
		Info: info,
	}, nil
}
