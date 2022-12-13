package subscriber

import (
	"context"
	"fmt"

	subscribermgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/subscriber"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	entapp "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	entsubscriber "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/subscriber"

	"github.com/google/uuid"
)

func GetSubscriber(ctx context.Context, id string) (*npool.Subscriber, error) {
	var infos []*npool.Subscriber

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.
			Subscriber.
			Query().
			Where(
				entsubscriber.ID(uuid.MustParse(id)),
			)
		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, fmt.Errorf("no record")
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	return infos[0], nil
}

func GetSubscriberes(ctx context.Context, conds *subscribermgrpb.Conds, offset, limit int32) ([]*npool.Subscriber, uint32, error) {
	return nil, 0, nil
}

func join(stm *ent.SubscriberQuery) *ent.SubscriberSelect {
	return stm.
		Select(
			entsubscriber.FieldID,
			entsubscriber.FieldAppID,
			entsubscriber.FieldEmailAddress,
			entsubscriber.FieldRegistered,
			entsubscriber.FieldCreatedAt,
			entsubscriber.FieldUpdatedAt,
		).
		Modify(func(s *sql.Selector) {
			t1 := sql.Table(entapp.Table)
			s.
				LeftJoin(t1).
				On(
					s.C(entsubscriber.FieldAppID),
					t1.C(entapp.FieldID),
				).
				AppendSelect(
					sql.As(t1.C(entapp.FieldName), "app_name"),
				)
		})
}
