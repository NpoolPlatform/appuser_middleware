package auth

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/auth"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
)

func (h *Handler) DeleteAuth(ctx context.Context) (*npool.Auth, error) {
	info, err := h.GetAuth(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := cli.
			Auth.
			UpdateOneID(*h.ID).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
