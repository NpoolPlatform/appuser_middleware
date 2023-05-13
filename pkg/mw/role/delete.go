package role

import (
	"context"
	"time"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	rolecrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"
)

func (h *Handler) DeleteRole(ctx context.Context) (*npool.Role, error) {
	info, err := h.GetRole(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := rolecrud.UpdateSet(
			cli.AppRole.UpdateOneID(*h.ID),
			&rolecrud.Req{
				ID:        h.ID,
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
