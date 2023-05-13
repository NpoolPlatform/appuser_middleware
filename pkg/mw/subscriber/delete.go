package subscriber

import (
	"context"
	"time"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	subscribercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/subscriber"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"
)

func (h *Handler) DeleteSubscriber(ctx context.Context) (*npool.Subscriber, error) {
	info, err := h.GetSubscriber(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := subscribercrud.UpdateSet(
			cli.Subscriber.UpdateOneID(*h.ID),
			&subscribercrud.Req{
				ID:        h.ID,
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
