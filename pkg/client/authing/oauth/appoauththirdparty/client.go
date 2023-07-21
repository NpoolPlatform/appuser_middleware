package appoauththirdparty

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/oauth/appoauththirdparty"

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

func CreateOAuthThirdParty(ctx context.Context, req *npool.OAuthThirdPartyReq) (*npool.OAuthThirdParty, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateOAuthThirdParty(ctx, &npool.CreateOAuthThirdPartyRequest{
			Info: req,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.OAuthThirdParty), nil
}

func UpdateOAuthThirdParty(ctx context.Context, req *npool.OAuthThirdPartyReq) (*npool.OAuthThirdParty, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateOAuthThirdParty(ctx, &npool.UpdateOAuthThirdPartyRequest{
			Info: req,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.OAuthThirdParty), nil
}

func ExistOAuthThirdParty(ctx context.Context, id string) (bool, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistOAuthThirdParty(ctx, &npool.ExistOAuthThirdPartyRequest{
			ID: id,
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

func GetOAuthThirdParty(ctx context.Context, id string) (*npool.OAuthThirdParty, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetOAuthThirdParty(ctx, &npool.GetOAuthThirdPartyRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.OAuthThirdParty), nil
}

func GetOAuthThirdParties(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.OAuthThirdParty, uint32, error) {
	var total uint32

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetOAuthThirdParties(ctx, &npool.GetOAuthThirdPartiesRequest{
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
	return infos.([]*npool.OAuthThirdParty), total, nil
}

func ExistOAuthThirdPartyConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistOAuthThirdPartyConds(ctx, &npool.ExistOAuthThirdPartyCondsRequest{
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

func DeleteOAuthThirdParty(ctx context.Context, id string) (*npool.OAuthThirdParty, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteOAuthThirdParty(ctx, &npool.DeleteOAuthThirdPartyRequest{
			Info: &npool.OAuthThirdPartyReq{
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
	return info.(*npool.OAuthThirdParty), nil
}

func GetOAuthThirdPartyOnly(ctx context.Context, conds *npool.Conds) (*npool.OAuthThirdParty, error) {
	const limit = 2
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetOAuthThirdParties(ctx, &npool.GetOAuthThirdPartiesRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.OAuthThirdParty)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.OAuthThirdParty)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.OAuthThirdParty)[0], nil
}

func GetOAuthThirdPartyDecryptOnly(ctx context.Context, conds *npool.Conds) (*npool.OAuthThirdParty, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetOAuthThirdPartyDecryptOnly(ctx, &npool.GetOAuthThirdPartyDecryptOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}

	return info.(*npool.OAuthThirdParty), nil
}
