package kyc

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	kyccrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/kyc"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"
)

func (h *Handler) UpdateKyc(ctx context.Context) (*npool.Kyc, error) {
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := kyccrud.UpdateSet(
			cli.Kyc.UpdateOneID(*h.ID),
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
				State:        h.State,
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
