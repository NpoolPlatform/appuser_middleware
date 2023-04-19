package subscriber

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	subscribercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/subscriber"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"

	"github.com/google/uuid"
)

func (h *Handler) CreateSubscriber(ctx context.Context) (*npool.Subscriber, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := subscribercrud.CreateSet(
			cli.Subscriber.Create(),
			&subscribercrud.Req{
				ID:           h.ID,
				AppID:        &h.AppID,
				EmailAddress: h.EmailAddress,
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
