package app

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	entapp "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/app"
)

func (h *Handler) ExistApp(ctx context.Context) (exist bool, err error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			App.
			Query().
			Where(
				entapp.ID(*h.ID),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
