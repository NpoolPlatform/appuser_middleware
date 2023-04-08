package user

import (
	"context"
	"fmt"

	mgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	servicename "github.com/NpoolPlatform/appuser-middleware/pkg/servicename"
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
	entappusersecret "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"

	entsecret "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"
	entbanappuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banappuser"
	entkyc "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/kyc"
	usermwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stm   *ent.AppUserSelect
	infos []*usermwpb.User
}

func (h *queryHandler) queryAppUser(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid userid")
	}

	h.stm = cli.
		AppUser.
		Query().
		Where(
			entuser.AppID(uuid.MustParse(h.AppID)),
			entuser.ID(uuid.MustParse(*h.ID)),
		).
		Select(
			entuser.FieldID,
			entuser.FieldAppID,
			entuser.FieldEmailAddress,
			entuser.FieldPhoneNo,
			entuser.FieldImportFromApp,
			entuser.FieldCreatedAt,
		)
	return nil
}

func (h *queryHandler) queryJoinAppUserExtra(s *sql.Selector) {
	t := sql.Table(entextra.Table)
	s.LeftJoin(t).
		On(
			s.C(entuser.FieldID),
			t.C(entextra.FieldUserID),
		).
		AppendSelect(
			sql.As(t.C(entextra.FieldUsername), "username"),
			sql.As(t.C(entextra.FieldFirstName), "first_name"),
			sql.As(t.C(entextra.FieldLastName), "last_name"),
			sql.As(t.C(entextra.FieldAddressFields), "address_fields"),
			sql.As(t.C(entextra.FieldGender), "gender"),
			sql.As(t.C(entextra.FieldPostalCode), "postal_code"),
			sql.As(t.C(entextra.FieldAge), "age"),
			sql.As(t.C(entextra.FieldBirthday), "birthday"),
			sql.As(t.C(entextra.FieldAvatar), "avatar"),
			sql.As(t.C(entextra.FieldOrganization), "organization"),
			sql.As(t.C(entextra.FieldIDNumber), "id_number"),
			sql.As(t.C(entextra.FieldActionCredits), "action_credits"),
		)
}

func (h *queryHandler) queryJoinAppUserControl(s *sql.Selector) {
	t := sql.Table(entappusercontrol.Table)
	s.LeftJoin(t).
		On(
			s.C(entuser.FieldID),
			t.C(entappusercontrol.FieldUserID),
		).
		AppendSelect(
			sql.As(t.C(entappusercontrol.FieldGoogleAuthenticationVerified), "google_authentication_verified"),
			t.C(entappusercontrol.FieldSigninVerifyType),
			t.C(entappusercontrol.FieldKol),
			t.C(entappusercontrol.FieldKolConfirmed),
		)
}

func (h *queryHandler) queryJoinApp(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			s.C(entuser.FieldImportFromApp),
			t.C(entapp.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entapp.FieldName), "imported_from_app_name"),
			sql.As(t.C(entapp.FieldLogo), "imported_from_app_logo"),
		)
}

func (h *queryHandler) queryJoinBanAppUser(s *sql.Selector) {
	t := sql.Table(entbanappuser.Table)
	s.LeftJoin(t).
		On(
			s.C(entuser.FieldID),
			t.C(entbanappuser.FieldUserID),
		).
		AppendSelect(
			sql.As(t.C(entbanappuser.FieldID), "ban_app_user_id"),
			sql.As(t.C(entbanappuser.FieldMessage), "ban_message"),
		)
}

func (h *queryHandler) queryJoinKyc(s *sql.Selector) {
	t := sql.Table(entkyc.Table)
	s.LeftJoin(t).
		On(
			s.C(entuser.FieldID),
			t.C(entkyc.FieldUserID),
		).
		AppendSelect(
			sql.As(t.C(entkyc.FieldState), "kyc_state"),
		)
}

func (h *queryHandler) queryJoinAppUserSecret(s *sql.Selector) {
	t := sql.Table(entappusersecret.Table)
	s.LeftJoin(t).
		On(
			s.C(entuser.FieldID),
			t.C(entappusersecret.FieldUserID),
		).
		AppendSelect(
			sql.As(t.C(entappusersecret.FieldGoogleSecret), "google_secret"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinAppUserExtra(s)
		h.queryJoinAppUserControl(s)
		h.queryJoinApp(s)
		h.queryJoinBanAppUser(s)
		h.queryJoinKyc(s)
		h.queryJoinAppUserSecret(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) queryUserRoles(ctx context.Context) error {
	if len(h.infos) == 0 {
		return nil
	}

	type role struct {
		UserID   uuid.UUID `json:"user_id"`
		RoleName string    `json:"role_name"`
	}

	roles := []*role{}

	uids := []uuid.UUID{}
	for _, info := range h.infos {
		uids = append(uids, uuid.MustParse(info.ID))
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return cli.
			AppRoleUser.
			Query().
			Where(
				entapproleuser.AppID(uuid.MustParse(h.AppID)),
				entapproleuser.UserIDIn(uids...),
			).
			Select(
				entapproleuser.FieldUserID,
			).
			Modify(func(s *sql.Selector) {
				t := sql.Table(entapprole.Table)
				s.LeftJoin(t).
					On(
						s.C(entapproleuser.FieldRoleID),
						t.C(entapprole.FieldID),
					).
					AppendSelect(
						sql.As(t.C(entapprole.FieldRole), "role_name"),
					)
			}).
			Scan(_ctx, &roles)
	})
	if err != nil {
		return err
	}

	for _, role := range roles {
		for _, info := range h.infos {
			if info.ID == role.UserID.String() {
				info.Roles = append(info.Roles, role.RoleName)
			}
		}
	}

	return nil
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.HasGoogleSecret = info.GoogleSecret != ""
	}
}

func (h *Handler) GetUser(ctx context.Context) (info *usermwpb.User, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppUser(cli); err != nil {
			return err
		}
		handler.queryJoin()
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, fmt.Errorf("invalid user")
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	if err := handler.queryUserRoles(ctx); err != nil {
		return nil, err
	}

	handler.formalize()

	return handler.infos[0], nil
}

func GetUser(ctx context.Context, appID, userID string) (*usermwpb.User, error) {
	var infos []*usermwpb.User
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetUser")
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

	span = commontracer.TraceInvoker(span, "user", "method", "Expand")

	infos, err = expand(ctx, []string{userID}, infos)
	if err != nil {
		return nil, err
	}

	return infos[0], nil
}

func GetUsers(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*usermwpb.User, int, error) {
	var infos []*usermwpb.User
	var err error
	var total int

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "user", "db", "CRUD")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppUser.
			Query()
		if conds != nil {
			if conds.ID != nil {
				stm.Where(
					entuser.ID(uuid.MustParse(conds.GetID().GetValue())),
				)
			}
			if conds.AppID != nil {
				stm.Where(
					entuser.AppID(uuid.MustParse(conds.GetAppID().GetValue())),
				)
			}
		}
		total, err = stm.Count(ctx)
		if err != nil {
			logger.Sugar().Errorw("GetUsers", "err", err.Error())
			return err
		}

		stm.
			Offset(int(offset)).
			Limit(int(limit))

		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetUsers", "err", err.Error())
		return nil, 0, err
	}

	users := []string{}
	for _, info := range infos {
		users = append(users, info.ID)
	}

	span = commontracer.TraceInvoker(span, "user", "method", "Expand")

	infos, err = expand(ctx, users, infos)
	if err != nil {
		return nil, total, err
	}

	return infos, total, nil
}

func GetManyUsers(ctx context.Context, userIDs []string) ([]*usermwpb.User, uint32, error) {
	var infos []*usermwpb.User
	var err error
	var total int

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetManyUsers")
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

	span = commontracer.TraceInvoker(span, "user", "db", "CRUD")

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		stm := cli.
			AppUser.
			Query().
			Where(
				entuser.IDIn(users...),
			)
		total, err = stm.Count(ctx)
		if err != nil {
			logger.Sugar().Errorw("GetUsers", "err", err.Error())
			return err
		}

		return join(stm).
			Scan(ctx, &infos)
	})
	if err != nil {
		logger.Sugar().Errorw("GetManyUsers", "err", err.Error())
		return nil, 0, err
	}

	for _, info := range infos {
		info.Banned = info.BanAppUserID != uuid.UUID{}.String()
	}

	span = commontracer.TraceInvoker(span, "user", "method", "Expand")

	infos, err = expand(ctx, userIDs, infos)
	if err != nil {
		return nil, 0, err
	}

	return infos, uint32(total), nil
}

func expand(ctx context.Context, userIDs []string, users []*usermwpb.User) ([]*usermwpb.User, error) {
	type extra struct {
		UserID       uuid.UUID `json:"user_id"`
		GoogleSecret string    `json:"google_secret"`
		RoleName     string    `json:"role_name"`
	}

	var infos []*extra
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "expand")
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

	span = commontracer.TraceInvoker(span, "user", "db", "CRUD")

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

	for _, user := range users {
		if credits, err := decimal.NewFromString(user.ActionCredits); err == nil {
			user.ActionCredits = credits.String()
			continue
		}
		user.ActionCredits = decimal.NewFromInt(0).String()
	}

	for _, info := range infos {
		for _, user := range users {
			if info.UserID.String() == user.ID {
				user.HasGoogleSecret = info.GoogleSecret != ""
				user.GoogleSecret = info.GoogleSecret
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
			entuser.FieldAppID,
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
					sql.As(t1.C(entextra.FieldActionCredits), "action_credits"),
				)

			t2 := sql.Table(entappusercontrol.Table)
			s.
				LeftJoin(t2).
				On(
					s.C(entuser.FieldID),
					t2.C(entappusercontrol.FieldUserID),
				).
				AppendSelect(
					sql.As(t2.C(entappusercontrol.FieldGoogleAuthenticationVerified), "google_authentication_verified"),
					t2.C(entappusercontrol.FieldSigninVerifyType),
					t2.C(entappusercontrol.FieldKol),
					t2.C(entappusercontrol.FieldKolConfirmed),
				)

			t3 := sql.Table(entapp.Table)
			s.
				LeftJoin(t3).
				On(
					s.C(entuser.FieldImportFromApp),
					t3.C(entapp.FieldID),
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
					t4.C(entbanappuser.FieldUserID),
				).
				AppendSelect(
					sql.As(t4.C(entbanappuser.FieldID), "ban_app_user_id"),
					sql.As(t4.C(entbanappuser.FieldMessage), "ban_message"),
				)
			t5 := sql.Table(entkyc.Table)
			s.
				LeftJoin(t5).
				On(
					s.C(entuser.FieldID),
					t5.C(entkyc.FieldUserID),
				).
				AppendSelect(
					sql.As(t5.C(entkyc.FieldState), "kyc_state"),
				)
		})
}
