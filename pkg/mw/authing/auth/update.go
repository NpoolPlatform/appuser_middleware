package auth

import (
	"context"

	authcrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/authing/auth"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/auth"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
)

func (h *Handler) UpdateAuth(ctx context.Context) (*npool.Auth, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := authcrud.UpdateSet(
			cli.Auth.UpdateOneID(*h.ID),
			&authcrud.Req{
				EntID:    h.EntID,
				AppID:    h.AppID,
				RoleID:   h.RoleID,
				UserID:   h.UserID,
				Resource: h.Resource,
				Method:   h.Method,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAuth(ctx)
}
