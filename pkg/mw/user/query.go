package user

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	entapp "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/app"
	entapprole "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approle"
	entapproleuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approleuser"
	entappuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuser"
	entappusercontrol "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appusercontrol"
	entextra "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuserextra"
	entappusersecret "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appusersecret"
	entappuserthirdparty "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuserthirdparty"
	entbanappuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/banappuser"
	entkyc "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/kyc"

	usercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stm            *ent.AppUserSelect
	infos          []*npool.User
	total          uint32
	joinThirdParty bool
}

func (h *queryHandler) selectAppUser(stm *ent.AppUserQuery) {
	h.stm = stm.Select(
		entappuser.FieldID,
		entappuser.FieldAppID,
		entappuser.FieldEmailAddress,
		entappuser.FieldPhoneNo,
		entappuser.FieldImportFromApp,
		entappuser.FieldCreatedAt,
	)
}

func (h *queryHandler) queryAppUser(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid userid")
	}

	h.selectAppUser(
		cli.AppUser.
			Query().
			Where(
				entappuser.ID(*h.ID),
				entappuser.AppID(h.AppID),
				entappuser.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryAppUserByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := usercrud.SetQueryConds(cli.AppUser.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectAppUser(stm)
	return nil
}

func (h *queryHandler) queryJoinAppUserExtra(s *sql.Selector) {
	t := sql.Table(entextra.Table)
	s.LeftJoin(t).
		On(
			s.C(entappuser.FieldID),
			t.C(entextra.FieldUserID),
		).
		On(
			s.C(entappuser.FieldAppID),
			t.C(entextra.FieldAppID),
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
			s.C(entappuser.FieldID),
			t.C(entappusercontrol.FieldUserID),
		).
		On(
			s.C(entappuser.FieldAppID),
			t.C(entappusercontrol.FieldAppID),
		).
		AppendSelect(
			sql.As(t.C(entappusercontrol.FieldGoogleAuthenticationVerified), "google_authentication_verified"),
			t.C(entappusercontrol.FieldSigninVerifyType),
			t.C(entappusercontrol.FieldKol),
			t.C(entappusercontrol.FieldKolConfirmed),
			t.C(entappusercontrol.FieldSelectedLangID),
		)
}

func (h *queryHandler) queryJoinApp(s *sql.Selector) {
	t := sql.Table(entapp.Table)
	s.LeftJoin(t).
		On(
			s.C(entappuser.FieldImportFromApp),
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
			s.C(entappuser.FieldID),
			t.C(entbanappuser.FieldUserID),
		).
		On(
			s.C(entappuser.FieldAppID),
			t.C(entbanappuser.FieldAppID),
		).
		On(
			s.C(entappuser.FieldDeletedAt),
			t.C(entbanappuser.FieldDeletedAt),
		).
		AppendSelect(
			sql.As(t.C(entbanappuser.FieldUserID), "ban_app_user_id"),
			sql.As(t.C(entbanappuser.FieldMessage), "ban_message"),
			sql.As(t.C(entbanappuser.FieldDeletedAt), "ban_deleted_at"),
		)
}

func (h *queryHandler) queryJoinKyc(s *sql.Selector) {
	t := sql.Table(entkyc.Table)
	s.LeftJoin(t).
		On(
			s.C(entappuser.FieldID),
			t.C(entkyc.FieldUserID),
		).
		On(
			s.C(entappuser.FieldAppID),
			t.C(entkyc.FieldAppID),
		).
		OnP(
			sql.EQ(t.C(entkyc.FieldDeletedAt), 0),
		).
		AppendSelect(
			sql.As(t.C(entkyc.FieldState), "kyc_state"),
		)
}

func (h *queryHandler) queryJoinAppUserSecret(s *sql.Selector) {
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
			sql.As(t.C(entappusersecret.FieldGoogleSecret), "google_secret"),
		)
}

func (h *queryHandler) queryJoinAppUserThirdParty(s *sql.Selector) error {
	if !h.joinThirdParty {
		return nil
	}

	t := sql.Table(entappuserthirdparty.Table)
	s.LeftJoin(t).
		On(
			s.C(entappuser.FieldID),
			t.C(entappuserthirdparty.FieldUserID),
		).
		On(
			s.C(entappuser.FieldAppID),
			t.C(entappuserthirdparty.FieldAppID),
		).
		On(
			s.C(entappuser.FieldDeletedAt),
			t.C(entappuserthirdparty.FieldDeletedAt),
		).
		AppendSelect(
			sql.As(t.C(entappuserthirdparty.FieldThirdPartyID), "third_party_id"),
			sql.As(t.C(entappuserthirdparty.FieldThirdPartyUserID), "third_party_user_id"),
			sql.As(t.C(entappuserthirdparty.FieldThirdPartyUsername), "third_party_username"),
			sql.As(t.C(entappuserthirdparty.FieldThirdPartyAvatar), "third_party_avatar"),
		)

	if h.Conds != nil && h.Conds.ThirdPartyID != nil {
		thirdPartyID, ok := h.Conds.ThirdPartyID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid oauth thirdpartyid")
		}
		s.Where(
			sql.EQ(t.C(entappuserthirdparty.FieldThirdPartyID), thirdPartyID),
		)
	}

	if h.Conds != nil && h.Conds.ThirdPartyUserID != nil {
		thirdPartyUserID, ok := h.Conds.ThirdPartyUserID.Val.(string)
		if !ok {
			return fmt.Errorf("invalid oauth thirdpartyuserid")
		}
		s.Where(
			sql.EQ(t.C(entappuserthirdparty.FieldThirdPartyUserID), thirdPartyUserID),
		)
	}

	return nil
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinAppUserExtra(s)
		h.queryJoinAppUserControl(s)
		h.queryJoinApp(s)
		h.queryJoinBanAppUser(s)
		h.queryJoinKyc(s)
		h.queryJoinAppUserSecret(s)
		err = h.queryJoinAppUserThirdParty(s)
	})
	if err != nil {
		return err
	}
	return err
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
				entapproleuser.UserIDIn(uids...),
				entapproleuser.DeletedAt(0),
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
		credits, err := decimal.NewFromString(info.ActionCredits)
		if err != nil {
			info.ActionCredits = decimal.NewFromInt(0).String()
		} else {
			info.ActionCredits = credits.String()
		}
		info.SigninVerifyType = basetypes.SignMethod(basetypes.SignMethod_value[info.SigninVerifyTypeStr])
		_ = json.Unmarshal([]byte(info.AddressFieldsString), &info.AddressFields)
		info.Banned = info.BanAppUserID != "" && info.BanDeletedAt == 0
		info.State = basetypes.KycState(basetypes.KycState_value[info.KycStateStr])
		if info.SelectedLangID != nil {
			if *info.SelectedLangID == uuid.Nil.String() {
				info.SelectedLangID = nil
			} else if _, err := uuid.Parse(*info.SelectedLangID); err != nil {
				info.SelectedLangID = nil
			}
		}
	}
}

func (h *Handler) GetUser(ctx context.Context) (info *npool.User, err error) {
	handler := &queryHandler{
		Handler:        h,
		joinThirdParty: false,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppUser(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
		if err := handler.scan(_ctx); err != nil {
			return err
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
		return nil, fmt.Errorf("too many records: %v", handler.infos)
	}

	if err := handler.queryUserRoles(ctx); err != nil {
		return nil, err
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetUsers(ctx context.Context) ([]*npool.User, uint32, error) {
	if h.AppID == uuid.Nil {
		return nil, 0, fmt.Errorf("invalid appid")
	}

	handler := &queryHandler{
		Handler:        h,
		joinThirdParty: true,
	}

	if h.Conds != nil && h.Conds.ID != nil {
		handler.joinThirdParty = false
	}
	if h.Conds != nil && h.Conds.IDs != nil {
		handler.joinThirdParty = false
	}
	if h.Conds != nil && h.PhoneNO != nil {
		handler.joinThirdParty = false
	}
	if h.Conds != nil && h.EmailAddress != nil {
		handler.joinThirdParty = false
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryAppUserByConds(ctx, cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))

		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, 0, err
	}

	if err := handler.queryUserRoles(ctx); err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
