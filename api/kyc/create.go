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
	if req == nil {
		logger.Sugar().Errorw(
			"CreateKyc",
			"In", in,
		)
		return &npool.CreateKycResponse{}, status.Error(codes.Aborted, "Info is empty")
	}
	handler, err := kyc1.NewHandler(
		ctx,
		kyc1.WithEntID(req.EntID, false),
		kyc1.WithAppID(req.AppID, true),
		kyc1.WithUserID(req.UserID, true),
		kyc1.WithDocumentType(req.DocumentType, true),
		kyc1.WithIDNumber(req.IDNumber, false),
		kyc1.WithFrontImg(req.FrontImg, true),
		kyc1.WithBackImg(req.BackImg, false),
		kyc1.WithSelfieImg(req.SelfieImg, true),
		kyc1.WithEntityType(req.EntityType, true),
		kyc1.WithReviewID(req.ReviewID, true),
		kyc1.WithState(req.State, false),
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
