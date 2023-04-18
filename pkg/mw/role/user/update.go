package user

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	usercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role/user"
)

func (h *Handler) UpdateUser(ctx context.Context) (*npool.User, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := usercrud.UpdateSet(
			cli.AppRoleUser.UpdateOneID(*h.ID),
			&usercrud.Req{
				ID:     h.ID,
				RoleID: h.RoleID,
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
