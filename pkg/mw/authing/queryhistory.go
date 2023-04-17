package authing

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	entapp "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/app"
	entappuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuser"
	entauth "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/auth"
	entauthhistory "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/authhistory"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing"
)

type queryHistoryHandler struct {
	*Handler
	stm   *ent.AuthHistorySelect
	infos []*npool.History
	total uint32
}

func (h *queryHistoryHandler) selectAuthHistory(stm *ent.AuthHistoryQuery) {
	h.stm = stm.Select(
		entauthhistory.FieldAppID,
		entauthhistory.FieldUserID,
		entauthhistory.FieldResource,
		entauthhistory.FieldMethod,
		entauthhistory.FieldAllowed,
		entauthhistory.FieldCreatedAt,
	)
}

func (h *queryHistoryHandler) queryAuthHistories(ctx context.Context, cli *ent.Client) error {
	stm := cli.
		AuthHistory.
		Query().
		Where(
			entauthhistory.AppID(h.AppID),
		)

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectAuthHistory(stm)
	return nil
}

func (h *queryHistoryHandler) queryJoinApp(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			s.C(entauth.FieldAppID),
			t.C(entapp.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entapp.FieldName), "app_name"),
			sql.As(t.C(entapp.FieldLogo), "app_logo"),
		)
}

func (h *queryHistoryHandler) queryJoinAppUser(s *sql.Selector) {
	t := sql.Table(entappuser.Table)
	s.LeftJoin(t).
		On(
			s.C(entauth.FieldUserID),
			t.C(entappuser.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entappuser.FieldEmailAddress), "email_address"),
			sql.As(t.C(entappuser.FieldPhoneNo), "phone_no"),
		)
}

func (h *queryHistoryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinApp(s)
		h.queryJoinAppUser(s)
	})
}

func (h *queryHistoryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetHistories(ctx context.Context) ([]*npool.History, uint32, error) {
	handler := &queryHistoryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAuthHistories(ctx, cli); err != nil {
			return nil
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
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
