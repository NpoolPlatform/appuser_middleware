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

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUserInfo")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	role, err := approlecrud.RowOnly(ctx, &approle.Conds{
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
		return nil, nil, err
	}

	roleUser, err := approleusercrud.RowOnly(ctx, &approleuser.Conds{
		AppID: &npool.StringVal{
			Value: in.GetUser().GetAppID(),
			Op:    cruder.EQ,
		},
		RoleID: &npool.StringVal{
			Value: role.ID.String(),
			Op:    cruder.EQ,
		},
	})
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil, fmt.Errorf("genesis user already exist")
		}
		return nil, nil, err
	}

	user, err := appusercrud.RowOnly(ctx, &appuser.Conds{
		AppID: &npool.StringVal{
			Value: in.GetUser().GetAppID(),
			Op:    cruder.EQ,
		},
		EmailAddress: &npool.StringVal{
			Value: in.GetUser().GetEmailAddress(),
			Op:    cruder.EQ,
		},
	})
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, nil, err
		}
	}

	userId := uuid.New()

	importFromApp := uuid.UUID{}
	if in.GetUser().GetImportFromApp() != "" {
		importFromApp = uuid.MustParse(in.GetUser().GetImportFromApp())
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if user == nil {
			userTx := tx.AppUser.Create()
			userTx.SetID(userId)
			userTx.SetAppID(uuid.MustParse(in.GetUser().GetAppID()))
			userTx.SetImportFromApp(importFromApp)
			userTx.SetEmailAddress(in.GetUser().GetEmailAddress())
			userTx.SetPhoneNo(in.GetUser().GetPhoneNo())
			userInfo, err = userTx.Save(ctx)
			if err != nil {
				return err
			}

			secretTx := tx.AppUserSecret.Create()
			secretTx.SetAppID(uuid.MustParse(in.GetUser().GetAppID()))
			secretTx.SetUserID(userId)
			secretTx.SetGoogleSecret(in.GetSecret().GetGoogleSecret())
			secretTx.SetPasswordHash(in.GetSecret().GetPasswordHash())
			secretTx.SetSalt(in.GetSecret().GetSalt())
			_, err = secretTx.Save(ctx)
			if err != nil {
				return err
			}
		}
		roleUserTx := tx.AppRoleUser.Create()
		roleUserTx.SetAppID(uuid.MustParse(in.GetUser().GetAppID()))
		roleUserTx.SetUserID(userId)
		roleUserTx.SetRoleID(role.ID)
		_, err = roleUserTx.Save(ctx)
		if err != nil {
			return err
		}
		return nil
	})

	return &appuser.AppUser{
			PhoneNo:       user.PhoneNo,
			ImportFromApp: user.ImportFromApp.String(),
			ID:            user.ID.String(),
			AppID:         user.AppID.String(),
			EmailAddress:  user.EmailAddress,
		}, &approleuser.AppRoleUser{
			UserID: roleUser.UserID.String(),
			ID:     roleUser.ID.String(),
			AppID:  roleUser.AppID.String(),
			RoleID: roleUser.RoleID.String(),
		}, nil
}
