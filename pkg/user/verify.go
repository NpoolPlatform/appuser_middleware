package user

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/appuser-manager/pkg/encrypt"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	entuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"

	entsecret "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"

	signmethod "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"
	usermwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"github.com/google/uuid"
)

func VerifyAccount(
	ctx context.Context,
	appID, account string,
	accountType signmethod.SignMethodType,
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

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "VerifyAccount")
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
				entuser.AppID(uuid.MustParse(appID)),
			)

		switch accountType {
		case signmethod.SignMethodType_Email:
			stm.Where(entuser.EmailAddress(account))
		case signmethod.SignMethodType_Mobile:
			stm.Where(entuser.PhoneNo(account))
		default:
			return fmt.Errorf("invalid account type")
		}

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
		logger.Sugar().Errorw("VerifyAccount", "err", "too many records")
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

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "VerifyUser")
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
				entuser.AppID(uuid.MustParse(appID)),
				entuser.ID(uuid.MustParse(userID)),
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
			entuser.FieldAppID,
			entuser.FieldID,
		).
		Limit(1).
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(entsecret.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(entuser.FieldID),
					t1.C(entsecret.FieldUserID),
				).
				On(
					s.C(entuser.FieldAppID),
					t1.C(entsecret.FieldAppID),
				).
				AppendSelect(
					sql.As(t1.C(entsecret.FieldPasswordHash), "password_hash"),
					sql.As(t1.C(entsecret.FieldSalt), "salt"),
					sql.As(t1.C(entsecret.FieldUserID), "user_id"),
				)
		})
}
