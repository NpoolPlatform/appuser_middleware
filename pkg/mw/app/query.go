package app

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	entapp "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/app"
	entappctrl "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appcontrol"
	entbanapp "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/banapp"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

type queryHandler struct {
	*Handler
	stm   *ent.AppSelect
	infos []*npool.App
	total uint32
}

func (h *queryHandler) selectApp(stm *ent.AppQuery) {
	h.stm = stm.Select(
		entapp.FieldID,
		entapp.FieldCreatedBy,
		entapp.FieldLogo,
		entapp.FieldName,
		entapp.FieldDescription,
		entapp.FieldCreatedAt,
	)
}

func (h *queryHandler) queryApp(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid appid")
	}

	h.selectApp(
		cli.App.
			Query().
			Where(
				entapp.ID(*h.ID),
				entapp.DeletedAt(0),
			),
	)

	return nil
}

func (h *queryHandler) queryApps(ctx context.Context, cli *ent.Client) (err error) {
	stm := cli.App.Query()

	if len(h.IDs) > 0 {
		stm.Where(
			entapp.IDIn(h.IDs...),
		)
	}

	if h.UserID != nil {
		stm.Where(
			entapp.CreatedBy(*h.UserID),
		)
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)
	h.selectApp(stm)
	return nil
}

func (h *queryHandler) queryJoinAppCtrl(s *sql.Selector) {
	t := sql.Table(entappctrl.Table)
	s.LeftJoin(t).
		On(
			s.C(entapp.FieldID),
			t.C(entappctrl.FieldAppID),
		).
		AppendSelect(
			t.C(entappctrl.FieldSignupMethods),
			t.C(entappctrl.FieldExternSigninMethods),
			t.C(entappctrl.FieldRecaptchaMethod),
			t.C(entappctrl.FieldKycEnable),
			t.C(entappctrl.FieldSigninVerifyEnable),
			t.C(entappctrl.FieldInvitationCodeMust),
			t.C(entappctrl.FieldCreateInvitationCodeWhen),
			t.C(entappctrl.FieldMaxTypedCouponsPerOrder),
			t.C(entappctrl.FieldMaintaining),
			t.C(entappctrl.FieldCommitButtonTargets),
		)
}

func (h *queryHandler) queryJoinBanApp(s *sql.Selector) {
	t := sql.Table(entbanapp.Table)
	s.LeftJoin(t).
		On(
			s.C(entapp.FieldID),
			t.C(entbanapp.FieldAppID),
		).
		AppendSelect(
			sql.As(t.C(entbanapp.FieldID), "ban_app_id"),
			sql.As(t.C(entbanapp.FieldMessage), "ban_message"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinAppCtrl(s)
		h.queryJoinBanApp(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.CreateInvitationCodeWhen =
			basetypes.CreateInvitationCodeWhen(
				basetypes.CreateInvitationCodeWhen_value[info.CreateInvitationCodeWhenStr],
			)
		_ = json.Unmarshal([]byte(info.CommitButtonTargetsStr), &info.CommitButtonTargets)

		methods := []string{}
		_methods := []basetypes.SignMethod{}

		_ = json.Unmarshal([]byte(info.SignupMethodsStr), &methods)
		for _, m := range methods {
			_methods = append(_methods, basetypes.SignMethod(basetypes.SignMethod_value[m]))
		}

		emethods := []string{}
		_emethods := []basetypes.SignMethod{}

		_ = json.Unmarshal([]byte(info.ExtSigninMethodsStr), &emethods)
		for _, m := range emethods {
			_emethods = append(_emethods, basetypes.SignMethod(basetypes.SignMethod_value[m]))
		}

		info.SignupMethods = _methods
		info.ExtSigninMethods = _emethods
		info.RecaptchaMethod = basetypes.RecaptchaMethod(basetypes.RecaptchaMethod_value[info.RecaptchaMethodStr])

		info.Banned = info.BanAppID != ""
	}
}

func (h *Handler) GetApp(ctx context.Context) (*npool.App, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryApp(cli); err != nil {
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
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetApps(ctx context.Context) ([]*npool.App, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryApps(_ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
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

	handler.formalize()

	return handler.infos, handler.total, nil
}
