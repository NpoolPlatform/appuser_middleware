package role

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/const"
	rolecrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role"
	app "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID          *uuid.UUID
	AppID       uuid.UUID
	CreatedBy   *uuid.UUID
	Role        *string
	Description *string
	Default     *bool
	Genesis     *bool
	Conds       *rolecrud.Conds
	Offset      int32
	Limit       int32
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
		h.AppID = _id
		return nil
	}
}

func WithCreatedBy(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		// TODO: check user exist
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CreatedBy = &_id
		return nil
	}
}

func WithRole(role *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if role == nil {
			return nil
		}
		if *role == "" {
			return fmt.Errorf("invalid role")
		}
		h.Role = role
		return nil
	}
}

func WithDescription(description *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Description = description
		return nil
	}
}

func WithDefault(defautl *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Default = defautl
		return nil
	}
}

func WithGenesis(genesis *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Genesis = genesis
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &rolecrud.Conds{}
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
		if conds.CreatedBy != nil {
			id, err := uuid.Parse(conds.GetCreatedBy().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CreatedBy = &cruder.Cond{Op: conds.GetCreatedBy().GetOp(), Val: id}
		}
		if conds.Role != nil {
			h.Conds.Role = &cruder.Cond{Op: conds.GetRole().GetOp(), Val: conds.GetRole().GetValue()}
		}
		if conds.Default != nil {
			h.Conds.Default = &cruder.Cond{Op: conds.GetDefault().GetOp(), Val: conds.GetDefault().GetValue()}
		}
		if conds.Roles != nil {
			h.Conds.Roles = &cruder.Cond{Op: conds.GetRoles().GetOp(), Val: conds.GetRoles().GetValue()}
		}
		if conds.Genesis != nil {
			h.Conds.Genesis = &cruder.Cond{Op: conds.GetGenesis().GetOp(), Val: conds.GetGenesis().GetValue()}
		}
		if len(conds.GetIDs().GetValue()) > 0 {
			_ids := []uuid.UUID{}
			for _, id := range conds.GetIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				_ids = append(_ids, _id)
			}
			h.Conds.IDs = &cruder.Cond{Op: conds.GetIDs().GetOp(), Val: _ids}
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
