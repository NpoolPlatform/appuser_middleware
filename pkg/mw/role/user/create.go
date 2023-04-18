package user

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	usercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role/user"

	"github.com/google/uuid"
)

func (h *Handler) CreateUser(ctx context.Context) (*npool.User, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
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
