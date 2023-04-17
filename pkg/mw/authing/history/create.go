package history

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	historycrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/authing/history"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/history"

	"github.com/google/uuid"
)

func (h *Handler) CreateHistory(ctx context.Context) (*npool.History, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := historycrud.CreateSet(
			cli.AuthHistory.Create(),
			&historycrud.Req{
				ID:       h.ID,
				AppID:    &h.AppID,
				UserID:   h.UserID,
				Resource: &h.Resource,
				Method:   &h.Method,
				Allowed:  &h.Allowed,
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
