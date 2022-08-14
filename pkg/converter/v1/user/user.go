package user

import (
	"encoding/json"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
)

func Ent2Grpc(row *npool.User) *npool.User {
	if row == nil {
		return nil
	}

	addressFields := []string{}
	_ = json.Unmarshal([]byte(row.AddressFieldsString), &addressFields)

	row.AddressFields = addressFields
	row.SigninVerifyByGoogleAuth = row.SigninVerifyByGoogleAuthInt != 0
	row.GoogleAuthVerified = row.GoogleAuthVerifiedInt != 0

	row.Banned = false
	if row.GetBanAppUserID() != "" {
		row.Banned = true
	}
	return row
}

func Ent2GrpcMany(rows []*npool.User) []*npool.User {
	users := []*npool.User{}
	for _, row := range rows {
		users = append(users, Ent2Grpc(row))
	}
	return users
}
