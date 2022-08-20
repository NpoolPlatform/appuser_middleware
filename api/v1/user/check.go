package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	mgrappuser "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
	mgrappusercontrol "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"
	mgrappuserextra "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/appuser-manager/api/v2/appuser"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appusercontrol"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appuserextra"
	"github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/approle"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
)

func validate(ctx context.Context, info *npool.UserReq) error {
	err := appuser.Validate(&mgrappuser.AppUserReq{
		ID:            info.ID,
		AppID:         info.AppID,
		PhoneNO:       info.PhoneNO,
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
		AppID:              info.AppID,
		UserID:             info.ID,
		GoogleAuthVerified: info.GoogleAuthVerified,
		SigninVerifyType:   info.SigninVerifyType,
	})
	if err != nil {
		logger.Sugar().Errorw("validate", err.Error())
		return status.Error(codes.InvalidArgument, err.Error())
	}

	if info.PasswordHash == nil && info.ThirdPartyID == nil {
		logger.Sugar().Errorw("validate", "PasswordHash or ThirdPartyID must be one")
		return status.Error(codes.InvalidArgument, "PasswordHash or ThirdPartyID must be one")
	}

	for _, val := range info.RoleIDs {
		roleID, err := uuid.Parse(val)
		if err != nil {
			logger.Sugar().Errorw("validate", "RoleID", val, "error", err)
			return status.Error(codes.InvalidArgument, "RoleID is invalid")
		}

		_, err = approle.Row(ctx, roleID)
		if err != nil {
			if ent.IsNotFound(err) {
				logger.Sugar().Errorw("validate", "RoleID", val, "error", err)
				return status.Error(codes.NotFound, "RoleID no found")
			}
			logger.Sugar().Errorw("validate", "RoleID", val, "error", err)
			return status.Error(codes.Internal, err.Error())
		}
	}
	return nil
}

func Validate(ctx context.Context, info *npool.UserReq) error {
	return validate(ctx, info)
}
