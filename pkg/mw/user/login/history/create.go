package history

import (
	"context"

	historycrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user/login/history"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/login/history"

	"github.com/google/uuid"
)

func (h *Handler) CreateHistory(ctx context.Context) (*npool.History, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := historycrud.CreateSet(
			cli.LoginHistory.Create(),
			&historycrud.Req{
				ID:        h.ID,
				AppID:     &h.AppID,
				UserID:    &h.UserID,
				ClientIP:  h.ClientIP,
				UserAgent: h.UserAgent,
				Location:  h.Location,
				LoginType: h.LoginType,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetHistory(ctx)
}
