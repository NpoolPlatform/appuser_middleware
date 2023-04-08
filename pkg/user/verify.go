package user

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/encrypt"

	entappuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"
	entappusersecret "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"

	usermwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

type verifyHandler struct {
	*Handler
	stm *ent.AppUserSelect
}

func (h *verifyHandler) queryAppUserByAccount(cli *ent.Client) error {
	if h.EmailAddress == nil && h.PhoneNO == nil {
		return fmt.Errorf("invalid account")
	}

	stm := cli.
		AppUser.
		Query().
		Where(
			entappuser.AppID(uuid.MustParse(h.AppID)),
			entappuser.DeletedAt(0),
		)
	if h.EmailAddress != nil {
		stm = stm.Where(
			entappuser.EmailAddress(*h.EmailAddress),
		)
	}
	if h.PhoneNO != nil {
		stm = stm.Where(
			entappuser.PhoneNo(*h.PhoneNO),
		)
	}
	h.stm = stm.Select(
		entappuser.FieldID,
		entappuser.FieldAppID,
	)
	return nil
}

func (h *verifyHandler) queryAppUserByID(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid user id")
	}

	h.stm = cli.
		AppUser.
		Query().
		Where(
			entappuser.AppID(uuid.MustParse(h.AppID)),
			entappuser.ID(uuid.MustParse(*h.ID)),
			entappuser.DeletedAt(0),
		).
		Select(
			entappuser.FieldID,
			entappuser.FieldAppID,
		)
	return nil
}

func (h *verifyHandler) queryJoinAppUserSecret() {
	h.stm.Modify(func(s *sql.Selector) {
		t := sql.Table(entappusersecret.Table)
		s.LeftJoin(t).
			On(
				s.C(entappuser.FieldID),
				t.C(entappusersecret.FieldUserID),
			).
			On(
				s.C(entappuser.FieldAppID),
				t.C(entappusersecret.FieldAppID),
			).
			AppendSelect(
				sql.As(t.C(entappusersecret.FieldPasswordHash), "password_hash"),
				sql.As(t.C(entappusersecret.FieldSalt), "salt"),
				sql.As(t.C(entappusersecret.FieldUserID), "user_id"),
			)
	})
}

type r struct {
	ID           string `sql:"id"`
	AppID        string `sql:"app_id"`
	UserID       string `sql:"user_id"`
	PasswordHash string `sql:"password_hash"`
	Salt         string `sql:"salt"`
}

func (h *Handler) VerifyAccount(ctx context.Context) (*usermwpb.User, error) {
	var infos []*r

	handler := &verifyHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryAppUserByAccount(cli)
		handler.queryJoinAppUserSecret()
		return handler.stm.Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("invalid user")
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	if h.PasswordHash == nil {
		return nil, fmt.Errorf("invalid password")
	}

	if err := encrypt.VerifyWithSalt(
		*h.PasswordHash,
		infos[0].PasswordHash,
		infos[0].Salt,
	); err != nil {
		return nil, err
	}

	h.ID = &infos[0].UserID
	return h.GetUser(ctx)
}

func (h *Handler) VerifyUser(ctx context.Context) (*usermwpb.User, error) {
	var infos []*r

	handler := &verifyHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryAppUserByID(cli)
		handler.queryJoinAppUserSecret()
		return handler.stm.Scan(_ctx, &infos)
	})

	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("invalid user")
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	if h.PasswordHash == nil {
		return nil, fmt.Errorf("invalid password")
	}

	if err := encrypt.VerifyWithSalt(
		*h.PasswordHash,
		infos[0].PasswordHash,
		infos[0].Salt,
	); err != nil {
		return nil, err
	}

	h.ID = &infos[0].UserID
	return h.GetUser(ctx)
}
