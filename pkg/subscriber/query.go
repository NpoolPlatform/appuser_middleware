package subscriber

import (
	"context"
	"fmt"

	mgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/subscriber"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"

	"entgo.io/ent/dialect/sql"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	crud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/subscriber"

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

func GetSubscriberes(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*npool.Subscriber, uint32, error) {
	var infos []*npool.Subscriber
	var total uint32

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		_total, err := stm.Count(ctx)
		if err != nil {
			return err
		}
		total = uint32(_total)

		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, 0, err
	}

	return infos, total, nil
}

func GetSubscriberOnly(ctx context.Context, conds *mgrpb.Conds) (*npool.Subscriber, error) {
	var infos []*npool.Subscriber

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := crud.SetQueryConds(conds, cli)
		if err != nil {
			return err
		}
		return join(stm).
			Scan(_ctx, &infos)
	})
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	return infos[0], nil
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
