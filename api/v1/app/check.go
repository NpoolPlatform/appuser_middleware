package app

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	mgrapp "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"
	mgrappcontrol "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
	mgrbanapp "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banapp"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/appuser-manager/api/v2/app"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appcontrol"
	"github.com/NpoolPlatform/appuser-manager/api/v2/banapp"
)

func validate(info *npool.AppReq) error {
	err := app.Validate(&mgrapp.AppReq{
		ID:          info.ID,
		Description: info.Description,
		CreatedBy:   info.CreatedBy,
		Name:        info.Name,
		Logo:        info.Logo,
	})
	if err != nil {
		logger.Sugar().Errorw("validate", err.Error())
		return status.Error(codes.InvalidArgument, err.Error())
	}

	err = appcontrol.Validate(&mgrappcontrol.AppControlReq{
		AppID:               info.ID,
		SignupMethods:       info.SignupMethods,
		ExternSigninMethods: info.ExtSigninMethods,
		RecaptchaMethod:     info.RecaptchaMethod,
		KycEnable:           info.KycEnable,
		SigninVerifyEnable:  info.SigninVerifyEnable,
		InvitationCodeMust:  info.InvitationCodeMust,
	})
	if err != nil {
		logger.Sugar().Errorw("validate", err.Error())
		return status.Error(codes.InvalidArgument, err.Error())
	}

	err = banapp.Validate(&mgrbanapp.BanAppReq{
		AppID:   info.ID,
		Message: info.BanMessage,
	})
	if err != nil {
		logger.Sugar().Errorw("validate", err.Error())
		return status.Error(codes.InvalidArgument, err.Error())
	}

	return nil
}
