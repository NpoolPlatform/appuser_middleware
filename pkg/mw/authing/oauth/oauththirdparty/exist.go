package oauththirdparty

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	oauththirdpartycrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/authing/oauth/oauththirdparty"
	entoauththirdparty "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/oauththirdparty"
)

type existHandler struct {
	*Handler
	stm *ent.OAuthThirdPartyQuery
}

func (h *existHandler) queryOAuthThirdParty(cli *ent.Client) {
	h.stm = cli.OAuthThirdParty.
		Query().
		Where(
			entoauththirdparty.ID(*h.ID),
			entoauththirdparty.DeletedAt(0),
		)
}

func (h *existHandler) queryOAuthThirdParties(cli *ent.Client) error {
	stm, err := oauththirdpartycrud.SetQueryConds(cli.OAuthThirdParty.Query(), h.Conds)
	if err != nil {
		return err
	}
	h.stm = stm
	return nil
}

func (h *Handler) ExistOAuthThirdParty(ctx context.Context) (bool, error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryOAuthThirdParty(cli)
		_exist, err := handler.stm.Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistOAuthThirdPartyConds(ctx context.Context) (bool, error) {
	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOAuthThirdParties(cli); err != nil {
			return err
		}
		_exist, err := handler.stm.Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
