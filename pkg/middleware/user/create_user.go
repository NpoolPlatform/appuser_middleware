package user

import (
	"context"
	"fmt"

	approlecrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approlev2"
	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/approle"
	"github.com/NpoolPlatform/message/npool/appusermw/user"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
)

func CreateUserWithSecret(ctx context.Context, in *user.CreateUserWithSecretRequest) (*ent.AppUser, error) {
	var err error
	var userInfo *ent.AppUser

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateUserWithSecret")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	userID := uuid.New()
	if in.User.ID != nil && in.GetUser().GetID() != "" {
		userID = uuid.MustParse(in.GetUser().GetID())
	}

	importFromApp := uuid.UUID{}
	if in.User.ImportFromApp != nil && in.GetUser().GetImportFromApp() != "" {
		importFromApp = uuid.MustParse(in.GetUser().GetImportFromApp())
	}

	defaultRole, err := approlecrud.RowOnly(ctx, &approle.Conds{
		AppID: &npool.StringVal{
			Value: in.GetUser().GetAppID(),
			Op:    cruder.EQ,
		},
	})
	if err != nil {
		logger.Sugar().Errorw("fail get default role: %v", err)
		return nil, fmt.Errorf("fail get default role: %v", err)
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		userTx := tx.AppUser.Create()
		userTx.SetID(userID)
		userTx.SetAppID(uuid.MustParse(in.GetUser().GetAppID()))
		userTx.SetImportFromApp(importFromApp)
		userTx.SetEmailAddress(in.GetUser().GetEmailAddress())
		userTx.SetPhoneNo(in.GetUser().GetPhoneNo())
		userInfo, err = userTx.Save(ctx)
		if err != nil {
			logger.Sugar().Errorw("fail create user: %v", err)
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
			logger.Sugar().Errorw("fail create secret: %v", err)
			return err
		}

		roleUserTx := tx.AppRoleUser.Create()
		roleUserTx.SetAppID(uuid.MustParse(in.GetUser().GetAppID()))
		roleUserTx.SetUserID(userID)
		roleUserTx.SetRoleID(defaultRole.ID)
		_, err = roleUserTx.Save(ctx)
		if err != nil {
			logger.Sugar().Errorw("fail create role user: %v", err)
			return err
		}

		return nil
	})

	if err != nil {
		logger.Sugar().Error("transaction fail :%v", err)
		return nil, err
	}

	return userInfo, nil
}

func CreateUserWithThirdParty(ctx context.Context, in *user.CreateUserWithThirdPartyRequest) (*ent.AppUser, error) {
	var err error
	var userInfo *ent.AppUser

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateUserWithThirdParty")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	userID := uuid.New()
	if in.User.ID != nil && in.GetUser().GetID() != "" {
		userID = uuid.MustParse(in.GetUser().GetID())
	}

	importFromApp := uuid.UUID{}
	if in.User.ImportFromApp != nil && in.GetUser().GetImportFromApp() != "" {
		importFromApp = uuid.MustParse(in.GetUser().GetImportFromApp())
	}

	defaultRole, err := approlecrud.RowOnly(ctx, &approle.Conds{
		AppID: &npool.StringVal{
			Value: in.GetUser().GetAppID(),
			Op:    cruder.EQ,
		},
	})
	if err != nil {
		logger.Sugar().Error("fail get role :%v", err)
		return nil, err
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
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

		secretTx := tx.AppUserThirdParty.Create()
		secretTx.SetAppID(uuid.MustParse(in.GetUser().GetAppID()))
		secretTx.SetUserID(userID)
		secretTx.SetThirdPartyID(in.GetThirdParty().GetThirdPartyID())
		secretTx.SetThirdPartyUserAvatar(in.GetThirdParty().GetThirdPartyUserAvatar())
		secretTx.SetThirdPartyUserID(in.GetThirdParty().GetThirdPartyUserID())
		secretTx.SetThirdPartyUsername(in.GetThirdParty().GetThirdPartyUsername())
		_, err = secretTx.Save(ctx)
		if err != nil {
			logger.Sugar().Error("fail create third party:%v", err)
			return err
		}

		roleUserTx := tx.AppRoleUser.Create()
		roleUserTx.SetAppID(uuid.MustParse(in.GetUser().GetAppID()))
		roleUserTx.SetUserID(userID)
		roleUserTx.SetRoleID(defaultRole.ID)
		_, err = roleUserTx.Save(ctx)
		if err != nil {
			logger.Sugar().Error("fail create role user:%v", err)
			return err
		}

		return nil
	})

	if err != nil {
		logger.Sugar().Error("transaction fail :%v", err)
		return nil, err
	}

	return userInfo, nil
}
