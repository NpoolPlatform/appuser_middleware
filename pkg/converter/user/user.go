package user

import (
	"encoding/json"

	mapp "github.com/NpoolPlatform/appuser-middleware/pkg/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
)

func Ent2Grpc(row *mapp.User) *npool.User {
	if row == nil {
		return nil
	}

	addressFields := []string{}
	_ = json.Unmarshal([]byte(row.AddressFields), &addressFields)

	return &npool.User{
		ID:                                 row.ID.String(),
		AppID:                              row.AppID.String(),
		EmailAddress:                       row.EmailAddress,
		PhoneNO:                            row.PhoneNO,
		ImportedFromAppID:                  row.ImportedFromAppID.String(),
		ImportedFromAppName:                row.ImportedFromAppName,
		ImportedFromAppLogo:                row.ImportedFromAppLogo,
		ImportedFromAppHome:                row.ImportedFromAppHome,
		Username:                           row.Username,
		AddressFields:                      addressFields,
		Gender:                             row.Gender,
		PostalCode:                         row.PostalCode,
		Age:                                row.Age,
		Birthday:                           row.Birthday,
		Avatar:                             row.Avatar,
		Organization:                       row.Organization,
		FirstName:                          row.FirstName,
		LastName:                           row.LastName,
		IDNumber:                           row.IDNumber,
		SigninVerifyByGoogleAuthentication: row.SigninVerifyByGoogleAuthentication != 0,
		GoogleAuthenticationVerified:       row.GoogleAuthenticationVerified != 0,
		Banned:                             row.Banned,
		BanMessage:                         row.BanMessage,
		HasGoogleSecret:                    row.HasGoogleSecret,
		Roles:                              row.Roles,
		Logined:                            false,
		LoginAccount:                       "",
		LoginAccountType:                   "",
		LoginToken:                         "",
		LoginClientIP:                      "",
		LoginClientUserAgent:               "",
		CreateAt:                           row.CreatedAt,
	}
}

func Ent2GrpcMany(rows []*mapp.User) []*npool.User {
	users := []*npool.User{}
	for _, row := range rows {
		users = append(users, Ent2Grpc(row))
	}
	return users
}
