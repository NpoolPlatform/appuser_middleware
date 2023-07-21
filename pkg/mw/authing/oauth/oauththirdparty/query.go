package oauththirdparty

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entoauththirdparty "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/oauththirdparty"

	oauththirdpartycrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/authing/oauth/oauththirdparty"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/oauth/oauththirdparty"
)

type queryHandler struct {
	*Handler
	stm   *ent.OAuthThirdPartySelect
	infos []*npool.OAuthThirdParty
	total uint32
}

func (h *queryHandler) selectOAuthThirdParty(stm *ent.OAuthThirdPartyQuery) {
	h.stm = stm.Select(
		entoauththirdparty.FieldID,
		entoauththirdparty.FieldClientName,
		entoauththirdparty.FieldClientTag,
		entoauththirdparty.FieldClientLogoURL,
		entoauththirdparty.FieldClientOauthURL,
		entoauththirdparty.FieldResponseType,
		entoauththirdparty.FieldScope,
		entoauththirdparty.FieldCreatedAt,
		entoauththirdparty.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryOAuthThirdParty(cli *ent.Client) {
	h.selectOAuthThirdParty(
		cli.OAuthThirdParty.
			Query().
			Where(
				entoauththirdparty.ID(*h.ID),
				entoauththirdparty.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryOAuthThirdParties(ctx context.Context, cli *ent.Client) error {
	stm, err := oauththirdpartycrud.SetQueryConds(cli.OAuthThirdParty.Query(), h.Conds)
	if err != nil {
		return err
	}

	_total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(_total)
	h.selectOAuthThirdParty(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetOAuthThirdParty(ctx context.Context) (*npool.OAuthThirdParty, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryOAuthThirdParty(cli)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetOAuthThirdParties(ctx context.Context) ([]*npool.OAuthThirdParty, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOAuthThirdParties(_ctx, cli); err != nil {
			return err
		}

		handler.
			stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Order(ent.Desc(entoauththirdparty.FieldCreatedAt))
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}
