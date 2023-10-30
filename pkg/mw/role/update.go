package role

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entapprole "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approle"

	rolecrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"
)

func (h *Handler) UpdateRole(ctx context.Context) (*npool.Role, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err := tx.
			AppRole.
			Query().
			Where(
				entapprole.ID(*h.ID),
				entapprole.DeletedAt(0),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}

		if _, err := rolecrud.UpdateSet(
			info.Update(),
			&rolecrud.Req{
				EntID:       h.EntID,
				CreatedBy:   h.CreatedBy,
				Role:        h.Role,
				Description: h.Description,
				Default:     h.Default,
				Genesis:     h.Genesis,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetRole(ctx)
}
