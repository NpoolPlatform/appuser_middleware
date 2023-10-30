package user

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	usercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role/user"
)

func (h *Handler) UpdateUser(ctx context.Context) (*npool.User, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := usercrud.UpdateSet(
			cli.AppRoleUser.UpdateOneID(*h.ID),
			&usercrud.Req{
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
