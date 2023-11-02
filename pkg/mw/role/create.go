package role

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	rolecrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

func (h *Handler) CreateRole(ctx context.Context) (*npool.Role, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateRole, *h.AppID, *h.Role)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := rolecrud.SetQueryConds(cli.AppRole.Query(), &rolecrud.Conds{
			AppID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			Role:  &cruder.Cond{Op: cruder.EQ, Val: *h.Role},
		})
		if err != nil {
			return err
		}

		exist, err := stm.Exist(_ctx)
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("role exist")
		}

		if _, err := rolecrud.CreateSet(
			cli.AppRole.Create(),
			&rolecrud.Req{
				EntID:       h.EntID,
				AppID:       h.AppID,
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
			if req.EntID == nil {
				req.EntID = &id
			}

			key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateRole, *req.AppID, *req.Role)
			if err := redis2.TryLock(key, 0); err != nil {
				return err
			}

			stm, err := rolecrud.SetQueryConds(cli.AppRole.Query(), &rolecrud.Conds{
				AppID: &cruder.Cond{Op: cruder.EQ, Val: *req.AppID},
				Role:  &cruder.Cond{Op: cruder.EQ, Val: *req.Role},
			})
			if err != nil {
				_ = redis2.Unlock(key)
				return err
			}

			exist, err := stm.Exist(_ctx)
			if err != nil {
				_ = redis2.Unlock(key)
				return err
			}
			if exist {
				_ = redis2.Unlock(key)
				return fmt.Errorf("role exist")
			}

			if _, err := rolecrud.CreateSet(
				cli.AppRole.Create(),
				req,
			).Save(ctx); err != nil {
				_ = redis2.Unlock(key)
				return err
			}
			_ = redis2.Unlock(key)
			ids = append(ids, id)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &rolecrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	infos, _, err := h.GetRoles(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
