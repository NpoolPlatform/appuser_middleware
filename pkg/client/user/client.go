//nolint:nolintlint,dupl
package user

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

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

func CreateUser(ctx context.Context, in *npool.UserReq) (*npool.User, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateUser(ctx, &npool.CreateUserRequest{
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
	return info.(*npool.User), nil
}

func UpdateUser(ctx context.Context, in *npool.UserReq) (*npool.User, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateUser(ctx, &npool.UpdateUserRequest{
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
	return info.(*npool.User), nil
}

func GetUser(ctx context.Context, appID, userID string) (*npool.User, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetUser(ctx, &npool.GetUserRequest{
			AppID:  appID,
			UserID: userID,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.User), nil
}

func GetUsers(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.User, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetUsers(ctx, &npool.GetUsersRequest{
			Conds:  conds,
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
	return infos.([]*npool.User), total, nil
}

func VerifyAccount(
	ctx context.Context,
	appID, account string,
	accountType basetypes.SignMethod,
	passwordHash string,
) (
	*npool.User, error,
) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.VerifyAccount(ctx, &npool.VerifyAccountRequest{
			AppID:        appID,
			Account:      account,
			AccountType:  accountType,
			PasswordHash: passwordHash,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.User), nil
}

func VerifyUser(
	ctx context.Context,
	appID, userID string,
	passwordHash string,
) (
	*npool.User, error,
) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.VerifyUser(ctx, &npool.VerifyUserRequest{
			AppID:        appID,
			UserID:       userID,
			PasswordHash: passwordHash,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.User), nil
}

func DeleteUser(ctx context.Context, appID, userID string) (*npool.User, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteUser(ctx, &npool.DeleteUserRequest{
			Info: &npool.UserReq{
				ID:    &userID,
				AppID: &appID,
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
	return info.(*npool.User), nil
}
