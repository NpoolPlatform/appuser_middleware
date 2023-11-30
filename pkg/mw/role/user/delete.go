package user

import (
	"context"
	"time"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	usercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role/user"
)

func (h *Handler) DeleteUser(ctx context.Context) (*npool.User, error) {
	info, err := h.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := usercrud.UpdateSet(
			cli.AppRoleUser.UpdateOneID(*h.ID),
			&usercrud.Req{
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
