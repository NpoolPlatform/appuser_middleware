package subscriber

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	subscribercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/subscriber"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"
)

func (h *Handler) UpdateSubscriber(ctx context.Context) (*npool.Subscriber, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := subscribercrud.UpdateSet(
			cli.Subscriber.UpdateOneID(*h.ID),
			&subscribercrud.Req{
				Registered: h.Registered,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetSubscriber(ctx)
}
