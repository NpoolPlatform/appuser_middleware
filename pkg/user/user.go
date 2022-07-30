package user

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	// entapprole "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approle"
	// entapproleuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approleuser"

	entapp "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	entuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"
	entappusercontrol "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusercontrol"
	entextra "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"

	// entsecret "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"
	// entbanappuser "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banappuser"

	"github.com/google/uuid"
)

func GetUser(ctx context.Context, appID, userID string) (*User, error) {
	var err error
	var infos []*User

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
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	infos, err = expand(ctx, []string{userID}, infos)
	if err != nil {
		return nil, err
	}

	return infos[0], nil
}

func GetUsers(ctx context.Context, appID string, offset, limit int32) ([]*User, error) {
	var err error
	var infos []*User

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
		return nil, err
	}

	users := []string{}
	for _, info := range infos {
		users = append(users, info.ID)
	}

	infos, err = expand(ctx, users, infos)
	if err != nil {
		return nil, err
	}

	return infos, nil
}

func GetManyUsers(ctx context.Context, userIDs []string) ([]*User, error) {
	var err error
	var infos []*User

	users := []uuid.UUID{}
	for _, user := range userIDs {
		users = append(users, uuid.MustParse(user))
	}

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
		return nil, err
	}

	infos, err = expand(ctx, userIDs, infos)
	if err != nil {
		return nil, err
	}

	return infos, nil
}

func expand(ctx context.Context, userIDs []string, users []*User) ([]*User, error) {
	// TODO: fill secret map, control, role
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
					t2.C(entextra.FieldUserID),
				).
				AppendSelect(
				// TODO: add expression
				)

			t3 := sql.Table(entapp.Table)
			s.
				LeftJoin(t3).
				On(
					s.C(entuser.FieldImportFromApp),
					t2.C(entextra.FieldID),
				).
				AppendSelect(
				// TODO: add expression
				)
		})
}
