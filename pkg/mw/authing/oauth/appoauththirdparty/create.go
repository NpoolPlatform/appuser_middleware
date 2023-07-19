package appoauththirdparty

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	appoauththirdpartycrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/authing/oauth/appoauththirdparty"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/oauth/appoauththirdparty"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	"github.com/google/uuid"
)

func (h *Handler) CreateOAuthThirdParty(ctx context.Context) (*npool.OAuthThirdParty, error) {
	if h.ThirdPartyID == nil {
		return nil, fmt.Errorf("invalid clientsecret")
	}

	key := fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixCreateUserTransfer, *h.ThirdPartyID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	handler, err := NewHandler(
		ctx,
		WithConds(&npool.Conds{
			AppID:        &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			ThirdPartyID: &basetypes.StringVal{Op: cruder.EQ, Value: h.ThirdPartyID.String()},
		}),
	)
	if err != nil {
		return nil, err
	}
	exist, err := handler.ExistOAuthThirdPartyConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("oauththirdparty exist")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := appoauththirdpartycrud.CreateSet(
			tx.AppOAuthThirdParty.Create(),
			&appoauththirdpartycrud.Req{
				ID:           h.ID,
				AppID:        &h.AppID,
				ThirdPartyID: h.ThirdPartyID,
			},
		).Save(ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOAuthThirdParty(ctx)
}
