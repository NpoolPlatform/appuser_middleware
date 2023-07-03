package appsubscribe

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	appsubscribecrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/subscriber/app/subscribe"
)

func (h *Handler) ExistAppSubscribeConds(ctx context.Context) (bool, error) {
	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := appsubscribecrud.SetQueryConds(cli.AppSubscribe.Query(), h.Conds)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
