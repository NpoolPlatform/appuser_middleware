package authing

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	entapp "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/app"
	entapprole "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approle"
	entappuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuser"
	entauth "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/auth"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing"
)

type queryHandler struct {
	*Handler
	stm   *ent.AuthSelect
	infos []*npool.Auth
	total uint32
}

func (h *queryHandler) selectAuth(stm *ent.AuthQuery) {
	h.stm = stm.Select(
		entauth.FieldID,
		entauth.FieldResource,
		entauth.FieldMethod,
		entauth.FieldCreatedAt,
		entauth.FieldAppID,
		entauth.FieldRoleID,
		entauth.FieldUserID,
	)
}

func (h *queryHandler) queryAuth(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid id")
	}

	h.selectAuth(
		cli.Auth.
			Query().
			Where(
				entauth.ID(*h.ID),
				entauth.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryAuths(ctx context.Context, cli *ent.Client) error {
	stm := cli.
		Auth.
		Query().
		Where(
			entauth.AppID(h.AppID),
		)

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectAuth(stm)
	return nil
}

func (h *queryHandler) queryJoinApp(s *sql.Selector) {
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

func (h *queryHandler) queryJoinAppRole(s *sql.Selector) {
	t := sql.Table(entapprole.Table)
	s.LeftJoin(t).
		On(
			s.C(entauth.FieldRoleID),
			t.C(entapprole.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entapprole.FieldRole), "role_name"),
		)
}

func (h *queryHandler) queryJoinAppUser(s *sql.Selector) {
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

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinApp(s)
		h.queryJoinAppRole(s)
		h.queryJoinAppUser(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetAuth(ctx context.Context) (*npool.Auth, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAuth(cli); err != nil {
			return nil
		}
		handler.queryJoin()
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
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

func (h *Handler) GetAuths(ctx context.Context) ([]*npool.Auth, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAuths(ctx, cli); err != nil {
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
