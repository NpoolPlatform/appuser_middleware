package user

import (
	"context"

	"github.com/google/uuid"

	appusersecretcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appusersecret"
	appuserthirdpartycrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appuserthirdparty"
	"github.com/NpoolPlatform/appuser-manager/pkg/encrypt"
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
	entappusercontrol "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusercontrol"
	entappuserextra "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"
	entappusersecret "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"
	entappuserthirdparty "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserthirdparty"
)

//nolint:funlen
func UpdateUser(ctx context.Context, in *npool.UserReq) (*npool.User, error) {
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
		_, err := appusercrud.UpdateSet(
			tx.AppUser.UpdateOneID(uuid.MustParse(in.GetID())),
			&appusermgrpb.AppUserReq{
				PhoneNO:       in.PhoneNO,
				EmailAddress:  in.EmailAddress,
				ImportFromApp: in.ImportedFromAppID,
			}).Save(ctx)
		if err != nil {
			logger.Sugar().Errorw("UpdateUser", "err", err.Error())
			return err
		}

		extra, err := tx.
			AppUserExtra.
			Query().
			Where(
				entappuserextra.AppID(uuid.MustParse(in.GetAppID())),
				entappuserextra.UserID(uuid.MustParse(in.GetID())),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		if _, err = appuserextracrud.UpdateSet(
			extra,
			&appuserextramgrpb.AppUserExtraReq{
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

		ctrl, err := tx.
			AppUserControl.
			Query().
			Where(
				entappusercontrol.AppID(uuid.MustParse(in.GetAppID())),
				entappusercontrol.UserID(uuid.MustParse(in.GetID())),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		if _, err = appusercontrolcrud.UpdateSet(
			ctrl,
			&appusercontrolmgrpb.AppUserControlReq{
				GoogleAuthVerified: in.GoogleAuthVerified,
				SigninVerifyType:   in.SigninVerifyType,
			}).Save(ctx); err != nil {
			logger.Sugar().Errorw("UpdateUser", "err", err.Error())
			return err
		}

		var password *string
		var salt *string

		if in.PasswordHash != nil {
			saltStr := encrypt.Salt()
			salt = &saltStr

			passwordStr, err := encrypt.EncryptWithSalt(in.GetPasswordHash(), saltStr)
			if err != nil {
				logger.Sugar().Errorw("UpdateUser", "err", err.Error())
				return err
			}
			password = &passwordStr
		}

		secret, err := tx.
			AppUserSecret.
			Query().
			Where(
				entappusersecret.AppID(uuid.MustParse(in.GetAppID())),
				entappusersecret.UserID(uuid.MustParse(in.GetID())),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		if _, err = appusersecretcrud.UpdateSet(
			secret,
			&appusersecretamgrpb.AppUserSecretReq{
				AppID:        in.AppID,
				PasswordHash: password,
				Salt:         salt,
				GoogleSecret: in.GoogleSecret,
			}).Save(ctx); err != nil {
			logger.Sugar().Errorw("UpdateUser", "err", err.Error())
			return err
		}

		thirdparty, err := tx.
			AppUserThirdParty.
			Query().
			Where(
				entappuserthirdparty.AppID(uuid.MustParse(in.GetAppID())),
				entappuserthirdparty.UserID(uuid.MustParse(in.GetID())),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		if _, err = appuserthirdpartycrud.UpdateSet(
			thirdparty,
			&appuserthirdpartymgrpb.AppUserThirdPartyReq{
				AppID:              in.AppID,
				ThirdPartyID:       in.ThirdPartyID,
				ThirdPartyUserID:   in.ThirdPartyUserID,
				ThirdPartyUsername: in.ThirdPartyUsername,
				ThirdPartyAvatar:   in.ThirdPartyAvatar,
			}).Save(ctx); err != nil {
			logger.Sugar().Errorw("UpdateUser", "err", err.Error())
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetUser(ctx, in.GetAppID(), in.GetID())
}
