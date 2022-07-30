package user

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	entapp "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	entapprole "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approle"
	entapproleuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approleuser"
	entuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"
	entappusercontrol "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusercontrol"
	entextra "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"
	entsecret "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"
	entbanappuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banappuser"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
)

func GetUser(ctx context.Context, appID, userID string) (*UseQueryrResp, error) {
	info, err := GetUsersRealization(ctx, appID, userID, nil, 0, 1)
	if err != nil {
		return nil, err
	}
	if len(info) == 0 {
		logger.Sugar().Errorw("user not found")
		return nil, fmt.Errorf("user not found")
	}
	return info[0], err
}

func GetUsers(ctx context.Context, appID string, offset, limit int32) ([]*UseQueryrResp, error) {
	info, err := GetUsersRealization(ctx, appID, "", nil, offset, limit)
	if err != nil {
		return nil, err
	}

	return info, err
}

func GetManyUsers(ctx context.Context, userIDs []string) ([]*UseQueryrResp, error) {
	info, err := GetUsersRealization(ctx, "", "", userIDs, 0, 0)
	if err != nil {
		return nil, err
	}

	return info, err
}

func GetUsersRealization(ctx context.Context, appID, userID string, userIDs []string, offset, limit int32) ([]*UseQueryrResp, error) {
	var err error
	var userInfos []*UseQueryrResp

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUserInfos")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.AddEvent("call db query")
	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		q := cli.Debug().AppUser.Query().
			Select(
				entuser.FieldID,
				entuser.FieldCreatedAt,
				entuser.FieldAppID,
				entuser.FieldEmailAddress,
				entuser.FieldPhoneNo,
				entuser.FieldImportFromApp,
			)

		if appID != "" {
			q.Where(
				entuser.AppID(uuid.MustParse(appID)),
			)
		}

		if userID != "" {
			q.Where(
				entuser.ID(uuid.MustParse(userID)),
			)
		}

		if userIDs != nil {
			q.Where(
				entuser.IDIn(uuid.MustParse(userID)),
			)
		}

		err = q.Offset(int(offset)).Limit(int(limit)).
			Modify(func(s *sql.Selector) {
				leftExtra(s)
				leftControl(s)
				leftBanApp(s)
				leftSecret(s)
			}).Scan(ctx, &userInfos)
		return err
	})
	if err != nil {
		logger.Sugar().Errorw("fail get user infos:%v", err)
		return nil, err
	}

	respUserIDs := []uuid.UUID{}

	for _, val := range userInfos {
		respUserID, err := uuid.Parse(val.ID)
		if err != nil {
			logger.Sugar().Errorw("userID string to uuid type fail :%v", err)
			return nil, err
		}
		respUserIDs = append(respUserIDs, respUserID)
	}

	roles, err := GetRoles(ctx, entapproleuser.UserIDIn(respUserIDs...))
	if err != nil {
		logger.Sugar().Errorw("fail get roles:%v", err)
		return nil, err
	}

	for key, userInfo := range userInfos {
		for _, role := range roles {
			if userInfo.ID == role.UserID {
				userInfos[key].Role = roles
			}
		}
	}

	return userInfos, nil
}

func GetRoles(ctx context.Context, ps predicate.AppRoleUser) ([]*AppRole, error) {
	var appRole []*AppRole
	var err error

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		err = cli.Debug().AppRoleUser.Query().
			Select(
				entapproleuser.FieldUserID,
			).
			Where(ps).
			Modify(func(s *sql.Selector) {
				role := sql.Table(entapprole.Table)
				s.LeftJoin(role).
					On(
						s.C(entapproleuser.FieldRoleID),
						role.C(entapprole.FieldID),
					).AppendSelect(
					role.C(entapprole.FieldRole),
					role.C(entapprole.FieldCreatedBy),
					role.C(entapprole.FieldDescription),
					role.C(entapprole.FieldDefault),
				).OnP(
					sql.P(func(builder *sql.Builder) {
						builder.Ident(role.C(entapprole.FieldDeletedAt)).WriteOp(sql.OpEQ).Arg(0)
					}),
				)
			}).Scan(ctx, &appRole)
		return err
	})
	if err != nil {
		logger.Sugar().Errorw("fail query roles:%v", err)
		return nil, err
	}
	return appRole, nil
}

func leftExtra(s *sql.Selector) {
	extra := sql.Table(entextra.Table)
	s.LeftJoin(extra).
		On(
			s.C(entapp.FieldID),
			extra.C(entextra.FieldUserID),
		).AppendSelect(
		extra.C(entextra.FieldUsername),
		extra.C(entextra.FieldFirstName),
		extra.C(entextra.FieldLastName),
		extra.C(entextra.FieldAddressFields),
		extra.C(entextra.FieldGender),
		extra.C(entextra.FieldPostalCode),
		extra.C(entextra.FieldAge),
		extra.C(entextra.FieldBirthday),
		extra.C(entextra.FieldAvatar),
		extra.C(entextra.FieldOrganization),
		extra.C(entextra.FieldIDNumber),
	).OnP(
		sql.P(func(builder *sql.Builder) {
			builder.Ident(extra.C(entextra.FieldDeletedAt)).WriteOp(sql.OpEQ).Arg(0)
		}),
	)
}

func leftControl(s *sql.Selector) {
	control := sql.Table(entappusercontrol.Table)
	s.LeftJoin(control).
		On(
			s.C(entapp.FieldID),
			control.C(entappusercontrol.FieldUserID),
		).AppendSelect(
		control.C(entappusercontrol.FieldSigninVerifyByGoogleAuthentication),
		control.C(entappusercontrol.FieldGoogleAuthenticationVerified),
	).OnP(
		sql.P(func(builder *sql.Builder) {
			builder.Ident(control.C(entappusercontrol.FieldDeletedAt)).WriteOp(sql.OpEQ).Arg(0)
		}),
	)
}

func leftBanApp(s *sql.Selector) {
	ban := sql.Table(entbanappuser.Table)
	s.LeftJoin(ban).
		On(
			s.C(entapp.FieldID),
			ban.C(entbanappuser.FieldUserID),
		).AppendSelect(
		sql.As(ban.C(entbanappuser.FieldUserID), "ban_app_user_id"),
	).OnP(
		sql.P(func(builder *sql.Builder) {
			builder.Ident(ban.C(entbanappuser.FieldDeletedAt)).WriteOp(sql.OpEQ).Arg(0)
		}),
	)
}

func leftSecret(s *sql.Selector) {
	secret := sql.Table(entsecret.Table)
	s.LeftJoin(secret).
		On(
			s.C(entapp.FieldID),
			secret.C(entsecret.FieldUserID),
		).AppendSelect(
		sql.As(secret.C(entsecret.FieldGoogleSecret), "has_google_secret"),
	).OnP(
		sql.P(func(builder *sql.Builder) {
			builder.Ident(secret.C(entsecret.FieldDeletedAt)).WriteOp(sql.OpEQ).Arg(0)
		}),
	)
}
