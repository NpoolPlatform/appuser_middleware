package app

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"

	appcrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/app"
	ctrlcrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/app/control"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createApp(ctx context.Context, tx *ent.Tx) error {
	if _, err := appcrud.CreateSet(
		tx.App.Create(),
		&appcrud.Req{
			ID:          h.ID,
			CreatedBy:   &h.CreatedBy,
			Name:        h.Name,
			Logo:        h.Logo,
			Description: h.Description,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createAppCtrl(ctx context.Context, tx *ent.Tx) error {
	if _, err := ctrlcrud.CreateSet(
		tx.AppControl.Create(),
		&ctrlcrud.Req{
			AppID:                    h.ID,
			SignupMethods:            h.SignupMethods,
			ExtSigninMethods:         h.ExtSigninMethods,
			RecaptchaMethod:          h.RecaptchaMethod,
			KycEnable:                h.KycEnable,
			SigninVerifyEnable:       h.SigninVerifyEnable,
			InvitationCodeMust:       h.InvitationCodeMust,
			CreateInvitationCodeWhen: h.CreateInvitationCodeWhen,
			MaxTypedCouponsPerOrder:  h.MaxTypedCouponsPerOrder,
			Maintaining:              h.Maintaining,
			CommitButtonTargets:      h.CommitButtonTargets,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateApp(ctx context.Context) (info *npool.App, err error) {
	handler := &createHandler{
		Handler: h,
	}

	id := uuid.New()
	if handler.ID == nil {
		handler.ID = &id
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createApp(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createAppCtrl(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetApp(ctx)
}

func CreateApp(ctx context.Context, in *npool.AppReq) (*npool.App, error) {
	return nil, nil
}
