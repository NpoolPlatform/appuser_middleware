package app

import (
	"encoding/json"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"

	rcpt "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/recaptcha"
	sm "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"
)

func Ent2Grpc(row *npool.App) (*npool.App, error) {
	if row == nil {
		return nil, nil
	}

	methods := []string{}
	methods1 := []sm.SignMethodType{}

	if row.SignupMethodsStr != "" {
		err := json.Unmarshal([]byte(row.SignupMethodsStr), &methods)
		if err != nil {
			return nil, err
		}
	}
	for _, m := range methods {
		methods1 = append(methods1, sm.SignMethodType(sm.SignMethodType_value[m]))
	}

	emethods := []string{}
	emethods1 := []sm.SignMethodType{}

	if row.ExtSigninMethodsStr != "" {
		err := json.Unmarshal([]byte(row.ExtSigninMethodsStr), &emethods)
		if err != nil {
			return nil, err
		}
	}
	for _, m := range emethods {
		emethods1 = append(emethods1, sm.SignMethodType(sm.SignMethodType_value[m]))
	}

	row.SignupMethods = methods1
	row.ExtSigninMethods = emethods1
	row.KycEnable = row.KycEnableInt != 0
	row.SigninVerifyEnable = row.SigninVerifyEnableInt != 0
	row.InvitationCodeMust = row.InvitationCodeMustInt != 0
	row.RecaptchaMethod = rcpt.RecaptchaType(rcpt.RecaptchaType_value[row.RecaptchaMethodStr])

	return row, nil
}

func Ent2GrpcMany(rows []*npool.App) ([]*npool.App, error) {
	apps := []*npool.App{}
	for _, row := range rows {
		app, err := Ent2Grpc(row)
		if err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}
	return apps, nil
}
