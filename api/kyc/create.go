package kyc

import (
	"context"

	common "github.com/NpoolPlatform/appuser-middleware/api/common"
	kyc1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/kyc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateKyc(ctx context.Context, in *npool.CreateKycRequest) (*npool.CreateKycResponse, error) {
	req := in.GetInfo()
	handler, err := kyc1.NewHandler(
		ctx,
		kyc1.WithID(req.ID),
		kyc1.WithAppID(req.GetAppID()),
		kyc1.WithUserID(req.GetUserID()),
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
			"CreateKyc",
			"In", in,
			"Error", err,
		)
		return &npool.CreateKycResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if req.UserID != nil {
		if err := common.ValidateUser(ctx, req.GetAppID(), req.GetUserID()); err != nil {
			logger.Sugar().Errorw(
				"CreateAuth",
				"In", in,
				"Error", err,
			)
			return &npool.CreateKycResponse{}, status.Error(codes.Aborted, err.Error())
		}
	}

	info, err := handler.CreateKyc(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateKyc",
			"In", in,
			"Error", err,
		)
		return &npool.CreateKycResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateKycResponse{
		Info: info,
	}, nil
}
