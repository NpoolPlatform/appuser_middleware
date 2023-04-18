package role

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	rolecrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"
)

func (h *Handler) UpdateRole(ctx context.Context) (*npool.Role, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := rolecrud.UpdateSet(
			cli.AppRole.UpdateOneID(*h.ID),
			&rolecrud.Req{
				ID:          h.ID,
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
