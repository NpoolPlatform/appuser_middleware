//nolint:dupl
package auth

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	entapp "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/app"
	entapproleuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approleuser"
	entappuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuser"
	entauth "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/auth"
	entbanapp "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/banapp"
	entbanappuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/banappuser"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

type auth struct {
	AppID   string `sql:"app_id"`
	RoleID  string `sql:"role_id"`
	UserID  string `sql:"user_id"`
	AppVID  string `sql:"app_vid"`
	AppBID  string `sql:"app_bid"`
	UserVID string `sql:"user_vid"`
	UserBID string `sql:"user_bid"`
}

type existHandler struct {
	*Handler
	infos []*auth
}

func (h *existHandler) queryJoinApp(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			t.C(entapp.FieldID),
			s.C(entappuser.FieldAppID),
		).
		AppendSelect(
			sql.As(t.C(entapp.FieldID), "app_vid"),
		)
}

func (h *existHandler) queryJoinBanApp(s *sql.Selector) {
	t := sql.Table(entbanapp.Table)
	s.LeftJoin(t).
		On(
			t.C(entbanapp.FieldAppID),
			s.C(entappuser.FieldAppID),
		).
		AppendSelect(
			sql.As(t.C(entbanapp.FieldAppID), "app_bid"),
		)
}

func (h *existHandler) queryJoinBanAppUser(s *sql.Selector) {
	t := sql.Table(entbanappuser.Table)
	s.LeftJoin(t).
		On(
			t.C(entbanappuser.FieldAppID),
			s.C(entappuser.FieldAppID),
		).
		On(
			t.C(entbanappuser.FieldUserID),
			s.C(entappuser.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entbanappuser.FieldUserID), "user_bid"),
		)
}

type existUserHandler struct {
	*existHandler
	stm *ent.AppUserSelect
}

func (h *existUserHandler) queryJoinAuth(s *sql.Selector) {
	t := sql.Table(entauth.Table)
	s.LeftJoin(t).
		On(
			t.C(entauth.FieldAppID),
			s.C(entappuser.FieldAppID),
		).
		On(
			t.C(entauth.FieldUserID),
			s.C(entappuser.FieldID),
		).
		Where(
			sql.And(
				sql.EQ(s.C(entappuser.FieldAppID), h.AppID),
				sql.EQ(t.C(entauth.FieldUserID), *h.UserID),
				sql.EQ(t.C(entauth.FieldResource), h.Resource),
				sql.EQ(t.C(entauth.FieldMethod), h.Method),
			),
		)
}

func (h *existUserHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *existUserHandler) queryAppUser(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid id")
	}

	h.stm = cli.
		AppUser.
		Query().
		Where(
			entappuser.AppID(h.AppID),
			entappuser.ID(*h.ID),
		).
		Select(
			entappuser.FieldAppID,
			entappuser.FieldID,
		)
	return nil
}

func (h *existUserHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinApp(s)
		h.queryJoinBanApp(s)
		h.queryJoinBanAppUser(s)
		h.queryJoinAuth(s)
	})
}

func (h *existHandler) existUserAuth(ctx context.Context) (bool, error) {
	handler := &existUserHandler{
		existHandler: h,
	}

	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppUser(cli); err != nil {
			return err
		}
		handler.queryJoin()
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		logger.Sugar().Errorw("existUserAuth", "error", err)
		return false, err
	}
	if len(h.infos) == 0 {
		logger.Sugar().Infow("existUserAuth", "Reason", "no record")
		return false, nil
	}
	if h.infos[0].AppBID == h.infos[0].AppID {
		logger.Sugar().Infow("existUserAuth", "Reason", "banned appid")
		return false, nil
	}
	if h.infos[0].AppID != h.infos[0].AppVID {
		logger.Sugar().Infow("existUserAuth", "Reason", "mismatch appid")
		return false, nil
	}
	if h.infos[0].UserBID == h.infos[0].UserID {
		logger.Sugar().Infow("existUserAuth", "Reason", "banned userid")
		return false, nil
	}

	return true, nil
}

type existRoleHandler struct {
	*existHandler
	stm *ent.AppRoleUserSelect
}

func (h *existRoleHandler) queryAppRoleUser(cli *ent.Client) error {
	if h.UserID == nil {
		return fmt.Errorf("invalid user")
	}

	h.stm = cli.
		AppRoleUser.
		Query().
		Where(
			entapproleuser.AppID(h.AppID),
			entapproleuser.UserID(*h.UserID),
		).
		Select(
			entapproleuser.FieldAppID,
			entapproleuser.FieldRoleID,
			entapproleuser.FieldUserID,
		)
	return nil
}

func (h *existRoleHandler) queryJoinAppUser(s *sql.Selector) {
	t := sql.Table(entappuser.Table)
	s.LeftJoin(t).
		On(
			t.C(entappuser.FieldAppID),
			s.C(entapproleuser.FieldAppID),
		).
		On(
			t.C(entappuser.FieldID),
			s.C(entapproleuser.FieldUserID),
		).
		AppendSelect(
			sql.As(t.C(entappuser.FieldID), "user_vid"),
		)
}

func (h *existRoleHandler) queryJoinAuth(s *sql.Selector) {
	t := sql.Table(entauth.Table)
	s.LeftJoin(t).
		On(
			t.C(entauth.FieldAppID),
			s.C(entapproleuser.FieldAppID),
		).
		On(
			t.C(entauth.FieldRoleID),
			s.C(entapproleuser.FieldRoleID),
		).
		Where(
			sql.And(
				sql.EQ(s.C(entapproleuser.FieldAppID), h.AppID),
				sql.EQ(s.C(entapproleuser.FieldUserID), *h.UserID),
				sql.EQ(t.C(entauth.FieldResource), h.Resource),
				sql.EQ(t.C(entauth.FieldMethod), h.Method),
			),
		)
}

func (h *existRoleHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinApp(s)
		h.queryJoinBanApp(s)
		h.queryJoinAppUser(s)
		h.queryJoinBanAppUser(s)
		h.queryJoinAuth(s)
	})
}

func (h *existRoleHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *existHandler) existRoleAuth(ctx context.Context) (bool, error) {
	handler := &existRoleHandler{
		existHandler: h,
	}

	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppRoleUser(cli); err != nil {
			return err
		}
		handler.queryJoin()
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logger.Sugar().Errorw("existRoleAuth", "error", err)
		return false, err
	}
	if len(h.infos) == 0 {
		logger.Sugar().Infow("existRoleAuth", "Reason", "no record")
		return false, nil
	}
	if h.infos[0].AppBID == h.infos[0].AppID {
		logger.Sugar().Infow("existRoleAuth", "Reason", "banned appid")
		return false, nil
	}
	if h.infos[0].AppID != h.infos[0].AppVID {
		logger.Sugar().Infow("existRoleAuth", "Reason", "mismatch appid")
		return false, nil
	}
	if h.infos[0].UserBID == h.infos[0].UserID {
		logger.Sugar().Infow("existRoleAuth", "Reason", "banned userid")
		return false, nil
	}
	if h.infos[0].UserID != h.infos[0].UserVID {
		logger.Sugar().Infow("existRoleAuth", "Reason", "mismatch userid")
		return false, nil
	}

	return true, nil
}

type existAppHandler struct {
	*existHandler
	stm *ent.AppSelect
}

func (h *existAppHandler) queryApp(cli *ent.Client) {
	h.stm = cli.
		App.
		Query().
		Where(
			entapp.ID(h.AppID),
		).
		Select(
			sql.As(entapp.FieldID, "app_id"),
		)
}

func (h *existAppHandler) queryJoinAuth(s *sql.Selector) {
	t := sql.Table(entauth.Table)
	s.LeftJoin(t).
		On(
			t.C(entauth.FieldAppID),
			s.C(entapproleuser.FieldAppID),
		).
		AppendSelect(
			sql.As(t.C(entauth.FieldAppID), "app_vid"),
		).
		Where(
			sql.And(
				sql.EQ(s.C(entapproleuser.FieldAppID), h.AppID),
				sql.EQ(s.C(entauth.FieldUserID), uuid.UUID{}),
				sql.EQ(s.C(entauth.FieldRoleID), uuid.UUID{}),
				sql.EQ(t.C(entauth.FieldResource), h.Resource),
				sql.EQ(t.C(entauth.FieldMethod), h.Method),
			),
		)
}

func (h *existAppHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinBanApp(s)
		h.queryJoinAuth(s)
	})
}

func (h *existAppHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *existHandler) existAppAuth(ctx context.Context) (bool, error) {
	handler := &existAppHandler{
		existHandler: h,
	}

	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		handler.queryApp(cli)
		handler.queryJoin()
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		logger.Sugar().Errorw("existAppAuth", "error", err)
		return false, err
	}
	if len(h.infos) == 0 {
		logger.Sugar().Infow("existAppAuth", "Reason", "no record")
		return false, nil
	}
	if h.infos[0].AppBID == h.infos[0].AppID {
		logger.Sugar().Infow("existAppAuth", "Reason", "banned appid")
		return false, nil
	}
	if h.infos[0].AppID != h.infos[0].AppVID {
		logger.Sugar().Infow("existAppAuth", "Reason", "mismatch appid")
		return false, nil
	}

	return true, nil
}

func (h *Handler) ExistAuth(ctx context.Context) (bool, error) {
	handler := &existHandler{
		Handler: h,
	}
	if h.UserID != nil {
		exist, err := handler.existUserAuth(ctx)
		if err != nil {
			return false, err
		}
		if exist {
			return true, nil
		}
	}
	exist, err := handler.existRoleAuth(ctx)
	if err != nil {
		return false, err
	}
	if exist {
		return true, nil
	}
	return handler.existAppAuth(ctx)
}
