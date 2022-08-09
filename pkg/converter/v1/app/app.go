package app

import (
	"encoding/json"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
)

func Ent2Grpc(row *npool.App) (*npool.App, error) {
	if row == nil {
		return nil, nil
	}

	methods := []string{}
	if row.SignupMethodsString != "" {
		err := json.Unmarshal([]byte(row.SignupMethodsString), &methods)
		if err != nil {
			return nil, err
		}
	}

	emethods := []string{}
	if row.ExtSigninMethodsString != "" {
		err := json.Unmarshal([]byte(row.ExtSigninMethodsString), &emethods)
		if err != nil {
			return nil, err
		}
	}

	row.SignupMethods = methods
	row.ExtSigninMethods = emethods
	row.KycEnable = row.KycEnableInt != 0
	row.SigninVerifyEnable = row.SigninVerifyEnableInt != 0
	row.InvitationCodeMust = row.InvitationCodeMustInt != 0

	return row, nil
}
