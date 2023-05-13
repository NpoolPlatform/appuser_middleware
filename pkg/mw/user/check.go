package user

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	usercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

func (h *Handler) checkAccountExist(ctx context.Context) error {
	if h.PhoneNO == nil && h.EmailAddress == nil {
		return fmt.Errorf("invalid account")
	}

	conds := &usercrud.Conds{
		AppID: &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
	}
	if h.EmailAddress != nil {
		conds.EmailAddress = &cruder.Cond{Op: cruder.EQ, Val: *h.EmailAddress}
	}
	if h.PhoneNO != nil {
		conds.PhoneNO = &cruder.Cond{Op: cruder.EQ, Val: *h.PhoneNO}
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := usercrud.SetQueryConds(cli.AppUser.Query(), conds)
		if err != nil {
			return err
		}
		exist, err := stm.Exist(ctx)
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("user already exist")
		}
		return nil
	})
}
