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
	handler, err := app1.NewHandler(
		ctx,
		app1.WithID(req.ID),
		app1.WithCreatedBy(req.GetCreatedBy()),
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
		app1.WithReqs(in.GetInfos()),
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
