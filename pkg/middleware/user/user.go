package user

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	entapp "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	entuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"
	entappusercontrol "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusercontrol"
	entextra "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"
	entbanappuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banappuser"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/NpoolPlatform/message/npool/appusermw/user"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
)

type Info struct {
	ID                                 string `json:"id"`
	AppID                              string `json:"app_id"`
	EmailAddress                       string `json:"email_address"`
	PhoneNO                            string `json:"phone_no"`
	ImportFromApp                      string `json:"import_from_app"`
	CreateAt                           uint32 `json:"create_at"`
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
	RoleCreatedBy                      string `json:"role_created_by"`
	Role                               string `json:"role"`
	RoleDescription                    string `json:"role_description"`
	RoleDefault                        uint32 `json:"role_default"`
	HasGoogleSecret                    uint32 `json:"has_google_secret"`
}

func GetUserInfo(ctx context.Context, userID uuid.UUID) (*user.AppUserInfo, error) {
	var err error
	var resp Info

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
		err = cli.Debug().AppUser.Query().Select(
			entuser.FieldID,
			entuser.FieldCreatedAt,
			entuser.FieldAppID,
			entuser.FieldEmailAddress,
			entuser.FieldPhoneNo,
			entuser.FieldImportFromApp,
		).
			Modify(func(s *sql.Selector) {
				extra := sql.Table(entextra.Table)
				control := sql.Table(entappusercontrol.Table)
				ban := sql.Table(entbanappuser.Table)
				// roleuser := sql.Table(entroleuser.Table)

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
				)

				s.LeftJoin(control).
					On(
						s.C(entapp.FieldID),
						control.C(entappusercontrol.FieldUserID),
					).AppendSelect(
					control.C(entappusercontrol.FieldSigninVerifyByGoogleAuthentication),
					control.C(entappusercontrol.FieldGoogleAuthenticationVerified),
				)

				s.LeftJoin(ban).
					On(
						s.C(entapp.FieldID),
						ban.C(entbanappuser.FieldUserID),
					).AppendSelect(
					sql.As(ban.C(entbanappuser.FieldUserID), "ban_app_user_id"),
				)
				// 1对多结果。。。。。。。。

				//s.LeftJoin(roleuser).
				//	On(
				//		s.C(entapp.FieldID),
				//		ban.C(entbanappuser.FieldUserID),
				//	).AppendSelect(
				//	sql.As(ban.C(entbanappuser.FieldUserID), "ban_app_user_id"),
				//)

			}).Scan(ctx, &resp)
		return err
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
