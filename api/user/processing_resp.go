package user

import (
	"encoding/json"

	mw "github.com/NpoolPlatform/appuser-middleware/pkg/middleware/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/approle"
	"github.com/NpoolPlatform/message/npool/appusermw/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func InfoRowToObject(row *mw.Info) (*user.AppUserInfo, error) {
	var addressFields []string
	if row.AddressFields != "" {
		err := json.Unmarshal([]byte(row.AddressFields), &addressFields)
		if err != nil {
			logger.Sugar().Errorw("fail json unmarshal addressFields: %v", err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	signinVerifyByGoogleAuthentication := false
	if row.SigninVerifyByGoogleAuthentication == 1 {
		signinVerifyByGoogleAuthentication = true
	}

	googleAuthenticationVerified := false
	if row.GoogleAuthenticationVerified == 1 {
		googleAuthenticationVerified = true
	}

	isBanApp := false
	if row.BanAppUserID != "" {
		isBanApp = true
	}

	hasGoogleSecret := false
	if row.HasGoogleSecret != "" {
		hasGoogleSecret = true
	}

	roles := []*approle.AppRole{}
	for _, val := range row.Role {
		isDefault := false
		if val.Default == 1 {
			isDefault = true
		}
		roles = append(roles, &approle.AppRole{
			CreatedBy:   val.CreatedBy,
			Role:        val.Role,
			Description: val.Description,
			Default:     isDefault,
		})
	}

	return &user.AppUserInfo{
		ID:                                 row.ID,
		AppID:                              row.AppID,
		EmailAddress:                       row.EmailAddress,
		PhoneNO:                            row.PhoneNO,
		ImportFromApp:                      row.ImportFromApp,
		CreateAt:                           row.CreatedAt,
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
		SigninVerifyByGoogleAuthentication: signinVerifyByGoogleAuthentication,
		GoogleAuthenticationVerified:       googleAuthenticationVerified,
		IsBanApp:                           isBanApp,
		Roles:                              roles,
		HasGoogleSecret:                    hasGoogleSecret,
	}, nil
}
