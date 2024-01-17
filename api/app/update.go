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
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateApp",
			"In", in,
		)
		return &npool.UpdateAppResponse{}, status.Error(codes.Aborted, "Info is empty")
	}
	handler, err := app1.NewHandler(
		ctx,
		app1.WithID(req.ID, true),
		app1.WithEntID(req.EntID, false),
		app1.WithName(req.Name, false),
		app1.WithLogo(req.Logo, false),
		app1.WithDescription(req.Description, false),
		app1.WithSignupMethods(req.GetSignupMethods(), false),
		app1.WithExtSigninMethods(req.GetExtSigninMethods(), false),
		app1.WithRecaptchaMethod(req.RecaptchaMethod, false),
		app1.WithKycEnable(req.KycEnable, false),
		app1.WithSigninVerifyEnable(req.SigninVerifyEnable, false),
		app1.WithInvitationCodeMust(req.InvitationCodeMust, false),
		app1.WithCreateInvitationCodeWhen(req.CreateInvitationCodeWhen, false),
		app1.WithMaxTypedCouponsPerOrder(req.MaxTypedCouponsPerOrder, false),
		app1.WithMaintaining(req.Maintaining, false),
		app1.WithCommitButtonTargets(req.GetCommitButtonTargets(), false),
		app1.WithBanned(req.Banned, false),
		app1.WithBanMessage(req.BanMessage, false),
		app1.WithCouponWithdrawEnable(req.CouponWithdrawEnable, false),
		app1.WithResetUserMethod(req.ResetUserMethod, false),
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
