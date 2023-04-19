package app

import (
	"context"

	app1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateApp(ctx context.Context, in *npool.UpdateAppRequest) (*npool.UpdateAppResponse, error) {
	req := in.GetInfo()
	handler, err := app1.NewHandler(
		ctx,
		app1.WithID(req.ID),
		app1.WithName(req.Name),
		app1.WithLogo(req.Logo),
		app1.WithDescription(req.Description),
		app1.WithSignupMethods(req.GetSignupMethods()),
		app1.WithExtSigninMethods(req.GetExtSigninMethods()),
		app1.WithRecaptchaMethod(req.RecaptchaMethod),
		app1.WithKycEnable(req.KycEnable),
		app1.WithSigninVerifyEnable(req.SigninVerifyEnable),
		app1.WithInvitationCodeMust(req.InvitationCodeMust),
		app1.WithCreateInvitationCodeWhen(req.CreateInvitationCodeWhen),
		app1.WithMaxTypedCouponsPerOrder(req.MaxTypedCouponsPerOrder),
		app1.WithMaintaining(req.Maintaining),
		app1.WithCommitButtonTargets(req.GetCommitButtonTargets()),
		app1.WithBanned(req.Banned),
		app1.WithBanMessage(req.BanMessage),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateApp",
			"In", in,
			"error", err,
		)
		return &npool.UpdateAppResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateApp(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateApp",
			"In", in,
			"error", err,
		)
		return &npool.UpdateAppResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateAppResponse{
		Info: info,
	}, nil
}
