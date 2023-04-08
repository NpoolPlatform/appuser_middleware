package user

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-manager/pkg/encrypt"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	entsubscriber "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/subscriber"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	approleusermgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approleuser"
	appusermgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
	appusercontrolmgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"
	appuserextramgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"
	appusersecretmgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusersecret"
	appuserthirdpartymgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserthirdparty"

	approleusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/approleuser"
	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuser"
	appusercontrolcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appusercontrol"
	appuserextracrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuserextra"
	appusersecretcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appusersecret"
	appuserthirdpartycrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuserthirdparty"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createAppUser(ctx context.Context, tx *ent.Tx) error {
	if h.ID == nil {
		return fmt.Errorf("invalid id")
	}
	if h.PhoneNO == nil && h.EmailAddress == nil {
		return fmt.Errorf("invalid account")
	}

	if _, err := appusercrud.CreateSet(
		tx.AppUser.Create(),
		&appusermgrpb.AppUserReq{
			ID:            h.ID,
			AppID:         &h.AppID,
			PhoneNO:       h.PhoneNO,
			EmailAddress:  h.EmailAddress,
			ImportFromApp: h.ImportedFromAppID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createAppUserExtra(ctx context.Context, tx *ent.Tx) error {
	if _, err := appuserextracrud.CreateSet(
		tx.AppUserExtra.Create(),
		&appuserextramgrpb.AppUserExtraReq{
			AppID:  &h.AppID,
			UserID: h.ID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createAppUserControl(ctx context.Context, tx *ent.Tx) error {
	if _, err := appusercontrolcrud.CreateSet(
		tx.AppUserControl.Create(),
		&appusercontrolmgrpb.AppUserControlReq{
			AppID:  &h.AppID,
			UserID: h.ID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createAppUserSecret(ctx context.Context, tx *ent.Tx) error {
	if h.PasswordHash == nil {
		return fmt.Errorf("invalid password")
	}

	saltStr := encrypt.Salt()
	pwdStr, err := encrypt.EncryptWithSalt(*h.PasswordHash, saltStr)
	if err != nil {
		return err
	}

	if _, err := appusersecretcrud.CreateSet(
		tx.AppUserSecret.Create(),
		&appusersecretmgrpb.AppUserSecretReq{
			AppID:        &h.AppID,
			UserID:       h.ID,
			PasswordHash: &pwdStr,
			Salt:         &saltStr,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *createHandler) createAppUserThirdParty(ctx context.Context, tx *ent.Tx) error {
	if h.ThirdPartyID == nil {
		return nil
	}

	if _, err := appuserthirdpartycrud.CreateSet(
		tx.AppUserThirdParty.Create(),
		&appuserthirdpartymgrpb.AppUserThirdPartyReq{
			AppID:              &h.AppID,
			UserID:             h.ID,
			ThirdPartyID:       h.ThirdPartyID,
			ThirdPartyUserID:   h.ThirdPartyUserID,
			ThirdPartyUsername: h.ThirdPartyUsername,
			ThirdPartyAvatar:   h.ThirdPartyAvatar,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *createHandler) createAppRoleUser(ctx context.Context, tx *ent.Tx) error {
	if len(h.RoleIDs) == 0 {
		return nil
	}

	bulk := make([]*ent.AppRoleUserCreate, len(h.RoleIDs))
	for i, roleID := range h.RoleIDs {
		bulk[i] = approleusercrud.CreateSet(
			tx.AppRoleUser.Create(),
			&approleusermgrpb.AppRoleUserReq{
				AppID:  &h.AppID,
				RoleID: &roleID,
				UserID: h.ID,
			})
	}
	if _, err := tx.
		AppRoleUser.
		CreateBulk(bulk...).
		Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *createHandler) updateSubscriber(ctx context.Context, tx *ent.Tx) error {
	if h.EmailAddress == nil {
		return nil
	}

	sub, err := tx.
		Subscriber.
		Query().
		Where(
			entsubscriber.AppID(uuid.MustParse(h.AppID)),
			entsubscriber.EmailAddress(*h.EmailAddress),
		).
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}

	if _, err := tx.
		Subscriber.
		UpdateOneID(sub.ID).
		SetRegistered(true).
		Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) CreateUser(ctx context.Context) (info *npool.User, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := h.checkAccountExist(ctx); err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createAppUser(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createAppUserExtra(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createAppUserControl(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createAppUserSecret(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createAppUserThirdParty(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createAppRoleUser(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updateSubscriber(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetUser(ctx)
}
