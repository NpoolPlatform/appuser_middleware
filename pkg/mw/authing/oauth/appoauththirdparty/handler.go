package appoauththirdparty

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/const"
	appoauththirdpartycrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/authing/oauth/appoauththirdparty"
	app "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/oauth/appoauththirdparty"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID            *uuid.UUID
	AppID         uuid.UUID
	ThirdPartyID  *uuid.UUID
	ClientID      *string
	ClientSecret  *string
	CallbackURL   *string
	Salt          *string
	ThirdPartyIDs []*uuid.UUID
	Reqs          []*npool.OAuthThirdPartyReq
	Conds         *appoauththirdpartycrud.Conds
	Offset        int32
	Limit         int32
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

func WithThirdPartyID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ThirdPartyID = &_id
		return nil
	}
}

func WithClientID(clientID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if clientID == nil {
			return nil
		}
		if *clientID == "" {
			return fmt.Errorf("invalid clientid")
		}
		h.ClientID = clientID
		return nil
	}
}

func WithClientSecret(clientSecret *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if clientSecret == nil {
			return nil
		}
		if *clientSecret == "" {
			return fmt.Errorf("invalid clientsecret")
		}
		h.ClientSecret = clientSecret
		return nil
	}
}

func WithSalt(salt *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if salt == nil {
			return nil
		}
		h.Salt = salt
		return nil
	}
}

func WithCallbackURL(callbackURL *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if callbackURL == nil {
			return nil
		}
		if *callbackURL == "" {
			return fmt.Errorf("invalid callbackurl")
		}
		h.CallbackURL = callbackURL
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &appoauththirdpartycrud.Conds{}
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
		if conds.ThirdPartyID != nil {
			id, err := uuid.Parse(conds.GetThirdPartyID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ThirdPartyID = &cruder.Cond{Op: conds.GetThirdPartyID().GetOp(), Val: id}
		}
		if conds.ClientName != nil {
			h.Conds.ClientName = &cruder.Cond{Op: conds.GetClientName().GetOp(), Val: basetypes.SignMethod(conds.GetClientName().GetValue())}
		}
		if len(conds.GetThirdPartyIDs().GetValue()) > 0 {
			_ids := []uuid.UUID{}
			for _, id := range conds.GetThirdPartyIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				_ids = append(_ids, _id)
			}
			h.Conds.ThirdPartyIDs = &cruder.Cond{Op: conds.GetThirdPartyIDs().GetOp(), Val: _ids}
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

func WithReqs(reqs []*npool.OAuthThirdPartyReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, req := range reqs {
			if _, err := uuid.Parse(*req.ThirdPartyID); err != nil {
				return err
			}
			if _, err := uuid.Parse(*req.AppID); err != nil {
				return err
			}
			if req.ID != nil {
				if _, err := uuid.Parse(*req.ID); err != nil {
					return err
				}
			}
		}
		h.Reqs = reqs
		return nil
	}
}
