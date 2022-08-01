package user

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	mgrappuser "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
	mgrappusercontrol "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"
	mgrappuserextra "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/appuser-manager/api/v2/appuser"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appusercontrol"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appuserextra"
)

func validate(info *npool.UserReq) error {
	err := appuser.Validate(&mgrappuser.AppUserReq{
		ID:            info.ID,
		AppID:         info.AppID,
		PhoneNo:       info.PhoneNO,
		EmailAddress:  info.EmailAddress,
		ImportFromApp: info.ImportedFromAppID,
	})
	if err != nil {
		logger.Sugar().Errorw("validate", err.Error())
		return status.Error(codes.InvalidArgument, err.Error())
	}

	err = appuserextra.Validate(&mgrappuserextra.AppUserExtraReq{
		AppID:         info.AppID,
		UserID:        info.ID,
		FirstName:     info.FirstName,
		Birthday:      info.Birthday,
		LastName:      info.LastName,
		Gender:        info.Gender,
		Avatar:        info.Avatar,
		Username:      info.Username,
		PostalCode:    info.PostalCode,
		Age:           info.Age,
		Organization:  info.Organization,
		IDNumber:      info.IDNumber,
		AddressFields: info.AddressFields,
	})
	if err != nil {
		logger.Sugar().Errorw("validate", err.Error())
		return status.Error(codes.InvalidArgument, err.Error())
	}

	err = appusercontrol.Validate(&mgrappusercontrol.AppUserControlReq{
		AppID:                              info.AppID,
		UserID:                             info.ID,
		SigninVerifyByGoogleAuthentication: info.SigninVerifyByGoogleAuth,
		GoogleAuthenticationVerified:       info.GoogleAuthenticationVerified,
	})
	if err != nil {
		logger.Sugar().Errorw("validate", err.Error())
		return status.Error(codes.InvalidArgument, err.Error())
	}

	err = appusercontrol.Validate(&mgrappusercontrol.AppUserControlReq{
		AppID:                              info.AppID,
		UserID:                             info.ID,
		SigninVerifyByGoogleAuthentication: info.SigninVerifyByGoogleAuth,
		GoogleAuthenticationVerified:       info.GoogleAuthenticationVerified,
	})
	if err != nil {
		logger.Sugar().Errorw("validate", err.Error())
		return status.Error(codes.InvalidArgument, err.Error())
	}

	return nil
}
