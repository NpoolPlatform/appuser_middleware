package user

import (
	"context"
	"fmt"

	usercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entappuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuser"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (h *Handler) ExistUser(ctx context.Context) (exist bool, err error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			AppUser.
			Query().
			Where(
				entappuser.AppID(h.AppID),
				entappuser.ID(*h.ID),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistUserConds(ctx context.Context) (exist bool, err error) {
	if h.Conds == nil {
		h.Conds = &usercrud.Conds{}
	}
	if h.ID != nil {
		h.Conds.ID = &cruder.Cond{Op: cruder.EQ, Val: *h.ID}
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := usercrud.SetQueryConds(cli.AppUser.Query(), h.Conds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
