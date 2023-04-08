package user

import (
	"context"
	"time"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	entapproleuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approleuser"
	entappusercontrol "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusercontrol"
	entappuserextra "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"
	entappusersecret "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"
	entappuserthirdparty "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserthirdparty"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteAppUser(ctx context.Context, tx *ent.Tx) error {
	if _, err := tx.
		AppUser.
		UpdateOneID(uuid.MustParse(*h.ID)).
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx); err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	return nil
}

func (h *deleteHandler) deleteAppUserExtra(ctx context.Context, tx *ent.Tx) error {
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

	if _, err := info.
		Update().
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteAppUserControl(ctx context.Context, tx *ent.Tx) error {
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

	if _, err := info.
		Update().
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteAppUserSecret(ctx context.Context, tx *ent.Tx) error {
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
		if !ent.IsNotFound(err) {
			return err
		}
	}

	if _, err := info.
		Update().
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteAppUserThirdParty(ctx context.Context, tx *ent.Tx) error {
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

	if _, err := info.
		Update().
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteAppRoleUser(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppRoleUser.
		Query().
		Where(
			entapproleuser.AppID(uuid.MustParse(h.AppID)),
			entapproleuser.UserID(uuid.MustParse(*h.ID)),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	if _, err := info.
		Update().
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteUser(ctx context.Context) (*npool.User, error) {
	info, err := h.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteAppUser(ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppUserExtra(ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppUserControl(ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppUserSecret(ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppUserThirdParty(ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppRoleUser(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
