package history

import (
	"context"

	handler "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/handler"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/history"
)

type Handler struct {
	*handler.Handler
	Conds   *npool.Conds
	Allowed bool
}

func NewHandler(ctx context.Context, options ...interface{}) (*Handler, error) {
	_handler, err := handler.NewHandler(ctx, options...)
	if err != nil {
		return nil, err
	}

	handler := &Handler{
		Handler: _handler,
	}
	for _, opt := range options {
		_opt, ok := opt.(func(context.Context, *Handler) error)
		if !ok {
			continue
		}
		if err := _opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithAllowed(allowed bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Allowed = allowed
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		// TODO: verify conds
		h.Conds = conds
		return nil
	}
}
