package role

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	rolecrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role"
	entapp "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/app"
	entapprole "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approle"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"
)

type queryHandler struct {
	*Handler
	stm   *ent.AppRoleSelect
	infos []*npool.Role
	total uint32
}

func (h *queryHandler) selectAppRole(stm *ent.AppRoleQuery) {
	h.stm = stm.Select(
		entapprole.FieldID,
		entapprole.FieldCreatedBy,
		entapprole.FieldRole,
		entapprole.FieldDescription,
		entapprole.FieldDefault,
		entapprole.FieldGenesis,
	)
}

func (h *queryHandler) queryAppRole(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid roleid")
	}

	h.selectAppRole(
		cli.AppRole.
			Query().
			Where(
				entapprole.ID(*h.ID),
				entapprole.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryAppRoles(ctx context.Context, cli *ent.Client) error {
	stm, err := rolecrud.SetQueryConds(cli.AppRole.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectAppRole(stm)
	return nil
}

func (h *queryHandler) queryJoinApp(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			s.C(entapprole.FieldAppID),
			t.C(entapp.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entapp.FieldID), "app_id"),
			sql.As(t.C(entapp.FieldName), "app_name"),
			sql.As(t.C(entapp.FieldLogo), "app_logo"),
			t.C(entapp.FieldCreatedAt),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinApp(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetRole(ctx context.Context) (*npool.Role, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppRole(cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(handler.Offset)).
			Limit(2).
			Modify(func(s *sql.Selector) {})
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

func (h *Handler) GetRoles(ctx context.Context) ([]*npool.Role, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppRoles(ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
		handler.stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Modify(func(s *sql.Selector) {})
		if err := handler.scan(ctx); err != nil {
			return nil
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}
