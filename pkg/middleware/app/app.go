package app

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	entapp "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	entappcontrol "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appcontrol"
	entbanapp "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banapp"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
)

type Info struct {
	ID                  string `json:"id"`
	CreatedBy           string `sql:"created_by"`
	Name                string `sql:"name"`
	Logo                string `sql:"logo"`
	CreatedAt           uint32 `sql:"created_at"`
	Description         string `sql:"description"`
	BanAppAppID         string `sql:"ban_app_app_id"`
	IsBanApp            string `sql:"is_ban_app"`
	BanAppMessage       string `sql:"ban_app_message"`
	SignupMethods       string `sql:"signup_methods"`
	ExternSigninMethods string `sql:"extern_signin_methods"`
	RecaptchaMethod     string `sql:"recaptcha_method"`
	KycEnable           int    `sql:"kyc_enable"`
	SigninVerifyEnable  int    `sql:"signin_verify_enable"`
	InvitationCodeMust  int    `sql:"invitation_code_must"`
}

func GetAppInfo(ctx context.Context, id uuid.UUID) (*Info, error) {
	var err error
	var resp []*Info

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppInfo")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.AddEvent("call db query")
	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		err = cli.Debug().App.Query().Select(
			entapp.FieldID,
			entapp.FieldName,
			entapp.FieldLogo,
			entapp.FieldCreatedBy,
			entapp.FieldCreatedAt,
			entapp.FieldDescription,
		).Where(
			entapp.ID(id),
		).Limit(1).
			Modify(func(s *sql.Selector) {
				b := sql.Table(entbanapp.Table)
				c := sql.Table(entappcontrol.Table)
				s.LeftJoin(b).
					On(
						s.C(entapp.FieldID),
						b.C(entbanapp.FieldAppID),
					).AppendSelect(
					sql.As(b.C(entbanapp.FieldAppID), "ban_app_app_id"),
					sql.As(b.C(entbanapp.FieldMessage), "ban_app_message"),
				)
				s.LeftJoin(c).On(
					s.C(entapp.FieldID),
					c.C(entappcontrol.FieldAppID),
				).AppendSelect(
					c.C(entappcontrol.FieldSignupMethods),
					c.C(entappcontrol.FieldExternSigninMethods),
					c.C(entappcontrol.FieldRecaptchaMethod),
					c.C(entappcontrol.FieldKycEnable),
					c.C(entappcontrol.FieldSigninVerifyEnable),
					c.C(entappcontrol.FieldInvitationCodeMust),
				)
			}).Scan(ctx, &resp)
		return err
	})
	if err != nil {
		return nil, err
	}
	if len(resp) == 0 {
		return nil, fmt.Errorf("not found app")
	}
	return resp[0], nil
}

func GetAppInfos(ctx context.Context, offset int32, limit int32) ([]*Info, error) {
	var err error
	var resp []*Info

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppInfo")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.AddEvent("call db query")
	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		err = cli.Debug().App.Query().Select(
			entapp.FieldID,
			entapp.FieldName,
			entapp.FieldLogo,
			entapp.FieldCreatedBy,
			entapp.FieldCreatedAt,
			entapp.FieldDescription,
		).
			Limit(int(limit)).
			Offset(int(offset)).
			Modify(func(s *sql.Selector) {
				b := sql.Table(entbanapp.Table)
				c := sql.Table(entappcontrol.Table)
				s.LeftJoin(b).
					On(
						s.C(entapp.FieldID),
						b.C(entbanapp.FieldAppID),
					).AppendSelect(
					sql.As(b.C(entbanapp.FieldAppID), "ban_app_app_id"),
					sql.As(b.C(entbanapp.FieldMessage), "ban_app_message"),
				)
				s.LeftJoin(c).On(
					s.C(entapp.FieldID),
					c.C(entappcontrol.FieldAppID),
				).AppendSelect(
					c.C(entappcontrol.FieldSignupMethods),
					c.C(entappcontrol.FieldExternSigninMethods),
					c.C(entappcontrol.FieldRecaptchaMethod),
					c.C(entappcontrol.FieldKycEnable),
					c.C(entappcontrol.FieldSigninVerifyEnable),
					c.C(entappcontrol.FieldInvitationCodeMust),
				)
			}).Scan(ctx, &resp)
		return err
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetAppInfosByCreator(ctx context.Context, creatorID uuid.UUID, offset int32, limit int32) ([]*Info, error) {
	var err error
	var resp []*Info

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppInfo")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.AddEvent("call db query")
	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		err = cli.Debug().App.Query().Select(
			entapp.FieldID,
			entapp.FieldName,
			entapp.FieldLogo,
			entapp.FieldCreatedBy,
			entapp.FieldCreatedAt,
			entapp.FieldDescription,
		).Where(
			entapp.CreatedBy(creatorID),
		).
			Limit(int(limit)).
			Offset(int(offset)).
			Modify(func(s *sql.Selector) {
				b := sql.Table(entbanapp.Table)
				c := sql.Table(entappcontrol.Table)
				s.LeftJoin(b).
					On(
						s.C(entapp.FieldID),
						b.C(entbanapp.FieldAppID),
					).AppendSelect(
					sql.As(b.C(entbanapp.FieldAppID), "ban_app_app_id"),
					sql.As(b.C(entbanapp.FieldMessage), "ban_app_message"),
				)
				s.LeftJoin(c).On(
					s.C(entapp.FieldID),
					c.C(entappcontrol.FieldAppID),
				).AppendSelect(
					c.C(entappcontrol.FieldSignupMethods),
					c.C(entappcontrol.FieldExternSigninMethods),
					c.C(entappcontrol.FieldRecaptchaMethod),
					c.C(entappcontrol.FieldKycEnable),
					c.C(entappcontrol.FieldSigninVerifyEnable),
					c.C(entappcontrol.FieldInvitationCodeMust),
				)
			}).Scan(ctx, &resp)
		return err
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
