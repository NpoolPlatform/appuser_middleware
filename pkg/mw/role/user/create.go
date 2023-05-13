package user

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	usercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role/user"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role/user"

	"github.com/google/uuid"
)

func (h *Handler) CreateUser(ctx context.Context) (*npool.User, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}
	if h.RoleID == nil || h.UserID == nil {
		return nil, fmt.Errorf("invalid roleid or userid")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := usercrud.SetQueryConds(
			cli.AppRoleUser.Query(),
			&usercrud.Conds{
				AppID:  &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
				RoleID: &cruder.Cond{Op: cruder.EQ, Val: *h.RoleID},
				UserID: &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
			},
		)
		if err != nil {
			return err
		}

		info, err := stm.Only(_ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}
		if info != nil {
			h.ID = &info.ID
			return nil
		}

		if _, err := usercrud.CreateSet(
			cli.AppRoleUser.Create(),
			&usercrud.Req{
				ID:     h.ID,
				AppID:  &h.AppID,
				RoleID: h.RoleID,
				UserID: h.UserID,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetUser(ctx)
}
