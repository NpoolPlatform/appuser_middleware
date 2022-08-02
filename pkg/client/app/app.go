//nolint:nolintlint,dupl
package app

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.AppMwClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second) //nolint
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := npool.NewAppMwClient(conn)

	return fn(_ctx, cli)
}

func CreateApp(ctx context.Context, in *npool.AppReq) (*npool.App, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.AppMwClient) (cruder.Any, error) {
		resp, err := cli.CreateApp(ctx, &npool.CreateAppRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.App), nil
}

func UpdateApp(ctx context.Context, in *npool.AppReq) (*npool.App, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.AppMwClient) (cruder.Any, error) {
		resp, err := cli.UpdateApp(ctx, &npool.UpdateAppRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.App), nil
}

func GetApp(ctx context.Context, appID, userID string) (*npool.App, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.AppMwClient) (cruder.Any, error) {
		resp, err := cli.GetApp(ctx, &npool.GetAppRequest{
			AppID: appID,
		})
		if err != nil {
			return nil, err
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.App), nil
}

func GetApps(ctx context.Context, limit, offset int32) ([]*npool.App, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.AppMwClient) (cruder.Any, error) {
		resp, err := cli.GetApps(ctx, &npool.GetAppsRequest{
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, err
	}
	return info.([]*npool.App), nil
}

func GetUserApps(ctx context.Context, userID string, offset, limit int32) ([]*npool.App, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.AppMwClient) (cruder.Any, error) {
		resp, err := cli.GetUserApps(ctx, &npool.GetUserAppsRequest{
			UserID: userID,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.App), total, nil
}

func GetSignMethods(ctx context.Context) ([]*npool.App, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.AppMwClient) (cruder.Any, error) {
		resp, err := cli.GetSignMethods(ctx, &npool.GetSignMethodsRequest{})
		if err != nil {
			return nil, err
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.App), total, nil
}

func GetRecaptchas(ctx context.Context) ([]*npool.App, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.AppMwClient) (cruder.Any, error) {
		resp, err := cli.GetRecaptchas(ctx, &npool.GetRecaptchasRequest{})
		if err != nil {
			return nil, err
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.App), total, nil
}
