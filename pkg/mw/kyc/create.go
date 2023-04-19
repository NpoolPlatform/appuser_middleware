package kyc

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	kyccrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/kyc"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"

	"github.com/google/uuid"
)

func (h *Handler) CreateKyc(ctx context.Context) (*npool.Kyc, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := kyccrud.CreateSet(
			cli.Kyc.Create(),
			&kyccrud.Req{
				ID:           h.ID,
				AppID:        &h.AppID,
				UserID:       &h.UserID,
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
