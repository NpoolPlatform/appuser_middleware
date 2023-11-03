package appsubscribe

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	appsubscribecrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/subscriber/app/subscribe"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber/app/subscribe"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

func (h *Handler) CreateAppSubscribe(ctx context.Context) (*npool.AppSubscribe, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	if *h.AppID == *h.SubscribeAppID {
		return nil, fmt.Errorf("cannot subscribe self")
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateAppSubscribe, *h.AppID, *h.SubscribeAppID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := appsubscribecrud.SetQueryConds(
			cli.AppSubscribe.Query(),
			&appsubscribecrud.Conds{
				AppID:          &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
				SubscribeAppID: &cruder.Cond{Op: cruder.EQ, Val: *h.SubscribeAppID},
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
				EntID:          h.EntID,
				AppID:          h.AppID,
				SubscribeAppID: h.SubscribeAppID,
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
