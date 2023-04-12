package user

import (
	"context"

	"github.com/NpoolPlatform/appuser-manager/pkg/encrypt"

	appusersecretcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appusersecret"
	appuserthirdpartycrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuserthirdparty"

	appusersecretamgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusersecret"
	appuserthirdpartymgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserthirdparty"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	appusermgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
	appusercontrolmgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"
	appuserextramgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"

	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuser"
	appusercontrolcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appusercontrol"
	appuserextracrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuserextra"
	entappusercontrol "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusercontrol"
	entappuserextra "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"
	entappusersecret "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"
	entappuserthirdparty "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserthirdparty"

	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateAppUser(ctx context.Context, tx *ent.Tx) error {
	if _, err := appusercrud.UpdateSet(
		tx.AppUser.UpdateOneID(uuid.MustParse(*h.ID)),
		&appusermgrpb.AppUserReq{
			PhoneNO:       h.PhoneNO,
			EmailAddress:  h.EmailAddress,
			ImportFromApp: h.ImportedFromAppID,
		}).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateAppUserExtra(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppUserExtra.
		Query().
		Where(
			entappuserextra.AppID(uuid.MustParse(h.AppID)),
			entappuserextra.UserID(uuid.MustParse(*h.ID)),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	req := &appuserextramgrpb.AppUserExtraReq{
		AppID:         &h.AppID,
		UserID:        h.ID,
		FirstName:     h.FirstName,
		Birthday:      h.Birthday,
		LastName:      h.LastName,
		Gender:        h.Gender,
		Avatar:        h.Avatar,
		Username:      h.Username,
		PostalCode:    h.PostalCode,
		Age:           h.Age,
		Organization:  h.Organization,
		IDNumber:      h.IDNumber,
		AddressFields: h.AddressFields,
		ActionCredits: h.ActionCredits,
	}

	if info == nil {
		if _, err = appuserextracrud.CreateSet(
			tx.AppUserExtra.Create(),
			req,
		).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	if _, err = appuserextracrud.UpdateSet(info, req).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateAppUserControl(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppUserControl.
		Query().
		Where(
			entappusercontrol.AppID(uuid.MustParse(h.AppID)),
			entappusercontrol.UserID(uuid.MustParse(*h.ID)),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	req := &appusercontrolmgrpb.AppUserControlReq{
		AppID:              &h.AppID,
		UserID:             h.ID,
		GoogleAuthVerified: h.GoogleAuthVerified,
		SigninVerifyType:   h.SigninVerifyType,
		Kol:                h.Kol,
		KolConfirmed:       h.KolConfirmed,
	}

	if info == nil {
		if _, err := appusercontrolcrud.CreateSet(
			tx.AppUserControl.Create(),
			req,
		).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	if _, err = appusercontrolcrud.UpdateSet(info, req).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateAppUserSecret(ctx context.Context, tx *ent.Tx) error {
	var salt, password *string

	if h.PasswordHash != nil {
		saltStr := encrypt.Salt()
		salt = &saltStr

		passwordStr, err := encrypt.EncryptWithSalt(*h.PasswordHash, saltStr)
		if err != nil {
			return err
		}
		password = &passwordStr
	}

	info, err := tx.
		AppUserSecret.
		Query().
		Where(
			entappusersecret.AppID(uuid.MustParse(h.AppID)),
			entappusersecret.UserID(uuid.MustParse(*h.ID)),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	if _, err = appusersecretcrud.UpdateSet(
		info,
		&appusersecretamgrpb.AppUserSecretReq{
			PasswordHash: password,
			Salt:         salt,
			GoogleSecret: h.GoogleSecret,
		}).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateAppUserThirdParty(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppUserThirdParty.
		Query().
		Where(
			entappuserthirdparty.AppID(uuid.MustParse(h.AppID)),
			entappuserthirdparty.UserID(uuid.MustParse(*h.ID)),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	if info == nil && h.ThirdPartyID == nil {
		return nil
	}

	req := &appuserthirdpartymgrpb.AppUserThirdPartyReq{
		AppID:              &h.AppID,
		UserID:             h.ID,
		ThirdPartyID:       h.ThirdPartyID,
		ThirdPartyUserID:   h.ThirdPartyUserID,
		ThirdPartyUsername: h.ThirdPartyUsername,
		ThirdPartyAvatar:   h.ThirdPartyAvatar,
	}

	if info == nil {
		if _, err := appuserthirdpartycrud.CreateSet(
			tx.AppUserThirdParty.Create(),
			req,
		).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	if _, err = appuserthirdpartycrud.UpdateSet(info, req).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateUser(ctx context.Context) (*npool.User, error) {
	_, err := h.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateAppUser(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateAppUserExtra(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateAppUserControl(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateAppUserSecret(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateAppUserThirdParty(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetUser(ctx)
}
