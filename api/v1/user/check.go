package user

import (
	"context"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/NpoolPlatform/appuser-manager/api/appuser"
	"github.com/NpoolPlatform/appuser-manager/api/appusercontrol"
	"github.com/NpoolPlatform/appuser-manager/api/appuserextra"
	"github.com/NpoolPlatform/appuser-manager/pkg/crud/approle"
	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuser"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npoolpb "github.com/NpoolPlatform/message/npool"
	appuserpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
	mgrappusercontrol "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"
	mgrappuserextra "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//nolint:dupl,funlen,gocyclo
func validate(ctx context.Context, info *npool.UserReq) error {
	err := appuser.Validate(&appuserpb.AppUserReq{
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

	if info.PasswordHash == nil && info.ThirdPartyID == nil && info.GetPasswordHash() == "" && info.GetThirdPartyID() == "" {
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

	if info.PhoneNO != nil && info.GetPhoneNO() != "" {
		phoneExist, err := appusercrud.ExistConds(ctx, &appuserpb.Conds{
			PhoneNO: &npoolpb.StringVal{
				Op:    cruder.EQ,
				Value: info.GetPhoneNO(),
			},
			AppID: &npoolpb.StringVal{
				Op:    cruder.EQ,
				Value: info.GetAppID(),
			},
		})

		if err != nil {
			logger.Sugar().Errorw("validate", "err", err)
			return status.Error(codes.Internal, err.Error())
		}

		if phoneExist {
			logger.Sugar().Errorw("validate", "phoneExsit", phoneExist)
			return status.Error(codes.AlreadyExists, "phone already exists")
		}
	}

	if info.EmailAddress != nil && info.GetEmailAddress() != "" {
		emailExist, err := appusercrud.ExistConds(ctx, &appuserpb.Conds{
			EmailAddress: &npoolpb.StringVal{
				Op:    cruder.EQ,
				Value: info.GetEmailAddress(),
			},
			AppID: &npoolpb.StringVal{
				Op:    cruder.EQ,
				Value: info.GetAppID(),
			},
		})
		if err != nil {
			logger.Sugar().Errorw("validate", "err", err)
			return status.Error(codes.Internal, err.Error())
		}

		if emailExist {
			logger.Sugar().Errorw("validate", "emailExist", emailExist)
			return status.Error(codes.AlreadyExists, "email already exists")
		}
	}

	return nil
}

func Validate(ctx context.Context, info *npool.UserReq) error {
	return validate(ctx, info)
}
