package kyc

import (
	"context"
	"time"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	kyccrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/kyc"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"
)

func (h *Handler) DeleteKyc(ctx context.Context) (*npool.Kyc, error) {
	info, err := h.GetKyc(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	h.ID = &info.ID

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := kyccrud.UpdateSet(
			cli.Kyc.UpdateOneID(*h.ID),
			&kyccrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
