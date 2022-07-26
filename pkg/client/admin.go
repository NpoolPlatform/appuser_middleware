//nolint:dupl
package client

import (
	"context"
	"fmt"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/appuser"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/appusersecret"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appusermw/admin"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
)

func doAdmin(ctx context.Context, fn func(_ctx context.Context, cli npool.AppUserMiddlewareAdminClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get app connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewAppUserMiddlewareAdminClient(conn)

	return fn(_ctx, cli)
}

func CreateGenesisRoleUser(ctx context.Context, user *appuser.AppUser, secret *appusersecret.AppUserSecret) (*npool.CreateGenesisRoleUserResponse, error) {
	info, err := doAdmin(ctx, func(_ctx context.Context, cli npool.AppUserMiddlewareAdminClient) (cruder.Any, error) {
		resp, err := cli.CreateGenesisRoleUser(ctx, &npool.CreateGenesisRoleUserRequest{
			User:   user,
			Secret: secret,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create genesis role user : %v", err)
		}
		return resp, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create genesis role user : %v", err)
	}
	return info.(*npool.CreateGenesisRoleUserResponse), nil
}
