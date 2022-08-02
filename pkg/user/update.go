package user

import (
	"context"

	appusersecretcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appusersecret"
	appuserthirdpartycrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appuserthirdparty"
	"github.com/NpoolPlatform/appuser-manager/pkg/middleware/encrypt"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-middleware/pkg/tracer/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	appusersecretamgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusersecret"
	appuserthirdpartymgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserthirdparty"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	appusermgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
	appusercontrolmgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"
	appuserextramgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"

	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appuser"
	appusercontrolcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appusercontrol"
	appuserextracrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appuserextra"
)

func UpdateUser(ctx context.Context, in *npool.UserReq) (*User, error) {
	var id string
	var appID string
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	span = commontracer.TraceInvoker(span, "user", "db", "UpdateTx")

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info, err := appusercrud.UpdateTx(tx, &appusermgrpb.AppUserReq{
			PhoneNo:       in.PhoneNO,
			EmailAddress:  in.EmailAddress,
			ImportFromApp: in.ImportedFromAppID,
		}).Save(ctx)
		if err != nil {
			logger.Sugar().Errorw("update app user", "err", err.Error())
			return err
		}

		id = info.ID.String()
		appID = info.AppID.String()

		if _, err = appuserextracrud.UpdateTx(tx, &appuserextramgrpb.AppUserExtraReq{
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
			return err
		}

		if _, err = appusercontrolcrud.UpdateTx(tx, &appusercontrolmgrpb.AppUserControlReq{
			AppID:                              in.AppID,
			UserID:                             in.ID,
			SigninVerifyByGoogleAuthentication: in.SigninVerifyByGoogleAuth,
			GoogleAuthenticationVerified:       in.GoogleAuthenticationVerified,
		}).Save(ctx); err != nil {
			logger.Sugar().Errorw("update app user control", "err", err.Error())
			return err
		}

		var password *string
		var salt *string

		if in.PasswordHash != nil {
			saltStr := encrypt.Salt()
			salt = &saltStr

			passwordStr, err := encrypt.EncryptWithSalt(in.GetPasswordHash(), saltStr)
			if err != nil {
				logger.Sugar().Errorw("make password", "err", err.Error())
				return err
			}
			password = &passwordStr
		}

		if _, err = appusersecretcrud.UpdateTx(tx, &appusersecretamgrpb.AppUserSecretReq{
			PasswordHash: password,
			Salt:         salt,
			GoogleSecret: in.GoogleSecret,
		}).Save(ctx); err != nil {
			logger.Sugar().Errorw("update app user secret", "err", err.Error())
			return err
		}

		if _, err = appuserthirdpartycrud.UpdateTx(tx, &appuserthirdpartymgrpb.AppUserThirdPartyReq{
			ThirdPartyID:         in.ThirdPartyID,
			ThirdPartyUserID:     in.ThirdPartyUserID,
			ThirdPartyUsername:   in.ThirdPartyUsername,
			ThirdPartyUserAvatar: in.ThirdPartyUserAvatar,
		}).Save(ctx); err != nil {
			logger.Sugar().Errorw("update app user third party", "err", err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetUser(ctx, appID, id)
}
