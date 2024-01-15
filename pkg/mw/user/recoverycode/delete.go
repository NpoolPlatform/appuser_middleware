package recoverycode

import (
	"context"
	"time"

	recoverycodecrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user/recoverycode"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/recoverycode"
)

func (h *Handler) DeleteKyc(ctx context.Context) (*npool.RecoveryCode, error) {
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
