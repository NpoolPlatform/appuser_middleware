package role

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	rolecrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"

	"github.com/google/uuid"
)

func (h *Handler) CreateRole(ctx context.Context) (*npool.Role, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := rolecrud.CreateSet(
			cli.AppRole.Create(),
			&rolecrud.Req{
				ID:          h.ID,
				AppID:       &h.AppID,
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

func (h *Handler) CreateRoles(ctx context.Context) ([]*npool.Role, error) {
	ids := []uuid.UUID{}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		for _, req := range h.Reqs {
			id := uuid.New()
			if req.ID != nil {
				id = uuid.MustParse(*req.ID)
			}
			appID := uuid.MustParse(*req.AppID)
			createdBy := uuid.MustParse(*req.CreatedBy)
			if _, err := rolecrud.CreateSet(
				cli.AppRole.Create(),
				&rolecrud.Req{
					ID:          &id,
					AppID:       &appID,
					CreatedBy:   &createdBy,
					Role:        req.Role,
					Description: req.Description,
					Default:     req.Default,
					Genesis:     req.Genesis,
				},
			).Save(ctx); err != nil {
				return err
			}
			ids = append(ids, id)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &rolecrud.Conds{
		IDs: &cruder.Cond{Op: cruder.EQ, Val: ids},
	}
	infos, _, err := h.GetRoles(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
