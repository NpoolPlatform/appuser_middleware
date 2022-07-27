package admin

import (
	"context"
	"fmt"

	approleusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approleuserv2"
	approlecrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approlev2"
	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuserv2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	bconst "github.com/NpoolPlatform/appuser-middleware/pkg/const"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/approle"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/approleuser"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/appuser"
	"github.com/NpoolPlatform/message/npool/appusermw/admin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
)

func CreateGenesisRoleUser(ctx context.Context, in *admin.CreateGenesisRoleUserRequest) (*appuser.AppUser, *approleuser.AppRoleUser, error) {
	var err error
	var userInfo *ent.AppUser

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateGenesisRoleUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	role, err := getRole(ctx)
	if err != nil {
		logger.Sugar().Error("fail get role:%v", err)
		return nil, nil, err
	}

	roleUser, err := getRoleUser(ctx, in.GetUser().GetAppID(), role.ID.String())
	if err != nil {
		logger.Sugar().Error("fail get role user:%v", err)
		return nil, nil, err
	}

	userInfo, err = getUser(ctx, in.GetUser().GetAppID(), in.GetUser().GetEmailAddress())
	if err != nil {
		logger.Sugar().Error("fail get user:%v", err)
		return nil, nil, err
	}

	userID := uuid.New()
	if userInfo != nil {
		userID = userInfo.ID
	}

	importFromApp := uuid.UUID{}
	if in.GetUser().GetImportFromApp() != "" {
		importFromApp = uuid.MustParse(in.GetUser().GetImportFromApp())
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if userInfo == nil {
			userTx := tx.AppUser.Create()
			userTx.SetID(userID)
			userTx.SetAppID(uuid.MustParse(in.GetUser().GetAppID()))
			userTx.SetImportFromApp(importFromApp)
			userTx.SetEmailAddress(in.GetUser().GetEmailAddress())
			userTx.SetPhoneNo(in.GetUser().GetPhoneNo())
			userInfo, err = userTx.Save(ctx)
			if err != nil {
				logger.Sugar().Error("fail create user:%v", err)
				return err
			}

			secretTx := tx.AppUserSecret.Create()
			secretTx.SetAppID(uuid.MustParse(in.GetUser().GetAppID()))
			secretTx.SetUserID(userID)
			secretTx.SetGoogleSecret(in.GetSecret().GetGoogleSecret())
			secretTx.SetPasswordHash(in.GetSecret().GetPasswordHash())
			secretTx.SetSalt(in.GetSecret().GetSalt())
			_, err = secretTx.Save(ctx)
			if err != nil {
				logger.Sugar().Error("fail create secret:%v", err)
				return err
			}
		}
		roleUserTx := tx.AppRoleUser.Create()
		roleUserTx.SetAppID(uuid.MustParse(in.GetUser().GetAppID()))
		roleUserTx.SetUserID(userID)
		roleUserTx.SetRoleID(role.ID)
		_, err = roleUserTx.Save(ctx)
		if err != nil {
			logger.Sugar().Error("fail create role user:%v", err)
			return err
		}
		return nil
	})

	if err != nil {
		logger.Sugar().Error("transaction fail :%v", err)
		return nil, nil, err
	}

	return &appuser.AppUser{
			PhoneNo:       userInfo.PhoneNo,
			ImportFromApp: userInfo.ImportFromApp.String(),
			ID:            userInfo.ID.String(),
			AppID:         userInfo.AppID.String(),
			EmailAddress:  userInfo.EmailAddress,
		}, &approleuser.AppRoleUser{
			UserID: roleUser.UserID.String(),
			ID:     roleUser.ID.String(),
			AppID:  roleUser.AppID.String(),
			RoleID: roleUser.RoleID.String(),
		}, nil
}

func getRole(ctx context.Context) (*ent.AppRole, error) {
	roleInfo, err := approlecrud.RowOnly(ctx, &approle.Conds{
		AppID: &npool.StringVal{
			Value: uuid.UUID{}.String(),
			Op:    cruder.EQ,
		},
		Role: &npool.StringVal{
			Value: bconst.GenesisRole,
			Op:    cruder.EQ,
		},
	})
	if err != nil {
		logger.Sugar().Error("fail get role:%v", err)
		return nil, err
	}
	return roleInfo, nil
}

func getRoleUser(ctx context.Context, appID, roleID string) (*ent.AppRoleUser, error) {
	roleUser, err := approleusercrud.RowOnly(ctx, &approleuser.Conds{
		AppID: &npool.StringVal{
			Value: appID,
			Op:    cruder.EQ,
		},
		RoleID: &npool.StringVal{
			Value: roleID,
			Op:    cruder.EQ,
		},
	})
	if err != nil {
		if ent.IsNotFound(err) {
			logger.Sugar().Error("genesis user already exist")
			return nil, fmt.Errorf("genesis user already exist")
		}
		logger.Sugar().Error("fail get role user:%v", err)
		return nil, err
	}
	return roleUser, nil
}

func getUser(ctx context.Context, appID, emailAddress string) (*ent.AppUser, error) {
	userInfo, err := appusercrud.RowOnly(ctx, &appuser.Conds{
		AppID: &npool.StringVal{
			Value: appID,
			Op:    cruder.EQ,
		},
		EmailAddress: &npool.StringVal{
			Value: emailAddress,
			Op:    cruder.EQ,
		},
	})
	if err != nil {
		if !ent.IsNotFound(err) {
			logger.Sugar().Error("fail get user:%v", err)
			return nil, err
		}
	}
	return userInfo, nil
}
