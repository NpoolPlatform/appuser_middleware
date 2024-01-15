package recoverycode

import (
	"context"
	"time"

	recoverycodecrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user/recoverycode"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entrecoverycode "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/recoverycode"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/recoverycode"
)

func (h *Handler) DeleteRecoveryCode(ctx context.Context) (*npool.RecoveryCode, error) {
	info, err := h.GetRecoveryCode(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	h.ID = &info.ID

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := recoverycodecrud.UpdateSet(
			cli.RecoveryCode.UpdateOneID(*h.ID),
			&recoverycodecrud.Req{
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (h *Handler) DeleteRecoveryCodes(ctx context.Context) ([]*npool.RecoveryCode, error) {
	h.Conds = &recoverycodecrud.Conds{
		AppID:  &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		UserID: &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
	}
	infos, _, err := h.GetRecoveryCodes(ctx)
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := cli.
			RecoveryCode.
			Update().
			Where(
				entrecoverycode.AppID(*h.AppID),
				entrecoverycode.UserID(*h.UserID),
			).
			SetDeletedAt(now).
			Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return infos, nil
}
