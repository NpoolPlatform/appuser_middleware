package app

import (
	"encoding/json"

	appmw "github.com/NpoolPlatform/appuser-middleware/pkg/middleware/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/appusermw/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func appInfoRowToObject(row *appmw.Info) (*app.AppInfo, error) {
	isBanApp := false
	if row.BanAppAppID != "" {
		isBanApp = true
	}

	var signupMethods []string
	err := json.Unmarshal([]byte(row.SignupMethods), &signupMethods)
	if err != nil {
		logger.Sugar().Errorw("json unmarshal fail: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	var externSigninMethods []string
	err = json.Unmarshal([]byte(row.ExternSigninMethods), &externSigninMethods)
	if err != nil {
		logger.Sugar().Errorw("json unmarshal fail: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	kycEnable := false
	if row.KycEnable == 1 {
		kycEnable = true
	}

	signinVerifyEnable := false
	if row.SigninVerifyEnable == 1 {
		signinVerifyEnable = true
	}

	invitationCodeMust := false
	if row.InvitationCodeMust == 1 {
		invitationCodeMust = true
	}

	return &app.AppInfo{
		Id:                  row.ID,
		CreatedBy:           row.CreatedBy,
		Name:                row.Name,
		Logo:                row.Logo,
		CreatedAt:           row.CreatedAt,
		Description:         row.Description,
		IsBanApp:            isBanApp,
		BanAppMessage:       row.BanAppMessage,
		SignupMethods:       signupMethods,
		ExternSigninMethods: externSigninMethods,
		RecaptchaMethod:     row.RecaptchaMethod,
		KycEnable:           kycEnable,
		SigninVerifyEnable:  signinVerifyEnable,
		InvitationCodeMust:  invitationCodeMust,
	}, nil
}
