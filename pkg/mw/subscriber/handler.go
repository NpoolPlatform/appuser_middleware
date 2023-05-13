package subscriber

import (
	"context"
	"fmt"
	"net/mail"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/const"
	subscribercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/subscriber"
	app "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID           *uuid.UUID
	AppID        uuid.UUID
	EmailAddress *string
	Registered   *bool
	Conds        *subscribercrud.Conds
	Offset       int32
	Limit        int32
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
		handler, err := app.NewHandler(
			ctx,
			app.WithID(&id),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistApp(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid app")
		}
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.AppID = _id
		return nil
	}
}

func WithEmailAddress(addr *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if addr == nil {
			return nil
		}
		if _, err := mail.ParseAddress(*addr); err != nil {
			return err
		}
		h.EmailAddress = addr
		return nil
	}
}

func WithRegistered(registered *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Registered = registered
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &subscribercrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: id}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.EmailAddress != nil {
			h.Conds.EmailAddress = &cruder.Cond{
				Op:  conds.GetEmailAddress().GetOp(),
				Val: conds.GetEmailAddress().GetValue(),
			}
		}
		if conds.Registered != nil {
			h.Conds.Registered = &cruder.Cond{
				Op:  conds.GetRegistered().GetOp(),
				Val: conds.GetRegistered().GetValue(),
			}
		}
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
