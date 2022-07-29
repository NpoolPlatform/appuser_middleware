package user

import (
	"encoding/json"
	mapp "github.com/NpoolPlatform/appuser-middleware/pkg/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func QueryEnt2Grpc(row *mapp.UseQueryrResp) (*npool.User, error) {
	var addressFields []string
	if row.AddressFields != "" {
		err := json.Unmarshal([]byte(row.AddressFields), &addressFields)
		if err != nil {
			logger.Sugar().Errorw("fail json unmarshal addressFields: %v", err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	banned := false
	if row.BanAppUserID != "" {
		banned = true
	}

	hasGoogleSecret := false
	if row.HasGoogleSecret != "" {
		hasGoogleSecret = true
	}

	roles := []string{}
	for _, val := range row.Role {
		roles = append(roles, val.Role)
	}

	return &npool.User{
		ID:                                 row.ID,
		AppID:                              row.AppID,
		EmailAddress:                       row.EmailAddress,
		PhoneNO:                            row.PhoneNO,
		ImportedFromAppID:                  "",
		ImportedFromAppName:                "",
		ImportedFromAppLogo:                "",
		ImportedFromAppHome:                "",
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
		Banned:                             banned,
		BanMessage:                         row.BanAppUserMessage,
		HasGoogleSecret:                    hasGoogleSecret,
		Roles:                              roles,
		Logined:                            false,
		LoginAccount:                       "",
		LoginAccountType:                   "",
		LoginToken:                         "",
		LoginClientIP:                      "",
		LoginClientUserAgent:               "",
		CreateAt:                           row.CreatedAt,
	}, nil
}
