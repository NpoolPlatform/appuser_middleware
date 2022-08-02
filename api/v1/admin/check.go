package admin

import (
	"context"

	appusergrpc "github.com/NpoolPlatform/appuser-manager/pkg/client/appuser"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/const"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool"
	appuserpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/admin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(ctx context.Context, info *admin.CreateGenesisUserRequest) error {
	if info.GetAppID() != constant.GenesisAppID && info.GetAppID() != constant.ChurchAppID {
		return status.Error(codes.PermissionDenied, "invalid app id for genesis user")
	}
	exist, err := appusergrpc.ExistAppUserConds(ctx, &appuserpb.Conds{
		AppID: &npool.StringVal{
			Op:    cruder.EQ,
			Value: info.GetAppID(),
		},
		EmailAddress: &npool.StringVal{
			Op:    cruder.EQ,
			Value: info.GetEmailAddress(),
		},
	})
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	if exist {
		return status.Error(codes.AlreadyExists, "genesis user already exists")
	}

	return nil
}
