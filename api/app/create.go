package app

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	recaptcha "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/recaptcha"
	signmethod "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"

	appmgrapi "github.com/NpoolPlatform/appuser-manager/api/v2/app"
	appmgrcli "github.com/NpoolPlatform/appuser-manager/pkg/client/app"
)

func (s *Server) CreateApp(ctx context.Context, in *npool.CreateAppRequest) (*npool.CreateAppResponse, error) {
	if err := appmgrapi.Validate(in.GetInfo()); err != nil {
		logger.Sugar().Errorw("CreateApp", "error", err)
		return &npool.CreateAppResponse{}, err
	}

	info, err := appmgrcli.CreateApp(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateApp", "error", err)
		return &npool.CreateAppResponse{}, err
	}

	return &npool.CreateAppResponse{
		Info: &npool.App{
			ID:        info.ID,
			CreatedBy: info.CreatedBy,
			Name:      info.Name,
			Logo:      info.Logo,

			RecaptchaMethod: recaptcha.RecaptchaType_GoogleRecaptchaV3.String(),
			SignupMethods: []string{
				signmethod.SignMethodType_Mobile.String(),
				signmethod.SignMethodType_Email.String(),
			},
			ExtSigninMethods: []string{
				signmethod.SignMethodType_Github.String(),
				signmethod.SignMethodType_Google.String(),
			},

			KycEnable:          true,
			SigninVerifyEnable: false,
			InvitationCodeMust: false,
			CreatedAt:          info.CreatedAt,
		},
	}, nil
}
