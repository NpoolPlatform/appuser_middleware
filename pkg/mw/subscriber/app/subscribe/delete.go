package appsubscribe

import (
	"context"
	"time"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	appsubscribecrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/subscriber/app/subscribe"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber/app/subscribe"
)

func (h *Handler) DeleteAppSubscribe(ctx context.Context) (*npool.AppSubscribe, error) {
	info, err := h.GetAppSubscribe(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := appsubscribecrud.UpdateSet(
			cli.AppSubscribe.UpdateOneID(*h.ID),
			&appsubscribecrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
