package user

import (
	"context"

	"github.com/NpoolPlatform/appuser-manager/pkg/middleware/encrypt"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-middleware/pkg/tracer/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	appusermgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
	appusercontrolmgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"
	appuserextramgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"
	appusersecretamgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusersecret"
	appuserthirdpartymgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserthirdparty"

	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appuser"
	appusercontrolcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appusercontrol"
	appuserextracrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appuserextra"
	appusersecretcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appusersecret"
	appuserthirdpartycrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appuserthirdparty"
)

//nolint:funlen
func CreateUser(ctx context.Context, in *npool.UserReq) (*User, error) {
	var id string
	var appID string
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	span = commontracer.TraceInvoker(span, "user", "db", "CreateTx")

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info, err := appusercrud.CreateTx(tx, &appusermgrpb.AppUserReq{
			ID:            in.ID,
			AppID:         in.AppID,
			PhoneNo:       in.PhoneNO,
			EmailAddress:  in.EmailAddress,
			ImportFromApp: in.ImportedFromAppID,
		}).Save(ctx)
		if err != nil {
			logger.Sugar().Errorw("app user create", "error", err)
			return err
		}

		id = info.ID.String()
		appID = info.AppID.String()

		if _, err = appuserextracrud.CreateTx(tx, &appuserextramgrpb.AppUserExtraReq{
			AppID:         in.AppID,
			UserID:        in.ID,
			FirstName:     in.FirstName,
			Birthday:      in.Birthday,
			LastName:      in.LastName,
			Gender:        in.Gender,
			Avatar:        in.Avatar,
			Username:      in.Username,
			PostalCode:    in.PostalCode,
			Age:           in.Age,
			Organization:  in.Organization,
			IDNumber:      in.IDNumber,
			AddressFields: in.AddressFields,
		}).Save(ctx); err != nil {
			logger.Sugar().Errorw("app user extra create", "error", err)
			return err
		}

		if _, err = appusercontrolcrud.CreateTx(tx, &appusercontrolmgrpb.AppUserControlReq{
			AppID:                              in.AppID,
			UserID:                             in.ID,
			SigninVerifyByGoogleAuthentication: in.SigninVerifyByGoogleAuth,
			GoogleAuthenticationVerified:       in.GoogleAuthenticationVerified,
		}).Save(ctx); err != nil {
			logger.Sugar().Errorw("app user control create", "error", err)
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

		if _, err = appusersecretcrud.CreateTx(tx, &appusersecretamgrpb.AppUserSecretReq{
			AppID:        in.AppID,
			UserID:       in.ID,
			PasswordHash: password,
			Salt:         salt,
			GoogleSecret: in.GoogleSecret,
		}).Save(ctx); err != nil {
			logger.Sugar().Errorw("app user secret create", "error", err)
			return err
		}

		if _, err = appuserthirdpartycrud.CreateTx(tx, &appuserthirdpartymgrpb.AppUserThirdPartyReq{
			AppID:                in.AppID,
			UserID:               in.ID,
			ThirdPartyID:         in.ThirdPartyID,
			ThirdPartyUserID:     in.ThirdPartyUserID,
			ThirdPartyUsername:   in.ThirdPartyUsername,
			ThirdPartyUserAvatar: in.ThirdPartyUserAvatar,
		}).Save(ctx); err != nil {
			logger.Sugar().Errorw("app user third party create", "error", err)
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetUser(ctx, appID, id)
}
