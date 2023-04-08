package user

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/appuser-manager/pkg/encrypt"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	servicename "github.com/NpoolPlatform/appuser-middleware/pkg/servicename"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	entappuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"

	entappusersecret "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"

	usermwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"github.com/google/uuid"
)

type verifyHandler struct {
	*Handler
	stm *ent.AppUserSelect
}

func (h *verifyHandler) queryAppUser(cli *ent.Client) error {
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

func (h *Handler) VerifyAccount(ctx context.Context) (*usermwpb.User, error) {
	type r struct {
		ID           string `sql:"id"`
		AppID        string `sql:"app_id"`
		UserID       string `sql:"user_id"`
		PasswordHash string `sql:"password_hash"`
		Salt         string `sql:"salt"`
	}

	var infos []*r

	handler := &verifyHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryAppUser(cli)
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

func VerifyUser(
	ctx context.Context,
	appID, userID string,
	passwordHash string,
) (
	*usermwpb.User, error,
) {
	type r struct {
		ID           string `sql:"id"`
		AppID        string `sql:"app_id"`
		UserID       string `sql:"user_id"`
		PasswordHash string `sql:"password_hash"`
		Salt         string `sql:"salt"`
	}

	var infos []*r
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "VerifyUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "user", "middleware", "CRUD")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppUser.
			Query().
			Where(
				entappuser.AppID(uuid.MustParse(appID)),
				entappuser.ID(uuid.MustParse(userID)),
			)

		return joinV(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("get user", "err", err.Error())
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		logger.Sugar().Errorw("CreateUser", "err", "too many records")
		return nil, fmt.Errorf("too many records")
	}

	if err := encrypt.VerifyWithSalt(
		passwordHash,
		infos[0].PasswordHash,
		infos[0].Salt,
	); err != nil {
		return nil, err
	}

	return GetUser(ctx, appID, infos[0].UserID)
}

func joinV(stm *ent.AppUserQuery) *ent.AppUserSelect {
	return stm.
		Select(
			entappuser.FieldAppID,
			entappuser.FieldID,
		).
		Limit(1).
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(entappusersecret.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(entappuser.FieldID),
					t1.C(entappusersecret.FieldUserID),
				).
				On(
					s.C(entappuser.FieldAppID),
					t1.C(entappusersecret.FieldAppID),
				).
				AppendSelect(
					sql.As(t1.C(entappusersecret.FieldPasswordHash), "password_hash"),
					sql.As(t1.C(entappusersecret.FieldSalt), "salt"),
					sql.As(t1.C(entappusersecret.FieldUserID), "user_id"),
				)
		})
}
