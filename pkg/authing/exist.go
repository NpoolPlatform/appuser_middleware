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
				auth.Resource(resource),
				auth.Method(method),
				auth.UserID(uuid.UUID{}),
				auth.RoleID(uuid.UUID{}),
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
func existRoleAuth(ctx context.Context, appID, userID, resource, method string) (exist bool, err error) {
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
					On(
						t5.C(auth.FieldRoleID),
						s.C(approleuser.FieldRoleID),
					).
					Where(
						sql.And(
							sql.EQ(s.C(approleuser.FieldAppID), appID),
							sql.EQ(s.C(approleuser.FieldUserID), userID),
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
		logger.Sugar().Infow("existRoleAuth", "Reason", "no record")
		return false, nil
	}
	if res[0].AppBID == res[0].AppID {
		logger.Sugar().Infow("existRoleAuth", "Reason", "banned appid")
		return false, nil
	}
	if res[0].AppID != res[0].AppVID {
		logger.Sugar().Infow("existRoleAuth", "Reason", "mismatch appid")
		return false, nil
	}
	if res[0].UserBID == res[0].UserID {
		logger.Sugar().Infow("existRoleAuth", "Reason", "banned userid")
		return false, nil
	}
	if res[0].UserID != res[0].UserVID {
		logger.Sugar().Infow("existRoleAuth", "Reason", "mismatch userid")
		return false, nil
	}

	return true, nil
}

//nolint
func existUserAuth(ctx context.Context, appID, userID, resource, method string) (exist bool, err error) {
	type r struct {
		AppID   string `sql:"app_id"`
		UserID  string `sql:"id"`
		AppVID  string `sql:"app_vid"`
		AppBID  string `sql:"app_bid"`
		UserBID string `sql:"user_bid"`
	}

	res := []*r{}

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		return cli.
			AppUser.
			Query().
			Select(
				appuser.FieldAppID,
				appuser.FieldID,
			).
			Modify(func(s *sql.Selector) {
				t1 := sql.Table(app.Table)
				s.
					LeftJoin(t1).
					On(
						t1.C(app.FieldID),
						s.C(appuser.FieldAppID),
					).
					AppendSelect(
						sql.As(t1.C(app.FieldID), "app_vid"),
					)

				t2 := sql.Table(banapp.Table)
				s.
					LeftJoin(t2).
					On(
						t2.C(banapp.FieldAppID),
						s.C(appuser.FieldAppID),
					).
					AppendSelect(
						sql.As(t2.C(banapp.FieldAppID), "app_bid"),
					)

				t3 := sql.Table(banappuser.Table)
				s.
					LeftJoin(t3).
					On(
						t3.C(banappuser.FieldAppID),
						s.C(appuser.FieldAppID),
					).
					On(
						t3.C(banappuser.FieldUserID),
						s.C(appuser.FieldID),
					).
					AppendSelect(
						sql.As(t3.C(banappuser.FieldUserID), "user_bid"),
					)

				t4 := sql.Table(auth.Table)
				s.
					LeftJoin(t4).
					On(
						t4.C(auth.FieldAppID),
						s.C(appuser.FieldAppID),
					).
					On(
						t4.C(auth.FieldUserID),
						s.C(appuser.FieldID),
					).
					Where(
						sql.And(
							sql.EQ(s.C(appuser.FieldAppID), appID),
							sql.EQ(s.C(appuser.FieldID), userID),
							sql.EQ(t4.C(auth.FieldResource), resource),
							sql.EQ(t4.C(auth.FieldMethod), method),
						),
					).
					Limit(1)
			}).
			Scan(ctx, &res)
	})
	if err != nil {
		logger.Sugar().Errorw("existUserAuth", "error", err)
		return false, err
	}
	if len(res) == 0 {
		logger.Sugar().Infow("existUserAuth", "Reason", "no record")
		return false, nil
	}
	if res[0].AppBID == res[0].AppID {
		logger.Sugar().Infow("existUserAuth", "Reason", "banned appid")
		return false, nil
	}
	if res[0].AppID != res[0].AppVID {
		logger.Sugar().Infow("existUserAuth", "Reason", "mismatch appid")
		return false, nil
	}
	if res[0].UserBID == res[0].UserID {
		logger.Sugar().Infow("existUserAuth", "Reason", "banned userid")
		return false, nil
	}

	return true, nil
}

func ExistAuth(ctx context.Context, appID string, userID *string, resource, method string) (exist bool, err error) {
	if userID == nil {
		return existAppAuth(ctx, appID, resource, method)
	}
	if exist, err := existRoleAuth(ctx, appID, *userID, resource, method); err == nil && exist {
		return true, nil
	}
	return existUserAuth(ctx, appID, *userID, resource, method)
}
