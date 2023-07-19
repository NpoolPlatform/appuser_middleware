package oauththirdparty

import (
	"context"
	"fmt"

	oauththirdpartycrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/authing/oauth/oauththirdparty"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entoauththirdparty "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/oauththirdparty"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/oauth/oauththirdparty"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func (h *Handler) UpdateOAuthThirdParty(ctx context.Context) (*npool.OAuthThirdParty, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	info, err := h.GetOAuthThirdParty(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	if h.ClientName != nil {
		handler, err := NewHandler(
			ctx,
			WithConds(&npool.Conds{
				ClientName: &basetypes.StringVal{Op: cruder.EQ, Value: *h.ClientName},
			}),
		)
		if err != nil {
			return nil, err
		}
		exist, err := handler.ExistOAuthThirdPartyConds(ctx)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, fmt.Errorf("oauththirdparty exist")
		}
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		user, err := tx.OAuthThirdParty.
			Query().
			Where(
				entoauththirdparty.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if user == nil {
			return fmt.Errorf("invalid user")
		}

		if _, err := oauththirdpartycrud.UpdateSet(
			user.Update(),
			&oauththirdpartycrud.Req{
				ClientID:       h.ClientID,
				ClientSecret:   h.ClientSecret,
				ClientName:     h.ClientName,
				ClientTag:      h.ClientTag,
				ClientLogoURL:  h.ClientLogoURL,
				ClientOAuthURL: h.ClientOAuthURL,
				ResponseType:   h.ResponseType,
				Scope:          h.Scope,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
