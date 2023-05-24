//nolint:dupl
package user

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-middleware/pkg/encrypt"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	usercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user"
	userthirdpartycrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user/3rdparty"
	banappusercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user/ban"
	userctrlcrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user/control"
	userextracrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user/extra"
	usersecretcrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user/secret"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateAppUser(ctx context.Context, tx *ent.Tx) error {
	if _, err := usercrud.UpdateSet(
		tx.AppUser.UpdateOneID(*h.ID),
		&usercrud.Req{
			PhoneNO:       h.PhoneNO,
			EmailAddress:  h.EmailAddress,
			ImportFromApp: h.ImportFromAppID,
		}).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateAppUserExtra(ctx context.Context, tx *ent.Tx) error {
	stm, err := userextracrud.SetQueryConds(
		tx.AppUserExtra.Query(),
		&userextracrud.Conds{
			AppID:  &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
			UserID: &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.ForUpdate().Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	req := &userextracrud.Req{
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
		if _, err = userextracrud.CreateSet(
			tx.AppUserExtra.Create(),
			req,
		).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	_stm, err := userextracrud.UpdateSet(
		ctx,
		info.Update(),
		req,
	)
	if err != nil {
		return err
	}
	if _, err := _stm.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateAppUserControl(ctx context.Context, tx *ent.Tx) error {
	stm, err := userctrlcrud.SetQueryConds(
		tx.AppUserControl.Query(),
		&userctrlcrud.Conds{
			AppID:  &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
			UserID: &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.ForUpdate().Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	req := &userctrlcrud.Req{
		AppID:              &h.AppID,
		UserID:             h.ID,
		GoogleAuthVerified: h.GoogleAuthVerified,
		SigninVerifyType:   h.SigninVerifyType,
		Kol:                h.Kol,
		KolConfirmed:       h.KolConfirmed,
		SelectedLangID:     h.SelectedLangID,
	}

	if info == nil {
		if _, err := userctrlcrud.CreateSet(
			tx.AppUserControl.Create(),
			req,
		).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	if _, err = userctrlcrud.UpdateSet(
		info.Update(),
		req,
	).Save(ctx); err != nil {
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

	stm, err := usersecretcrud.SetQueryConds(
		tx.AppUserSecret.Query(),
		&usersecretcrud.Conds{
			AppID:  &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
			UserID: &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.ForUpdate().Only(ctx)
	if err != nil {
		return err
	}

	if _, err = usersecretcrud.UpdateSet(
		info.Update(),
		&usersecretcrud.Req{
			PasswordHash: password,
			Salt:         salt,
			GoogleSecret: h.GoogleSecret,
		}).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateAppUserThirdParty(ctx context.Context, tx *ent.Tx) error {
	stm, err := userthirdpartycrud.SetQueryConds(
		tx.AppUserThirdParty.Query(),
		&userthirdpartycrud.Conds{
			AppID:  &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
			UserID: &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.ForUpdate().Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	if info == nil && h.ThirdPartyID == nil {
		return nil
	}

	req := &userthirdpartycrud.Req{
		AppID:              &h.AppID,
		UserID:             h.ID,
		ThirdPartyID:       h.ThirdPartyID,
		ThirdPartyUserID:   h.ThirdPartyUserID,
		ThirdPartyUsername: h.ThirdPartyUsername,
		ThirdPartyAvatar:   h.ThirdPartyAvatar,
	}

	if info == nil {
		if _, err := userthirdpartycrud.CreateSet(
			tx.AppUserThirdParty.Create(),
			req,
		).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	if _, err = userthirdpartycrud.UpdateSet(
		info.Update(),
		req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateBanAppUser(ctx context.Context, tx *ent.Tx) error {
	if h.Banned == nil {
		return nil
	}
	if h.ID == nil {
		return fmt.Errorf("invalid id")
	}

	stm, err := banappusercrud.SetQueryConds(
		tx.BanAppUser.Query(),
		&banappusercrud.Conds{
			AppID:  &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
			UserID: &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		})
	if err != nil {
		return err
	}

	info, err := stm.Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	fmt.Printf("banned=%v,info=%v,err=%v,app_id=%v,user_id=%v\n", *h.Banned, info, err, h.AppID, h.ID)

	if info != nil {
		now := uint32(0)
		if !*h.Banned {
			now = uint32(time.Now().Unix())
		}
		if _, err := banappusercrud.UpdateSet(
			tx.BanAppUser.UpdateOneID(info.ID),
			&banappusercrud.Req{
				AppID:     &h.AppID,
				UserID:    h.ID,
				DeletedAt: &now,
				Message:   h.BanMessage,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	if !*h.Banned {
		return nil
	}

	if _, err := banappusercrud.CreateSet(
		tx.BanAppUser.Create(),
		&banappusercrud.Req{
			AppID:   &h.AppID,
			UserID:  h.ID,
			Message: h.BanMessage,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) UpdateUser(ctx context.Context) (*npool.User, error) {
	if err := h.checkAccountExist(ctx); err != nil {
		return nil, err
	}

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
		if err := handler.updateBanAppUser(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetUser(ctx)
}
