package authing

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/auth"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banapp"

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
		logger.Sugar().Infow("existAppAuth", "Reason", "banned")
		return false, nil
	}
	if res[0].AppID != res[0].AppVID {
		logger.Sugar().Infow("existAppAuth", "Reason", "mismatch appid")
		return false, nil
	}

	return true, nil
}

func existUserAuth(ctx context.Context, appID, userID, resource, method string) (exist bool, err error) {
	return false, nil
}

func ExistAuth(ctx context.Context, appID string, userID *string, resource, method string) (exist bool, err error) {
	if userID == nil {
		return existAppAuth(ctx, appID, resource, method)
	}
	return existUserAuth(ctx, appID, *userID, resource, method)
}
