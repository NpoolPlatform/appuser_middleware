//nolint:dupl
package client

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermw/user"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
)

func doUser(ctx context.Context, fn func(_ctx context.Context, cli npool.AppUserMiddlewareUserClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get app connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewAppUserMiddlewareUserClient(conn)

	return fn(_ctx, cli)
}

func GetUserInfo(ctx context.Context, appID, userID string) (*npool.AppUserInfo, error) {
	info, err := doUser(ctx, func(_ctx context.Context, cli npool.AppUserMiddlewareUserClient) (cruder.Any, error) {
		resp, err := cli.GetUserInfo(ctx, &npool.GetUserInfoRequest{
			AppID:  appID,
			UserID: userID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app info : %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app info : %v", err)
	}
	return info.(*npool.AppUserInfo), nil
}

func GetUserInfos(ctx context.Context, appID string, offset, limit int32) ([]*npool.AppUserInfo, error) {
	info, err := doUser(ctx, func(_ctx context.Context, cli npool.AppUserMiddlewareUserClient) (cruder.Any, error) {
		resp, err := cli.GetUserInfos(ctx, &npool.GetUserInfosRequest{
			AppID:  appID,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get app infos : %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app infos : %v", err)
	}
	return info.([]*npool.AppUserInfo), nil
}
