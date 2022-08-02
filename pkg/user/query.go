package user

import (
	"context"
	"fmt"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	entapprole "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approle"
	entapproleuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approleuser"

	entapp "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	entuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"
	entappusercontrol "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusercontrol"
	entextra "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"

	entsecret "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"
	entbanappuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banappuser"

	"github.com/google/uuid"
)

func GetUser(ctx context.Context, appID, userID string) (*User, error) {
	var infos []*User
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "user", "middleware", "query join")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppUser.
			Query().
			Where(
				entuser.ID(uuid.MustParse(userID)),
				entuser.AppID(uuid.MustParse(appID)),
			).
			Limit(1)

		return join(stm).
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

	span = commontracer.TraceInvoker(span, "user", "method", "expand")

	infos, err = expand(ctx, []string{userID}, infos)
	if err != nil {
		return nil, err
	}

	return infos[0], nil
}

func GetUsers(ctx context.Context, appID string, offset, limit int32) ([]*User, error) {
	var infos []*User
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "user", "db", "query join")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppUser.
			Query().
			Where(
				entuser.AppID(uuid.MustParse(appID)),
			).
			Offset(int(offset)).
			Limit(int(limit))

		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("CreateUser", "err", err.Error())
		return nil, err
	}

	users := []string{}
	for _, info := range infos {
		users = append(users, info.ID.String())
	}

	span = commontracer.TraceInvoker(span, "user", "method", "expand")

	infos, err = expand(ctx, users, infos)
	if err != nil {
		return nil, err
	}

	return infos, nil
}

func GetManyUsers(ctx context.Context, userIDs []string) ([]*User, error) {
	var infos []*User
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetManyUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	users := []uuid.UUID{}
	for _, user := range userIDs {
		users = append(users, uuid.MustParse(user))
	}

	span = commontracer.TraceInvoker(span, "user", "db", "query join")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppUser.
			Query().
			Where(
				entuser.IDIn(users...),
			)

		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("CreateUser", "err", err.Error())
		return nil, err
	}

	for _, info := range infos {
		info.Banned = info.BanAppUserID.String() != uuid.UUID{}.String()
	}

	span = commontracer.TraceInvoker(span, "user", "method", "expand")

	infos, err = expand(ctx, userIDs, infos)
	if err != nil {
		return nil, err
	}

	return infos, nil
}

func expand(ctx context.Context, userIDs []string, users []*User) ([]*User, error) {
	type extra struct {
		UserID       uuid.UUID `json:"user_id"`
		GoogleSecret string    `json:"google_secret"`
		RoleName     string    `json:"role_name"`
	}

	var infos []*extra
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "expand")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	uids := []uuid.UUID{}
	for _, user := range userIDs {
		uids = append(uids, uuid.MustParse(user))
	}

	span = commontracer.TraceInvoker(span, "user", "db", "query join")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		return cli.
			AppUserSecret.
			Query().
			Where(
				entsecret.UserIDIn(uids...),
			).
			Select(
				entsecret.FieldUserID,
				entsecret.FieldGoogleSecret,
			).
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(entapproleuser.Table)
				t2 := sql.Table(entapprole.Table)

				s.
					LeftJoin(t1).
					On(
						s.C(entsecret.FieldUserID),
						t1.C(entapproleuser.FieldUserID),
					).
					LeftJoin(t2).
					On(
						t1.C(entapproleuser.FieldRoleID),
						t2.C(entapprole.FieldID),
					).
					AppendSelect(
						sql.As(t2.C(entapprole.FieldRole), "role_name"),
					)
			}).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("expand", "err", err.Error())
		return nil, err
	}

	for _, info := range infos {
		for _, user := range users {
			if info.UserID == user.ID {
				user.HasGoogleSecret = info.GoogleSecret != ""
				user.Roles = append(user.Roles, info.RoleName)
				break
			}
		}
	}

	return users, nil
}

func join(stm *ent.AppUserQuery) *ent.AppUserSelect {
	return stm.
		Select(
			entuser.FieldID,
			entuser.FieldEmailAddress,
			entuser.FieldPhoneNo,
			entuser.FieldImportFromApp,
			entuser.FieldCreatedAt,
		).
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(entextra.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(entuser.FieldID),
					t1.C(entextra.FieldUserID),
				).
				AppendSelect(
					sql.As(t1.C(entextra.FieldUsername), "username"),
					sql.As(t1.C(entextra.FieldFirstName), "first_name"),
					sql.As(t1.C(entextra.FieldLastName), "last_name"),
					sql.As(t1.C(entextra.FieldAddressFields), "address_fields"),
					sql.As(t1.C(entextra.FieldGender), "gender"),
					sql.As(t1.C(entextra.FieldPostalCode), "postal_code"),
					sql.As(t1.C(entextra.FieldAge), "age"),
					sql.As(t1.C(entextra.FieldBirthday), "birthday"),
					sql.As(t1.C(entextra.FieldAvatar), "avatar"),
					sql.As(t1.C(entextra.FieldOrganization), "organization"),
					sql.As(t1.C(entextra.FieldIDNumber), "id_number"),
				)

			t2 := sql.Table(entappusercontrol.Table)
			s.
				LeftJoin(t2).
				On(
					s.C(entuser.FieldID),
					t2.C(entappusercontrol.FieldUserID),
				).
				AppendSelect(
					sql.As(t2.C(entappusercontrol.FieldSigninVerifyByGoogleAuthentication), "signin_verify_by_google_authentication"),
					sql.As(t2.C(entappusercontrol.FieldGoogleAuthenticationVerified), "google_authentication_verified"),
				)

			t3 := sql.Table(entapp.Table)
			s.
				LeftJoin(t3).
				On(
					s.C(entuser.FieldImportFromApp),
					t2.C(entapp.FieldID),
				).
				AppendSelect(
					sql.As(t3.C(entapp.FieldName), "imported_from_app_name"),
					sql.As(t3.C(entapp.FieldLogo), "imported_from_app_logo"),
				)

			t4 := sql.Table(entbanappuser.Table)
			s.
				LeftJoin(t4).
				On(
					s.C(entuser.FieldID),
					t2.C(entbanappuser.FieldUserID),
				).
				AppendSelect(
					sql.As(t4.C(entbanappuser.FieldID), "ban_app_user_id"),
					sql.As(t4.C(entbanappuser.FieldMessage), "ban_message"),
				)
		})
}
