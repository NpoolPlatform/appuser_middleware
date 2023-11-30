package kyc

import (
	"context"

	kyc1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/kyc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetKyc(ctx context.Context, in *npool.GetKycRequest) (*npool.GetKycResponse, error) {
	handler, err := kyc1.NewHandler(
		ctx,
		kyc1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetKyc",
			"In", in,
			"Error", err,
		)
		return &npool.GetKycResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetKyc(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetKyc",
			"In", in,
			"Error", err,
		)
		return &npool.GetKycResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetKycResponse{
		Info: info,
	}, nil
}

func (s *Server) GetKycs(ctx context.Context, in *npool.GetKycsRequest) (*npool.GetKycsResponse, error) {
	handler, err := kyc1.NewHandler(
		ctx,
		kyc1.WithConds(in.GetConds()),
		kyc1.WithOffset(in.GetOffset()),
		kyc1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetKycs",
			"In", in,
			"Error", err,
		)
		return &npool.GetKycsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetKycs(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetKycs",
			"In", in,
			"Error", err,
		)
		return &npool.GetKycsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetKycsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
