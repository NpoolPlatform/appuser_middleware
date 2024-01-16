package recoverycode

import (
	"context"

	recoverycodecrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user/recoverycode"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/recoverycode"
)

func (h *Handler) UpdateRecoveryCode(ctx context.Context) (*npool.RecoveryCode, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := recoverycodecrud.UpdateSet(
			cli.RecoveryCode.UpdateOneID(*h.ID),
			&recoverycodecrud.Req{
				Used: h.Used,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetRecoveryCode(ctx)
}
