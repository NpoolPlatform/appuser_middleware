package history

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/const"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/history"

	"github.com/google/uuid"
)

type Handler struct {
	ID       *uuid.UUID
	AppID    uuid.UUID
	UserID   *uuid.UUID
	Method   string
	Resource string
	Conds    *npool.Conds
	Offset   int32
	Limit    int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithAppID(id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.AppID = _id
		return nil
	}
}

func WithUserID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.UserID = &_id
		return nil
	}
}

func WithMethod(method string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		switch method {
		case "POST":
		case "GET":
		default:
			return fmt.Errorf("method %v invalid", method)
		}
		h.Method = method
		return nil
	}
}

func WithResource(resource string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		const leastResourceLen = 3
		if len(resource) < leastResourceLen {
			return fmt.Errorf("resource %v invalid", resource)
		}
		h.Resource = resource
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = conds
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
