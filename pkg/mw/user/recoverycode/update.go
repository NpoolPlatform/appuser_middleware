package recoverycode

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/recoverycode"
)

func (h *Handler) UpdateRecoveryCode(ctx context.Context) (*npool.RecoveryCode, error) {
	if h.Used == nil {
		return h.GetRecoveryCode(ctx)
	}
	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		if _, err := cli.
			RecoveryCode.UpdateOneID(*h.ID).
			SetUsed(*h.Used).
			Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetRecoveryCode(ctx)
}
