//nolint:nolintlint,dupl
package appsubscribe

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber/app/subscribe"

	servicename "github.com/NpoolPlatform/appuser-middleware/pkg/servicename"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second) //nolint
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return fn(_ctx, cli)
}

func CreateAppSubscribe(ctx context.Context, in *npool.AppSubscribeReq) (*npool.AppSubscribe, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateAppSubscribe(ctx, &npool.CreateAppSubscribeRequest{
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
	return info.(*npool.AppSubscribe), nil
}

func GetAppSubscribe(ctx context.Context, appID string) (*npool.AppSubscribe, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAppSubscribe(ctx, &npool.GetAppSubscribeRequest{
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
	return info.(*npool.AppSubscribe), nil
}

func GetAppSubscribes(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.AppSubscribe, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAppSubscribes(ctx, &npool.GetAppSubscribesRequest{
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
	return infos.([]*npool.AppSubscribe), total, nil
}

func GetAppSubscribeOnly(ctx context.Context, conds *npool.Conds) (*npool.AppSubscribe, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAppSubscribes(ctx, &npool.GetAppSubscribesRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  2, //nolint
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.AppSubscribe)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.AppSubscribe)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.AppSubscribe)[0], nil
}

func ExistAppSubscribeConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistAppSubscribeConds(ctx, &npool.ExistAppSubscribeCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func DeleteAppSubscribe(ctx context.Context, id string) (*npool.AppSubscribe, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteAppSubscribe(ctx, &npool.DeleteAppSubscribeRequest{
			Info: &npool.AppSubscribeReq{
				ID: &id,
			},
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.AppSubscribe), nil
}
