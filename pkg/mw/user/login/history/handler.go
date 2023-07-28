//nolint:dupl
package history

import (
	"context"
	"fmt"
	"net"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/const"
	historycrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user/login/history"
	app "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	user "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/login/history"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type Handler struct {
	ID        *uuid.UUID
	AppID     uuid.UUID
	UserID    uuid.UUID
	ClientIP  *string
	UserAgent *string
	Location  *string
	LoginType *basetypes.LoginType
	Conds     *historycrud.Conds
	Offset    int32
	Limit     int32
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

func WithUserID(id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := user.NewHandler(
			ctx,
			user.WithID(&id),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistUser(ctx)
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
		h.UserID = _id
		return nil
	}
}

func WithClientIP(ip *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if ip == nil {
			return nil
		}
		if ip := net.ParseIP(*ip); ip == nil {
			return fmt.Errorf("invalid client ip")
		}
		h.ClientIP = ip
		return nil
	}
}

func WithUserAgent(agent *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.UserAgent = agent
		return nil
	}
}

func WithLocation(location *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Location = location
		return nil
	}
}

func WithLoginType(loginType *basetypes.LoginType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if loginType == nil {
			return nil
		}
		switch *loginType {
		case basetypes.LoginType_FreshLogin:
		case basetypes.LoginType_RefreshLogin:
		default:
			return fmt.Errorf("invalid login type")
		}
		h.LoginType = loginType
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
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: id}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{Op: conds.GetUserID().GetOp(), Val: id}
		}
		if conds.LoginType != nil {
			h.Conds.LoginType = &cruder.Cond{
				Op:  conds.GetLoginType().GetOp(),
				Val: basetypes.LoginType(conds.GetLoginType().GetValue()),
			}
		}
		if conds.ClientIP != nil {
			h.Conds.ClientIP = &cruder.Cond{
				Op:  conds.GetClientIP().GetOp(),
				Val: conds.GetClientIP().GetValue(),
			}
		}
		if conds.Location != nil {
			h.Conds.Location = &cruder.Cond{
				Op:  conds.GetLocation().GetOp(),
				Val: conds.GetLocation().GetValue(),
			}
		}
		if conds.UserAgent != nil {
			h.Conds.UserAgent = &cruder.Cond{
				Op:  conds.GetUserAgent().GetOp(),
				Val: conds.GetUserAgent().GetValue(),
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
