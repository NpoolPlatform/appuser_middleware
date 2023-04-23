package user

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	entapp "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/app"
	entapprole "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approle"
	entapproleuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approleuser"
	entappuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuser"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role/user"

	roleusercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role/user"
)

type queryHandler struct {
	*Handler
	stm   *ent.AppRoleUserSelect
	infos []*npool.User
	total uint32
}

func (h *queryHandler) selectAppRoleUser(stm *ent.AppRoleUserQuery) {
	h.stm = stm.Select(
		entapproleuser.FieldID,
	)
}

func (h *queryHandler) queryAppRoleUser(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid roleuserid")
	}

	h.selectAppRoleUser(
		cli.AppRoleUser.
			Query().
			Where(
				entapproleuser.ID(*h.ID),
				entapproleuser.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryAppRoleUsers(cli *ent.Client) error {
	stm, err := roleusercrud.SetQueryConds(cli.AppRoleUser.Query(), h.Conds)
	if err != nil {
		return err
	}
	h.selectAppRoleUser(stm)
	return nil
}

func (h *queryHandler) queryJoinAppRole(s *sql.Selector) {
	t := sql.Table(entapprole.Table)
	stm := s.LeftJoin(t).
		On(
			s.C(entapproleuser.FieldRoleID),
			t.C(entapprole.FieldID),
		)
	if h.Conds.Genesis != nil {
		stm.Where(
			sql.EQ(t.C(entapprole.FieldGenesis), h.Conds.Genesis.Val.(bool)),
		)
	}

	stm.AppendSelect(
		t.C(entapprole.FieldCreatedBy),
		t.C(entapprole.FieldRole),
		t.C(entapprole.FieldDescription),
		t.C(entapprole.FieldDefault),
		t.C(entapprole.FieldGenesis),
		sql.As(t.C(entapprole.FieldID), "role_id"),
	)
}

func (h *queryHandler) queryJoinApp(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			s.C(entapproleuser.FieldAppID),
			t.C(entapp.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entapp.FieldID), "app_id"),
			sql.As(t.C(entapp.FieldName), "app_name"),
			sql.As(t.C(entapp.FieldLogo), "app_logo"),
			t.C(entapp.FieldCreatedAt),
		)
}

func (h *queryHandler) queryJoinAppUser(s *sql.Selector) {
	t := sql.Table(entappuser.Table)
	s.LeftJoin(t).
		On(
			s.C(entapproleuser.FieldUserID),
			t.C(entappuser.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entappuser.FieldID), "user_id"),
			sql.As(t.C(entappuser.FieldEmailAddress), "email_address"),
			t.C(entappuser.FieldPhoneNo),
		)
}

func (h *queryHandler) queryJoin(ctx context.Context) error {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinAppRole(s)
		h.queryJoinApp(s)
		h.queryJoinAppUser(s)
	})
	total, err := h.stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetUser(ctx context.Context) (*npool.User, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppRoleUser(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(ctx); err != nil {
			return err
		}
		if err := handler.scan(ctx); err != nil {
			return nil
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
		return nil, fmt.Errorf("too many record")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetUsers(ctx context.Context) ([]*npool.User, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppRoleUsers(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(ctx); err != nil {
			return err
		}
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
