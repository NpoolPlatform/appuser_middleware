//nolint:nolintlint,dupl
package admin

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/admin"
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	constant "github.com/NpoolPlatform/appuser-manager/pkg/message/const"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.AdminMwClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second) //nolint
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cli := npool.NewAdminMwClient(conn)

	return fn(_ctx, cli)
}

func CreateGenesisUser(ctx context.Context, appID, userID, role, emailAddress, passwordHash string) (*user.User, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.AdminMwClient) (cruder.Any, error) {
		resp, err := cli.CreateGenesisUser(ctx, &npool.CreateGenesisUserRequest{
			AppID:        &appID,
			UserID:       &userID,
			Role:         &role,
			EmailAddress: &emailAddress,
			PasswordHash: &passwordHash,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*user.User), nil
}
