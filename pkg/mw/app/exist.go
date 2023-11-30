package app

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	appcrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/app"
	entapp "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/app"
)

func (h *Handler) ExistApp(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			App.
			Query().
			Where(
				entapp.EntID(*h.EntID),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistAppConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := appcrud.SetQueryConds(cli.App.Query(), h.Conds)
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
