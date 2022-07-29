package app

import (
	"encoding/json"
	mapp "github.com/NpoolPlatform/appuser-middleware/pkg/app"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
)

func QueryEnt2Grpc(row *mapp.AppQueryResp) (*npool.App, error) {
	if row == nil {
		return nil, nil
	}

	methods := []string{}
	if row.SignupMethods != "" {
		err := json.Unmarshal([]byte(row.SignupMethods), &methods)
		if err != nil {
			return nil, err
		}
	}

	emethods := []string{}
	if row.ExtSigninMethods != "" {
		err := json.Unmarshal([]byte(row.ExtSigninMethods), &emethods)
		if err != nil {
			return nil, err
		}
	}

	return &npool.App{
		ID:                 row.ID.String(),
		CreatedBy:          row.CreatedBy.String(),
		Name:               row.Name,
		Logo:               row.Logo,
		Description:        row.Description,
		Banned:             row.Banned,
		BanMessage:         row.BanMessage,
		SignupMethods:      methods,
		ExtSigninMethods:   emethods,
		RecaptchaMethod:    row.RecaptchaMethod,
		KycEnable:          row.KycEnable != 0,
		SigninVerifyEnable: row.SigninVerifyEnable != 0,
		InvitationCodeMust: row.InvitationCodeMust != 0,
		CreatedAt:          row.CreatedAt,
	}, nil
}

func CreateEnt2Grpc(row *mapp.AppCreateResp) (*npool.App, error) {
	banned := false
	bannedMsg := ""
	if row.BanApp != nil {
		banned = true
		bannedMsg = row.BanApp.Message
	}

	return &npool.App{
		ID:          row.App.ID.String(),
		CreatedBy:   row.App.CreatedBy.String(),
		Name:        row.App.Name,
		Logo:        row.App.Logo,
		Description: row.App.Description,
		Banned:      banned,
		BanMessage:  bannedMsg,

		SignupMethods:    row.AppControl.SignupMethods,
		ExtSigninMethods: row.AppControl.ExternSigninMethods,

		RecaptchaMethod:    row.AppControl.RecaptchaMethod,
		KycEnable:          row.AppControl.KycEnable,
		SigninVerifyEnable: row.AppControl.SigninVerifyEnable,
		InvitationCodeMust: row.AppControl.InvitationCodeMust,

		CreatedAt: row.App.CreatedAt,
	}, nil
}
