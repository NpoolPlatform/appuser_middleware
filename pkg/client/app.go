//nolint:dupl
package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermiddleware/app"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
)

func doApp(ctx context.Context, fn func(_ctx context.Context, cli npool.AppUserMiddlewareAppClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get app connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewAppUserMiddlewareAppClient(conn)

	return fn(_ctx, cli)
}

func GetAppInfo(ctx context.Context, id string) (*npool.AppInfo, error) {
	info, err := doApp(ctx, func(_ctx context.Context, cli npool.AppUserMiddlewareAppClient) (cruder.Any, error) {
		resp, err := cli.GetAppInfo(ctx, &npool.GetAppInfoRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app info : %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app info : %v", err)
	}
	return info.(*npool.AppInfo), nil
}
