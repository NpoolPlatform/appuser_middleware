//nolint:nolintlint,dupl
package user

import (
	"context"
	"fmt"
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

func CreateThirdUser(ctx context.Context, in *npool.UserReq) (*npool.User, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateThirdUser(ctx, &npool.CreateThirdUserRequest{
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

func GetThirdUserOnly(ctx context.Context, conds *npool.Conds) (*npool.User, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetThirdUsers(ctx, &npool.GetThirdUsersRequest{
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
	if len(infos.([]*npool.User)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.User)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.User)[0], nil
}

func GetUserOnly(ctx context.Context, conds *npool.Conds) (*npool.User, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetUsers(ctx, &npool.GetUsersRequest{
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
	if len(infos.([]*npool.User)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.User)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.User)[0], nil
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

func ExistUser(ctx context.Context, appID, userID string) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistUser(ctx, &npool.ExistUserRequest{
			AppID:  appID,
			UserID: userID,
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

func ExistUserConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistUserConds(ctx, &npool.ExistUserCondsRequest{
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

func DeleteAppUserThirdParty(ctx context.Context, appID, userID, thirdPartyID, thirdPartyUserID string) (*npool.User, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteThirdUser(ctx, &npool.DeleteThirdUserRequest{
			Info: &npool.UserReq{
				ID:               &userID,
				AppID:            &appID,
				ThirdPartyID:     &thirdPartyID,
				ThirdPartyUserID: &thirdPartyUserID,
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
