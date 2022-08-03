package admin

import (
	"context"

	"github.com/google/uuid"

	approlegrpc "github.com/NpoolPlatform/appuser-manager/pkg/client/approle"
	approleusergrpc "github.com/NpoolPlatform/appuser-manager/pkg/client/approleuser"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/const"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool"
	approlepb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approle"
	approleuserpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approleuser"
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/admin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(ctx context.Context, info *admin.CreateGenesisUserRequest) error {
	if info.GetAppID() != constant.GenesisAppID && info.GetAppID() != constant.ChurchAppID {
		return status.Error(codes.PermissionDenied, "invalid app id for genesis user")
	}

	resp, err := approlegrpc.GetAppRoleOnly(ctx, &approlepb.Conds{
		AppID: &npool.StringVal{
			Op:    cruder.EQ,
			Value: uuid.UUID{}.String(),
		},
		Role: &npool.StringVal{
			Op:    cruder.EQ,
			Value: constant.GenesisRole,
		},
	})
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	exist, err := approleusergrpc.ExistAppRoleUserConds(ctx, &approleuserpb.Conds{
		AppID: &npool.StringVal{
			Op:    cruder.EQ,
			Value: uuid.UUID{}.String(),
		},
		RoleID: &npool.StringVal{
			Op:    cruder.EQ,
			Value: resp.ID,
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
