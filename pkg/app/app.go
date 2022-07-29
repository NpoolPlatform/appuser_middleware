package app

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	entapp "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	ctrl "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appcontrol"
	banapp "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banapp"

	"github.com/google/uuid"
)

func GetApp(ctx context.Context, id string) (*App, error) {
	var err error
	infos := []*App{}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		return cli.
			App.
			Query().
			Select(
				entapp.FieldID,
				entapp.FieldLogo,
				entapp.FieldCreatedBy,
				entapp.FieldCreatedAt,
				entapp.FieldDescription,
			).
			Where(
				entapp.ID(uuid.MustParse(id)),
			).
			Limit(1).
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(banapp.Table)
				s.
					LeftJoin(t1).
					On(
						s.C(entapp.FieldID),
						t1.C(banapp.FieldAppID),
					).
					AppendSelect(
						sql.As(t1.C(banapp.FieldID), "ban_app_id"),
						sql.As(t1.C(banapp.FieldMessage), "ban_message"),
					)

				t2 := sql.Table(ctrl.Table)
				s.
					LeftJoin(t2).
					On(
						s.C(entapp.FieldID),
						t2.C(ctrl.FieldAppID),
					).
					AppendSelect(
						t2.C(ctrl.FieldSignupMethods),
						t2.C(ctrl.FieldExternSigninMethods),
						t2.C(ctrl.FieldRecaptchaMethod),
						t2.C(ctrl.FieldKycEnable),
						t2.C(ctrl.FieldSigninVerifyEnable),
						t2.C(ctrl.FieldInvitationCodeMust),
					)
			}).
			Scan(ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	return infos[0], nil
}
