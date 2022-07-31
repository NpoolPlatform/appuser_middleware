package user

import (
	"context"
	"time"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.UserMwClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second) //nolint
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := npool.NewUserMwClient(conn)

	return fn(_ctx, cli)
}

func GetManyUsers(ctx context.Context, ids []string) ([]*npool.User, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.UserMwClient) (cruder.Any, error) {
		resp, err := cli.GetManyUsers(ctx, &npool.GetManyUsersRequest{
			IDs: ids,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.User), nil
}
