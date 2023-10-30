package auth

import (
	"context"
	"fmt"

	authcrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/authing/auth"
	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/auth"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func (h *Handler) CreateAuth(ctx context.Context) (*npool.Auth, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateAuth, *h.Resource, *h.Method)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	exist, err := h.ExistAuth(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("auth exist")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := authcrud.CreateSet(
			cli.Auth.Create(),
			&authcrud.Req{
				EntID:    h.EntID,
				AppID:    h.AppID,
				RoleID:   h.RoleID,
				UserID:   h.UserID,
				Resource: h.Resource,
				Method:   h.Method,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetAuth(ctx)
}

func (h *Handler) CreateAuths(ctx context.Context) ([]*npool.Auth, error) {
	ids := []uuid.UUID{}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			id := uuid.New()
			if req.EntID != nil {
				req.EntID = &id
			}
			if _, err := authcrud.CreateSet(
				tx.Auth.Create(),
				req,
			).Save(_ctx); err != nil {
				return err
			}
			ids = append(ids, id)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &authcrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))
	infos, _, err := h.GetAuths(ctx)
	if err != nil {
		return nil, err
	}
	return infos, nil
}
