package recoverycode

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/recoverycode"

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

func GetRecoveryCodeOnly(ctx context.Context, conds *npool.Conds) (*npool.RecoveryCode, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetRecoveryCodes(ctx, &npool.GetRecoveryCodesRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  2,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.RecoveryCode)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.RecoveryCode)) > 1 {
		return nil, fmt.Errorf("too many record")
	}
	return infos.([]*npool.RecoveryCode)[0], nil
}

func UpdateRecoveryCode(ctx context.Context, in *npool.RecoveryCodeReq) (*npool.RecoveryCode, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateRecoveryCode(ctx, &npool.UpdateRecoveryCodeRequest{
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
	return info.(*npool.RecoveryCode), nil
}

func GenerateRecoveryCodes(ctx context.Context, in *npool.RecoveryCodeReq) ([]*npool.RecoveryCode, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GenerateRecoveryCodes(ctx, &npool.GenerateRecoveryCodesRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.RecoveryCode), nil
}

func GetRecoveryCodes(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.RecoveryCode, uint32, error) {
	var total uint32
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetRecoveryCodes(ctx, &npool.GetRecoveryCodesRequest{
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
	return infos.([]*npool.RecoveryCode), total, nil
}

func ExistRecoveryCodeConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistRecoveryCodeConds(ctx, &npool.ExistRecoveryCodeCondsRequest{
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
