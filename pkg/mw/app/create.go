package app

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	appcrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/app"
	ctrlcrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/app/control"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"

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

func (h *Handler) CreateApps(ctx context.Context) (infos []*npool.App, err error) {
	ids := []uuid.UUID{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			id := uuid.New()
			if req.ID != nil {
				id = uuid.MustParse(*req.ID)
			}

			createdBy := uuid.MustParse(*req.CreatedBy)

			if _, err := appcrud.CreateSet(
				tx.App.Create(),
				&appcrud.Req{
					ID:          &id,
					CreatedBy:   &createdBy,
					Name:        req.Name,
					Logo:        req.Logo,
					Description: req.Description,
				},
			).Save(ctx); err != nil {
				return err
			}
			if _, err := ctrlcrud.CreateSet(
				tx.AppControl.Create(),
				&ctrlcrud.Req{
					AppID:                    &id,
					SignupMethods:            req.SignupMethods,
					ExtSigninMethods:         req.ExtSigninMethods,
					RecaptchaMethod:          req.RecaptchaMethod,
					KycEnable:                req.KycEnable,
					SigninVerifyEnable:       req.SigninVerifyEnable,
					InvitationCodeMust:       req.InvitationCodeMust,
					CreateInvitationCodeWhen: req.CreateInvitationCodeWhen,
					MaxTypedCouponsPerOrder:  req.MaxTypedCouponsPerOrder,
					Maintaining:              req.Maintaining,
					CommitButtonTargets:      req.CommitButtonTargets,
				},
			).Save(ctx); err != nil {
				return err
			}
			ids = append(ids, id)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &appcrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	infos, _, err = h.GetApps(ctx)
	if err != nil {
		return nil, err
	}
	return infos, nil
}
