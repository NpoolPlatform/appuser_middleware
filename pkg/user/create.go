package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/appuser-manager/pkg/encrypt"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-middleware/pkg/tracer/user"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	approleusermgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approleuser"
	appusermgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
	appusercontrolmgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"
	appuserextramgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"
	appusersecretamgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusersecret"
	appuserthirdpartymgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserthirdparty"

	approleusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/approleuser"
	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appuser"
	appusercontrolcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appusercontrol"
	appuserextracrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appuserextra"
	appusersecretcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appusersecret"
	appuserthirdpartycrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appuserthirdparty"

	uuid1 "github.com/NpoolPlatform/go-service-framework/pkg/const/uuid"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
)

//nolint:funlen
func CreateUser(ctx context.Context, in *npool.UserReq) (*npool.User, error) {
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

	importedFromAppID := uuid1.InvalidUUIDStr
	if in.ImportedFromAppID != nil {
		importedFromAppID = in.GetImportedFromAppID()
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info, err := appusercrud.CreateSet(tx.AppUser.Create(), &appusermgrpb.AppUserReq{
			ID:            in.ID,
			AppID:         in.AppID,
			PhoneNO:       in.PhoneNO,
			EmailAddress:  in.EmailAddress,
			ImportFromApp: &importedFromAppID,
		}).Save(ctx)
		if err != nil {
			logger.Sugar().Errorw("CreateUser", "error", err)
			return err
		}

		id = info.ID.String()
		appID = info.AppID.String()

		if _, err = appuserextracrud.CreateSet(tx.AppUserExtra.Create(), &appuserextramgrpb.AppUserExtraReq{
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
			ActionCredits: in.ActionCredits,
		}).Save(ctx); err != nil {
			logger.Sugar().Errorw("CreateUser", "error", err)
			return err
		}

		if _, err = appusercontrolcrud.CreateSet(tx.AppUserControl.Create(), &appusercontrolmgrpb.AppUserControlReq{
			AppID:              in.AppID,
			UserID:             in.ID,
			GoogleAuthVerified: in.GoogleAuthVerified,
			SigninVerifyType:   in.SigninVerifyType,
			Kol:                in.Kol,
		}).Save(ctx); err != nil {
			logger.Sugar().Errorw("CreateUser", "error", err)
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

			if _, err = appusersecretcrud.CreateSet(tx.AppUserSecret.Create(), &appusersecretamgrpb.AppUserSecretReq{
				AppID:        in.AppID,
				UserID:       in.ID,
				PasswordHash: password,
				Salt:         salt,
				GoogleSecret: in.GoogleSecret,
			}).Save(ctx); err != nil {
				logger.Sugar().Errorw("CreateUser", "error", err)
				return err
			}
		}

		if in.ThirdPartyID != nil {
			if _, err = appuserthirdpartycrud.CreateSet(tx.AppUserThirdParty.Create(), &appuserthirdpartymgrpb.AppUserThirdPartyReq{
				AppID:              in.AppID,
				UserID:             in.ID,
				ThirdPartyID:       in.ThirdPartyID,
				ThirdPartyUserID:   in.ThirdPartyUserID,
				ThirdPartyUsername: in.ThirdPartyUsername,
				ThirdPartyAvatar:   in.ThirdPartyAvatar,
			}).Save(ctx); err != nil {
				logger.Sugar().Errorw("CreateUser", "error", err)
				return err
			}
		}

		bulk := make([]*ent.AppRoleUserCreate, len(in.RoleIDs))
		for key := range in.RoleIDs {
			bulk[key] = approleusercrud.CreateSet(tx.AppRoleUser.Create(), &approleusermgrpb.AppRoleUserReq{
				AppID:  in.AppID,
				RoleID: &in.RoleIDs[key],
				UserID: in.ID,
			})
		}
		if _, err := tx.AppRoleUser.CreateBulk(bulk...).Save(ctx); err != nil {
			logger.Sugar().Errorw("CreateUser", "error", err)
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetUser(ctx, appID, id)
}
