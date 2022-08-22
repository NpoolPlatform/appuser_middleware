//nolint:dupl
package authing

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing"

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

func ExistAuth(ctx context.Context, appID, userID, resource, method string) (bool, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistAuth(ctx, &npool.ExistAuthRequest{
			AppID:    appID,
			UserID:   &userID,
			Resource: resource,
			Method:   method,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return infos.(bool), nil
}

func GetAuths(ctx context.Context, appID string, offset, limit int32) ([]*npool.Auth, uint32, error) {
	var total uint32

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAuths(ctx, &npool.GetAuthsRequest{
			AppID:  appID,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.Auth), total, nil
}

func GetHistories(ctx context.Context, appID string, offset, limit int32) ([]*npool.History, uint32, error) {
	var total uint32

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetHistories(ctx, &npool.GetHistoriesRequest{
			AppID:  appID,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.History), total, nil
}
