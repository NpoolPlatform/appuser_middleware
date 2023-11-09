//nolint:dupl
package user

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	entapproleuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approleuser"
	entappusercontrol "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appusercontrol"
	entappuserextra "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuserextra"
	entappusersecret "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appusersecret"
	entappuserthirdparty "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuserthirdparty"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteAppUser(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppUser.
		UpdateOneID(*h.ID).
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return err
	}
	h.AppID = &info.AppID
	h.EntID = &info.EntID
	return nil
}

func (h *deleteHandler) deleteAppUserExtra(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppUserExtra.
		Query().
		Where(
			entappuserextra.AppID(*h.AppID),
			entappuserextra.UserID(*h.EntID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
		return nil
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
			entappusercontrol.AppID(*h.AppID),
			entappusercontrol.UserID(*h.EntID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
		return nil
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
			entappusersecret.AppID(*h.AppID),
			entappusersecret.UserID(*h.EntID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
		return nil
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
			entappuserthirdparty.AppID(*h.AppID),
			entappuserthirdparty.UserID(*h.EntID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
		return nil
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
			entapproleuser.AppID(*h.AppID),
			entapproleuser.UserID(*h.EntID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
		return nil
	}

	if _, err := info.
		Update().
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteUser(ctx context.Context) (info *npool.User, err error) {
	info, err = h.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteAppUser(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppUserExtra(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppUserControl(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppUserSecret(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppUserThirdParty(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppRoleUser(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (h *Handler) DeleteThirdUser(ctx context.Context) (info *npool.User, err error) {
	info, err = h.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	if info.ThirdPartyID == nil && info.ThirdPartyUserID == nil {
		return nil, fmt.Errorf("invalid thirdparty")
	}

	id1, err := uuid.Parse(info.EntID)
	if err != nil {
		return nil, err
	}
	h.EntID = &id1

	id2, err := uuid.Parse(info.AppID)
	if err != nil {
		return nil, err
	}
	h.AppID = &id2

	h.ThirdPartyUserID = info.ThirdPartyUserID

	id3, err := uuid.Parse(*info.ThirdPartyID)
	if err != nil {
		return nil, err
	}
	h.ThirdPartyID = &id3

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.
			AppUserThirdParty.
			Query().
			Where(
				entappuserthirdparty.AppID(*h.AppID),
				entappuserthirdparty.UserID(*h.EntID),
				entappuserthirdparty.ThirdPartyUserID(*h.ThirdPartyUserID),
				entappuserthirdparty.ThirdPartyID(*h.ThirdPartyID),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
			return nil
		}

		if _, err := info.
			Update().
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
