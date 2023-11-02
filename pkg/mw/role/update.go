package role

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entapprole "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approle"

	rolecrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"
)

func (h *Handler) UpdateRole(ctx context.Context) (*npool.Role, error) {
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

		info, err = rolecrud.UpdateSet(
			info.Update(),
			&rolecrud.Req{
				EntID:       h.EntID,
				CreatedBy:   h.CreatedBy,
				Role:        h.Role,
				Description: h.Description,
				Default:     h.Default,
				Genesis:     h.Genesis,
			},
		).Save(ctx)
		if err != nil {
			return err
		}
		h.AppID = &info.AppID
		h.EntID = &info.EntID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetRole(ctx)
}
