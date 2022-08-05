package admin

import (
	"context"

	approleusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/approleuser"
	appusersecretcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appusersecret"
	"github.com/NpoolPlatform/appuser-manager/pkg/middleware/encrypt"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	"github.com/NpoolPlatform/appuser-middleware/pkg/user"

	approlecrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/approle"
	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appuser"
	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-middleware/pkg/tracer/admin"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool"
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approle"
	approleusermgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approleuser"
	appusermgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
	appusersecretmgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusersecret"
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/admin"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
)

func CreateGenesisUser(ctx context.Context, in *admin.CreateGenesisUserRequest) (*user.User, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateGenesisUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	span = commontracer.TraceInvoker(span, "app", "db", "CreateGenesisUser")
	roleInfo, err := approlecrud.RowOnly(ctx, &approle.Conds{
		AppID: &npool.StringVal{
			Value: in.GetAppID(),
			Op:    cruder.EQ,
		},
		Role: &npool.StringVal{
			Value: in.GetRole(),
			Op:    cruder.EQ,
		},
	})
	if err != nil {
		logger.Sugar().Errorw("CreateGenesisUser", "error", err)
		return nil, err
	}

	roleID := roleInfo.ID.String()

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := appusercrud.CreateSet(tx.AppUser.Create(), &appusermgrpb.AppUserReq{
			ID:           in.UserID,
			AppID:        in.AppID,
			EmailAddress: in.EmailAddress,
		}).Save(ctx); err != nil {
			logger.Sugar().Errorw("CreateGenesisUser", "error", err)
			return err
		}

		var password *string
		var salt *string

		if in.PasswordHash != nil {
			saltStr := encrypt.Salt()
			salt = &saltStr

			passwordStr, err := encrypt.EncryptWithSalt(in.GetPasswordHash(), saltStr)
			if err != nil {
				return err
			}
			password = &passwordStr
		}

		if _, err := appusersecretcrud.CreateSet(tx.AppUserSecret.Create(), &appusersecretmgrpb.AppUserSecretReq{
			AppID:        in.AppID,
			UserID:       in.UserID,
			PasswordHash: password,
			Salt:         salt,
		}).Save(ctx); err != nil {
			logger.Sugar().Errorw("CreateGenesisUser", "error", err)
			return err
		}

		if _, err := approleusercrud.CreateSet(tx.AppRoleUser.Create(), &approleusermgrpb.AppRoleUserReq{
			AppID:  in.AppID,
			RoleID: &roleID,
			UserID: in.UserID,
		}).Save(ctx); err != nil {
			logger.Sugar().Errorw("CreateGenesisUser", "error", err)
			return err
		}

		return nil
	})

	return user.GetUser(ctx, in.GetAppID(), in.GetUserID())
}
