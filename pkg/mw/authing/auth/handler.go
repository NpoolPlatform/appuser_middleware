package auth

import (
	"context"

	handler "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/handler"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/auth"
)

type Handler struct {
	*handler.Handler
	Conds *npool.Conds
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

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		// TODO: verify conds
		h.Conds = conds
		return nil
	}
}
