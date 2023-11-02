package kyc

import (
	"context"

	kyc1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/kyc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateKyc(ctx context.Context, in *npool.UpdateKycRequest) (*npool.UpdateKycResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateKyc",
			"In", in,
		)
		return &npool.UpdateKycResponse{}, status.Error(codes.Aborted, "Info is empty")
	}
	handler, err := kyc1.NewHandler(
		ctx,
		kyc1.WithID(req.ID, true),
		kyc1.WithDocumentType(req.DocumentType, false),
		kyc1.WithIDNumber(req.IDNumber, false),
		kyc1.WithFrontImg(req.FrontImg, false),
		kyc1.WithBackImg(req.BackImg, false),
		kyc1.WithSelfieImg(req.SelfieImg, false),
		kyc1.WithEntityType(req.EntityType, false),
		kyc1.WithReviewID(req.ReviewID, false),
		kyc1.WithState(req.State, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateKyc",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateKycResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateKyc(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateKyc",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateKycResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateKycResponse{
		Info: info,
	}, nil
}
