package history

import (
	"context"

	historycrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/authing/history"
	handler "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/handler"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/history"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	*handler.Handler
	Conds   *historycrud.Conds
	Allowed *bool
}

func NewHandler(ctx context.Context, options ...interface{}) (*Handler, error) {
	_handler, err := handler.NewHandler(ctx, options...)
	if err != nil {
		return nil, err
	}

	h := &Handler{
		Handler: _handler,
	}
	for _, opt := range options {
		_opt, ok := opt.(func(context.Context, *Handler) error)
		if !ok {
			continue
		}
		if err := _opt(ctx, h); err != nil {
			return nil, err
		}
	}
	return h, nil
}

func WithAllowed(allowed *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Allowed = allowed
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &historycrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{
				Op:  conds.GetID().GetOp(),
				Val: id,
			}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{
				Op:  conds.GetAppID().GetOp(),
				Val: id,
			}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{
				Op:  conds.GetUserID().GetOp(),
				Val: id,
			}
		}
		if conds.Resource != nil {
			h.Conds.Resource = &cruder.Cond{
				Op:  conds.GetResource().GetOp(),
				Val: conds.GetResource().GetValue(),
			}
		}
		if conds.Method != nil {
			h.Conds.Method = &cruder.Cond{
				Op:  conds.GetMethod().GetOp(),
				Val: conds.GetMethod().GetValue(),
			}
		}
		if conds.Allowed != nil {
			h.Conds.Allowed = &cruder.Cond{
				Op:  conds.GetAllowed().GetOp(),
				Val: conds.GetAllowed().GetValue(),
			}
		}
		return nil
	}
}
