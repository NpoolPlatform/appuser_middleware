package kyc

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	kyccrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/kyc"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

func (h *Handler) CreateKyc(ctx context.Context) (*npool.Kyc, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateUser, h.AppID, h.UserID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := kyccrud.SetQueryConds(
			cli.Kyc.Query(),
			&kyccrud.Conds{
				AppID:  &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
				UserID: &cruder.Cond{Op: cruder.EQ, Val: h.UserID},
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
			return fmt.Errorf("kyc exist")
		}

		if _, err := kyccrud.CreateSet(
			cli.Kyc.Create(),
			&kyccrud.Req{
				EntID:        h.EntID,
				AppID:        h.AppID,
				UserID:       h.UserID,
				DocumentType: h.DocumentType,
				IDNumber:     h.IDNumber,
				FrontImg:     h.FrontImg,
				BackImg:      h.BackImg,
				SelfieImg:    h.SelfieImg,
				EntityType:   h.EntityType,
				ReviewID:     h.ReviewID,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetKyc(ctx)
}
