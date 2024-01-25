package app

import (
	"context"

	app1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateApp(ctx context.Context, in *npool.CreateAppRequest) (*npool.CreateAppResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateApp",
			"In", in,
		)
		return &npool.CreateAppResponse{}, status.Error(codes.Aborted, "Info is empty")
	}
	handler, err := app1.NewHandler(
		ctx,
		app1.WithEntID(req.EntID, false),
		app1.WithCreatedBy(req.CreatedBy, true),
		app1.WithName(req.Name, true),
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
		app1.WithCouponWithdrawEnable(req.CouponWithdrawEnable, false),
		app1.WithResetUserMethod(req.ResetUserMethod, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateApp",
			"In", in,
			"error", err,
		)
		return &npool.CreateAppResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateApp(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateApp",
			"In", in,
			"error", err,
		)
		return &npool.CreateAppResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateAppResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateApps(ctx context.Context, in *npool.CreateAppsRequest) (*npool.CreateAppsResponse, error) {
	handler, err := app1.NewHandler(
		ctx,
		app1.WithReqs(in.GetInfos(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateApps",
			"In", in,
			"error", err,
		)
		return &npool.CreateAppsResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.CreateApps(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateApps",
			"In", in,
			"error", err,
		)
		return &npool.CreateAppsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateAppsResponse{
		Infos: infos,
	}, nil
}
