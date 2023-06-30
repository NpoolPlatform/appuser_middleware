package appsubscribe

import (
	"context"

	appsubscribecrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/subscriber/app/subscribe"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber/app/subscribe"

	"github.com/google/uuid"
)

func (h *Handler) CreateAppSubscribe(ctx context.Context) (*npool.AppSubscribe, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := appsubscribecrud.SetQueryConds(
			cli.AppSubscribe.Query(),
			&appsubscribecrud.Conds{
				AppID:        &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
				EmailAddress: &cruder.Cond{Op: cruder.EQ, Val: *h.EmailAddress},
			},
		)
		if err != nil {
			return err
		}

		info, err := stm.Only(_ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}
		if info != nil {
			h.ID = &info.ID
			return nil
		}

		if _, err := appsubscribecrud.CreateSet(
			cli.AppSubscribe.Create(),
			&appsubscribecrud.Req{
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

	return h.GetAppSubscribe(ctx)
}
