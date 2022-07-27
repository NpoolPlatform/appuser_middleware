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

type AppRole struct {
	UserID      string `json:"user_id"`
	Role        string `json:"role"`
	CreatedBy   string `json:"created_by"`
	Description string `json:"description"`
	Default     int32  `json:"default"`
}

type Info struct {
	ID                                 string `json:"id"`
	AppID                              string `json:"app_id"`
	EmailAddress                       string `json:"email_address"`
	PhoneNO                            string `json:"phone_no"`
	ImportFromApp                      string `json:"import_from_app"`
	CreatedAt                          uint32 `json:"created_at"`
	Username                           string `json:"username"`
	AddressFields                      string `json:"address_fields"`
	Gender                             string `json:"gender"`
	PostalCode                         string `json:"postal_code"`
	Age                                uint32 `json:"age"`
	Birthday                           uint32 `json:"birthday"`
	Avatar                             string `json:"avatar"`
	Organization                       string `json:"organization"`
	FirstName                          string `json:"first_name"`
	LastName                           string `json:"last_name"`
	IDNumber                           string `json:"id_number"`
	SigninVerifyByGoogleAuthentication uint32 `json:"signin_verify_by_google_authentication"`
	GoogleAuthenticationVerified       uint32 `json:"google_authentication_verified"`
	BanAppUserID                       string `json:"ban_app_user_id"`
	HasGoogleSecret                    string `json:"has_google_secret"`
	Role                               []*AppRole
}

func GetUserInfo(ctx context.Context, appID, userID uuid.UUID) (*Info, error) {
	var err error
	var userInfo []Info

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUserInfo")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.AddEvent("call db query")
	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		err = cli.Debug().AppUser.Query().
			Select(
				entuser.FieldID,
				entuser.FieldCreatedAt,
				entuser.FieldAppID,
				entuser.FieldEmailAddress,
				entuser.FieldPhoneNo,
				entuser.FieldImportFromApp,
			).Where(
			entuser.AppID(appID),
			entuser.ID(userID),
		).
			Modify(func(s *sql.Selector) {
				leftExtra(s)
				leftControl(s)
				leftBanApp(s)
				leftSecret(s)
			}).Scan(ctx, &userInfo)
		return err
	})
	if err != nil {
		logger.Sugar().Errorw("fail get user info: %v", err)
		return nil, err
	}

	if len(userInfo) == 0 {
		logger.Sugar().Errorw("user not found")
		return nil, fmt.Errorf("user not found")
	}

	appRole, err := GetRoles(ctx, entapproleuser.UserID(userID))
	if err != nil {
		logger.Sugar().Errorw("fail get roles :%v", err)
		return nil, err
	}

	userInfo[0].Role = appRole
	return &userInfo[0], nil
}

func GetUserInfos(ctx context.Context, appID uuid.UUID, offset, limit int32) ([]*Info, error) {
	var err error
	var userInfos []*Info

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
		err = cli.Debug().AppUser.Query().
			Select(
				entuser.FieldID,
				entuser.FieldCreatedAt,
				entuser.FieldAppID,
				entuser.FieldEmailAddress,
				entuser.FieldPhoneNo,
				entuser.FieldImportFromApp,
			).Where(
			entuser.AppID(appID),
		).Offset(int(offset)).Limit(int(limit)).
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

	userIDs := []uuid.UUID{}

	for _, val := range userInfos {
		userID, err := uuid.Parse(val.ID)
		if err != nil {
			logger.Sugar().Errorw("userID string to uuid type fail :%v", err)
			return nil, err
		}
		userIDs = append(userIDs, userID)
	}

	roles, err := GetRoles(ctx, entapproleuser.UserIDIn(userIDs...))
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
