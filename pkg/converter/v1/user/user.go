package user

import (
	"encoding/json"

	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/kyc"
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
)

func Ent2Grpc(row *npool.User) *npool.User {
	if row == nil {
		return nil
	}

	addressFields := []string{}
	_ = json.Unmarshal([]byte(row.AddressFieldsString), &addressFields)

	row.AddressFields = addressFields
	row.GoogleAuthVerified = row.GoogleAuthVerifiedInt != 0

	row.Banned = false
	if row.GetBanAppUserID() != "" {
		row.Banned = true
	}

	row.SigninVerifyType = signmethod.SignMethodType(signmethod.SignMethodType_value[row.SigninVerifyTypeStr])
	row.KycReviewState = kyc.KycReviewState(kyc.KycReviewState_value[row.KycReviewStateStr])
	return row
}

func Ent2GrpcMany(rows []*npool.User) []*npool.User {
	users := []*npool.User{}
	for _, row := range rows {
		users = append(users, Ent2Grpc(row))
	}
	return users
}
