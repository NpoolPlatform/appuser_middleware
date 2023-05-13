package auth

import (
	"context"
	"fmt"

	authcrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/authing/auth"
	handler "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/handler"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/auth"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	*handler.Handler
	Conds *authcrud.Conds
	Reqs  []*npool.AuthReq
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
		h.Conds = &authcrud.Conds{}
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
		return nil
	}
}

func WithReqs(reqs []*npool.AuthReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, req := range reqs {
			if req.ID != nil {
				if _, err := uuid.Parse(*req.ID); err != nil {
					return err
				}
			}
			if _, err := uuid.Parse(*req.AppID); err != nil {
				return err
			}
			if _, err := uuid.Parse(*req.UserID); err != nil {
				return err
			}
			if _, err := uuid.Parse(*req.RoleID); err != nil {
				return err
			}
			const leastResourceLen = 3
			if len(*req.Resource) < leastResourceLen {
				return fmt.Errorf("resource %v invalid", *req.Resource)
			}
			switch *req.Method {
			case "POST":
			case "GET":
			default:
				return fmt.Errorf("method %v invalid", *req.Method)
			}
		}
		h.Reqs = reqs
		return nil
	}
}
