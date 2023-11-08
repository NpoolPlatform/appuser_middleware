package user

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	entapprole "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approle"
	entapproleuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approleuser"

	roleusercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role/user"
)

type existHandler struct {
	*Handler
	stm *ent.AppRoleUserSelect
}

func (h *existHandler) selectAppRoleUser(stm *ent.AppRoleUserQuery) {
	h.stm = stm.Select(entapproleuser.FieldID)
}

func (h *existHandler) queryAppRoleUser(cli *ent.Client) error {
	stm := cli.AppRoleUser.
		Query().
		Where(entapproleuser.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entapproleuser.ID(*h.ID))
	}
	if h.AppID != nil {
		stm.Where(entapproleuser.AppID(*h.AppID))
	}
	if h.EntID != nil {
		stm.Where(entapproleuser.EntID(*h.EntID))
	}
	h.selectAppRoleUser(stm)
	return nil
}

func (h *existHandler) queryAppRoleUsers(ctx context.Context, cli *ent.Client) error {
	stm, err := roleusercrud.SetQueryConds(cli.AppRoleUser.Query(), h.Conds)
	if err != nil {
		return err
	}
	h.selectAppRoleUser(stm)
	return nil
}

func (h *existHandler) queryJoinAppRole(s *sql.Selector) {
	t := sql.Table(entapprole.Table)
	stm := s.LeftJoin(t).
		On(
			s.C(entapproleuser.FieldRoleID),
			t.C(entapprole.FieldEntID),
		)
	stm.AppendSelect(
		t.C(entapprole.FieldCreatedBy),
		t.C(entapprole.FieldRole),
		t.C(entapprole.FieldDescription),
		t.C(entapprole.FieldDefault),
		t.C(entapprole.FieldGenesis),
		sql.As(t.C(entapprole.FieldEntID), "role_id"),
	)
	if h.Conds != nil && h.Conds.Genesis != nil {
		stm.Where(
			sql.EQ(t.C(entapprole.FieldGenesis), h.Conds.Genesis.Val.(bool)),
		)
	}
}

func (h *existHandler) queryJoin(ctx context.Context) {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinAppRole(s)
	})
}

func (h *Handler) ExistUserConds(ctx context.Context) (bool, error) {
	handler := &existHandler{
		Handler: h,
	}

	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppRoleUsers(ctx, cli); err != nil {
			return err
		}
		handler.queryJoin(ctx)
		exist, err = handler.stm.Exist(ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
