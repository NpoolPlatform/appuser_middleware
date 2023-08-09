package oauththirdparty

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/const"
	thidpartycrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/oauth/oauththirdparty"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/oauth/oauththirdparty"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID             *uuid.UUID
	ClientName     *basetypes.SignMethod
	ClientTag      *string
	ClientLogoURL  *string
	ClientOAuthURL *string
	ResponseType   *string
	Scope          *string
	Conds          *thidpartycrud.Conds
	Offset         int32
	Limit          int32
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

func WithClientName(clientName *basetypes.SignMethod) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if clientName == nil {
			return nil
		}
		switch *clientName {
		case basetypes.SignMethod_Twitter:
		case basetypes.SignMethod_Github:
		case basetypes.SignMethod_Facebook:
		case basetypes.SignMethod_Linkedin:
		case basetypes.SignMethod_Wechat:
		case basetypes.SignMethod_Google:
		default:
			return fmt.Errorf("invalid clientname")
		}
		h.ClientName = clientName
		return nil
	}
}

func WithClientTag(clientTag *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if clientTag == nil {
			return nil
		}
		if *clientTag == "" {
			return fmt.Errorf("invalid clienttag")
		}
		h.ClientTag = clientTag
		return nil
	}
}

func WithClientLogoURL(clientLogoURL *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if clientLogoURL == nil {
			return nil
		}
		if *clientLogoURL == "" {
			return fmt.Errorf("invalid clientlogourl")
		}
		h.ClientLogoURL = clientLogoURL
		return nil
	}
}

func WithClientOAuthURL(clientOAuthURL *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if clientOAuthURL == nil {
			return nil
		}
		if *clientOAuthURL == "" {
			return fmt.Errorf("invalid clientoauthurl")
		}
		h.ClientOAuthURL = clientOAuthURL
		return nil
	}
}

func WithResponseType(responseType *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if responseType == nil {
			return nil
		}
		if *responseType == "" {
			return fmt.Errorf("invalid responsetype")
		}
		h.ResponseType = responseType
		return nil
	}
}

func WithScope(scope *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if scope == nil {
			return nil
		}
		h.Scope = scope
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &thidpartycrud.Conds{}
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

		if conds.ClientName != nil {
			h.Conds.ClientName = &cruder.Cond{Op: conds.GetClientName().GetOp(), Val: basetypes.SignMethod(conds.GetClientName().GetValue())}
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
