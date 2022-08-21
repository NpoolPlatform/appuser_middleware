package authing

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approleuser"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/auth"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banapp"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banappuser"

	"entgo.io/ent/dialect/sql"

	"github.com/google/uuid"
)

func existAppAuth(ctx context.Context, appID, resource, method string) (exist bool, err error) {
	type r struct {
		AppID  string `sql:"app_id"`
		AppVID string `sql:"app_vid"`
		AppBID string `sql:"app_bid"`
	}

	res := []*r{}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		return cli.
			Auth.
			Query().
			Where(
				auth.AppID(uuid.MustParse(appID)),
				auth.RoleID(uuid.UUID{}),
				auth.UserID(uuid.UUID{}),
				auth.Resource(resource),
				auth.Method(method),
			).
			Limit(1).
			Select(
				auth.FieldAppID,
			).
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(app.Table)
				s.
					LeftJoin(t1).
					On(
						t1.C(app.FieldID),
						s.C(auth.FieldAppID),
					).
					AppendSelect(
						sql.As(t1.C(app.FieldID), "app_vid"),
					)

				t2 := sql.Table(banapp.Table)
				s.
					LeftJoin(t2).
					On(
						t2.C(banapp.FieldAppID),
						s.C(auth.FieldAppID),
					).
					AppendSelect(
						sql.As(t2.C(banapp.FieldAppID), "app_bid"),
					)
			}).
			Scan(ctx, &res)
	})
	if err != nil {
		return false, err
	}
	if len(res) == 0 {
		logger.Sugar().Infow("existAppAuth", "Reason", "no record")
		return false, nil
	}
	if res[0].AppBID == res[0].AppID {
		logger.Sugar().Infow("existAppAuth", "Reason", "banned appid")
		return false, nil
	}
	if res[0].AppID != res[0].AppVID {
		logger.Sugar().Infow("existAppAuth", "Reason", "mismatch appid")
		return false, nil
	}

	return true, nil
}

// nolint
func existUserAuth(ctx context.Context, appID, userID, resource, method string) (exist bool, err error) {
	type r struct {
		AppID   string `sql:"app_id"`
		RoleID  string `sql:"role_id"`
		UserID  string `sql:"user_id"`
		AppVID  string `sql:"app_vid"`
		AppBID  string `sql:"app_bid"`
		UserVID string `sql:"user_vid"`
		UserBID string `sql:"user_bid"`
	}

	res := []*r{}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		return cli.
			AppRoleUser.
			Query().
			Where(
				approleuser.AppID(uuid.MustParse(appID)),
				approleuser.UserID(uuid.MustParse(userID)),
			).
			Select(
				approleuser.FieldAppID,
				approleuser.FieldRoleID,
				approleuser.FieldUserID,
			).
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(app.Table)
				s.
					LeftJoin(t1).
					On(
						t1.C(app.FieldID),
						s.C(approleuser.FieldAppID),
					).
					AppendSelect(
						sql.As(t1.C(app.FieldID), "app_vid"),
					)

				t2 := sql.Table(banapp.Table)
				s.
					LeftJoin(t2).
					On(
						t2.C(banapp.FieldAppID),
						s.C(approleuser.FieldAppID),
					).
					AppendSelect(
						sql.As(t2.C(banapp.FieldAppID), "app_bid"),
					)

				t3 := sql.Table(appuser.Table)
				s.
					LeftJoin(t3).
					On(
						t3.C(appuser.FieldAppID),
						s.C(approleuser.FieldAppID),
					).
					On(
						t3.C(appuser.FieldID),
						s.C(approleuser.FieldUserID),
					).
					AppendSelect(
						sql.As(t3.C(appuser.FieldID), "user_vid"),
					)

				t4 := sql.Table(banappuser.Table)
				s.
					LeftJoin(t4).
					On(
						t4.C(banappuser.FieldAppID),
						s.C(approleuser.FieldAppID),
					).
					On(
						t4.C(banappuser.FieldUserID),
						s.C(approleuser.FieldUserID),
					).
					AppendSelect(
						sql.As(t4.C(banappuser.FieldUserID), "user_bid"),
					)

				t5 := sql.Table(auth.Table)
				s.
					LeftJoin(t5).
					On(
						t5.C(auth.FieldAppID),
						s.C(approleuser.FieldAppID),
					).
					Where(
						sql.And(
							sql.Or(
								sql.EQ(s.C(approleuser.FieldUserID), userID),
								sql.EQ(s.C(approleuser.FieldRoleID), t5.C(auth.FieldRoleID)),
							),
							sql.EQ(t5.C(auth.FieldResource), resource),
							sql.EQ(t5.C(auth.FieldMethod), method),
						),
					).
					Limit(1)
			}).
			Scan(ctx, &res)
	})
	if err != nil {
		return false, err
	}
	if len(res) == 0 {
		logger.Sugar().Infow("existAppAuth", "Reason", "no record")
		return false, nil
	}
	if res[0].AppBID == res[0].AppID {
		logger.Sugar().Infow("existAppAuth", "Reason", "banned appid")
		return false, nil
	}
	if res[0].AppID != res[0].AppVID {
		logger.Sugar().Infow("existAppAuth", "Reason", "mismatch appid")
		return false, nil
	}
	if res[0].UserBID == res[0].UserID {
		logger.Sugar().Infow("existAppAuth", "Reason", "banned userid")
		return false, nil
	}
	if res[0].UserID != res[0].UserVID {
		logger.Sugar().Infow("existAppAuth", "Reason", "mismatch userid")
		return false, nil
	}

	return true, nil
}

func ExistAuth(ctx context.Context, appID string, userID *string, resource, method string) (exist bool, err error) {
	if userID == nil {
		return existAppAuth(ctx, appID, resource, method)
	}
	return existUserAuth(ctx, appID, *userID, resource, method)
}
