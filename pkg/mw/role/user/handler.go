package user

import (
	"context"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/const"
	usercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role/user"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID     *uuid.UUID
	AppID  uuid.UUID
	RoleID *uuid.UUID
	UserID *uuid.UUID
	Conds  *usercrud.Conds
	Offset int32
	Limit  int32
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

func WithRoleID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.RoleID = &_id
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

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error { //nolint
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &usercrud.Conds{}
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
		if conds.RoleID != nil {
			id, err := uuid.Parse(conds.GetRoleID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.RoleID = &cruder.Cond{Op: conds.GetRoleID().GetOp(), Val: id}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{Op: conds.GetUserID().GetOp(), Val: id}
		}
		if len(conds.GetAppIDs().GetValue()) > 0 {
			ids := []uuid.UUID{}
			for _, id := range conds.GetAppIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.AppIDs = &cruder.Cond{Op: conds.GetAppIDs().GetOp(), Val: ids}
		}
		if len(conds.GetRoleIDs().GetValue()) > 0 {
			ids := []uuid.UUID{}
			for _, id := range conds.GetRoleIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.RoleIDs = &cruder.Cond{Op: conds.GetRoleIDs().GetOp(), Val: ids}
		}
		if conds.Genesis != nil {
			h.Conds.Genesis = &cruder.Cond{
				Op:  conds.GetGenesis().GetOp(),
				Val: conds.GetGenesis().GetValue(),
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
