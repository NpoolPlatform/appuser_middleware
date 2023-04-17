package app

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	appcrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/app"
	ctrlcrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/app/control"
	entappctrl "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appcontrol"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateApp(ctx context.Context, tx *ent.Tx) error {
	if h.ID == nil {
		return fmt.Errorf("invalid id")
	}

	if _, err := appcrud.UpdateSet(
		tx.App.UpdateOneID(*h.ID),
		&appcrud.Req{
			Name:        h.Name,
			Logo:        h.Logo,
			Description: h.Description,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateAppCtrl(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppControl.
		Query().
		Where(
			entappctrl.AppID(*h.ID),
			entappctrl.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}

	req := &ctrlcrud.Req{
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
	}

	if info == nil {
		if _, err = ctrlcrud.CreateSet(
			tx.AppControl.Create(),
			req,
		).Save(ctx); err != nil {
			return err
		}
		return nil
	}

	if _, err = ctrlcrud.UpdateSet(info.Update(), req).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateApp(ctx context.Context) (*npool.App, error) {
	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateApp(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateAppCtrl(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetApp(ctx)
}
