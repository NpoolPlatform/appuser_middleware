package role

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entapprole "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approle"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

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

		if h.Default != nil && *h.Default {
			stm, err := rolecrud.SetQueryConds(tx.AppRole.Query(), &rolecrud.Conds{
				Default: &cruder.Cond{Op: cruder.EQ, Val: *h.Default},
				AppID:   &cruder.Cond{Op: cruder.EQ, Val: &info.AppID},
				Role:    &cruder.Cond{Op: cruder.NEQ, Val: &info.Role},
			})
			if err != nil {
				return err
			}

			exist, err := stm.Exist(_ctx)
			if err != nil {
				return err
			}
			if exist {
				return fmt.Errorf("default role exist")
			}
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
