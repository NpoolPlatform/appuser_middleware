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
	handler, err := kyc1.NewHandler(
		ctx,
		kyc1.WithID(req.ID),
		kyc1.WithDocumentType(req.DocumentType),
		kyc1.WithIDNumber(req.IDNumber),
		kyc1.WithFrontImg(req.FrontImg),
		kyc1.WithBackImg(req.BackImg),
		kyc1.WithSelfieImg(req.SelfieImg),
		kyc1.WithEntityType(req.EntityType),
		kyc1.WithReviewID(req.ReviewID),
		kyc1.WithState(req.State),
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
