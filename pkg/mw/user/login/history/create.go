package history

import (
	"context"
	"fmt"

	historycrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user/login/history"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	user1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/login/history"

	"github.com/google/uuid"
)

func (h *Handler) CreateHistory(ctx context.Context) (*npool.History, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	userID := h.UserID.String()
	appID := h.AppID.String()

	handler, err := user1.NewHandler(
		ctx,
		user1.WithEntID(&userID, true),
		user1.WithAppID(&appID, true),
	)
	if err != nil {
		return nil, err
	}
	exist, err := handler.ExistUser(ctx)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("invalid user")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := historycrud.CreateSet(
			cli.LoginHistory.Create(),
			&historycrud.Req{
				EntID:     h.EntID,
				AppID:     h.AppID,
				UserID:    h.UserID,
				ClientIP:  h.ClientIP,
				UserAgent: h.UserAgent,
				Location:  h.Location,
				LoginType: h.LoginType,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}
		h.ID = &info.ID
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetHistory(ctx)
}
