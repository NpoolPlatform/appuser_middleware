package app

import (
	"context"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npoolpb "github.com/NpoolPlatform/message/npool"
	mgrapp "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"
	mgrappcontrol "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/appuser-manager/api/app"
	"github.com/NpoolPlatform/appuser-manager/api/appcontrol"
	appcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/app"
)

func validate(ctx context.Context, info *npool.AppReq) error {
	err := app.Validate(&mgrapp.AppReq{
		Description: info.Description,
		CreatedBy:   info.CreatedBy,
		Name:        info.Name,
		Logo:        info.Logo,
	})
	if err != nil {
		logger.Sugar().Errorw("validate", err.Error())
		return status.Error(codes.InvalidArgument, err.Error())
	}

	err = appcontrol.Validate(&mgrappcontrol.AppControlReq{
		AppID: info.ID,
	})
	if err != nil {
		logger.Sugar().Errorw("validate", err.Error())
		return status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := appcrud.ExistConds(ctx, &mgrapp.Conds{
		Name: &npoolpb.StringVal{
			Op:    cruder.EQ,
			Value: info.GetName(),
		},
	})

	if err != nil {
		logger.Sugar().Errorw("validate", err.Error())
		return status.Error(codes.Internal, err.Error())
	}

	if exist {
		logger.Sugar().Errorw("validate", "app name already exists")
		return status.Error(codes.AlreadyExists, "app name already exists")
	}
	return nil
}

func Validate(ctx context.Context, info *npool.AppReq) error {
	return validate(ctx, info)
}
