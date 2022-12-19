//nolint:nolintlint,dupl
package subscriber

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	mgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/subscriber"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second) //nolint
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return fn(_ctx, cli)
}

func CreateSubscriber(ctx context.Context, in *mgrpb.SubscriberReq) (*npool.Subscriber, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateSubscriber(ctx, &npool.CreateSubscriberRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		fmt.Printf("1--- %v\n", err)
		return nil, err
	}
	return info.(*npool.Subscriber), nil
}

func UpdateSubscriber(ctx context.Context, in *mgrpb.SubscriberReq) (*npool.Subscriber, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateSubscriber(ctx, &npool.UpdateSubscriberRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Subscriber), nil
}

func GetSubscriber(ctx context.Context, appID string) (*npool.Subscriber, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetSubscriber(ctx, &npool.GetSubscriberRequest{
			ID: appID,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Subscriber), nil
}

func GetSubscriberes(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*npool.Subscriber, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetSubscriberes(ctx, &npool.GetSubscriberesRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}

		total = resp.Total

		return resp.Infos, nil
	})
	if err != nil {
		return nil, total, err
	}
	return infos.([]*npool.Subscriber), total, nil
}

func DeleteSubscriber(ctx context.Context, id string) (*npool.Subscriber, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteSubscriber(ctx, &npool.DeleteSubscriberRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Subscriber), nil
}
