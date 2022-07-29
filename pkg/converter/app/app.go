package app

import (
	"encoding/json"

	mapp "github.com/NpoolPlatform/appuser-middleware/pkg/app"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
)

func Ent2Grpc(row *mapp.App) (*npool.App, error) {
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
		BanAppID:           row.BanAppID.String(),
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
