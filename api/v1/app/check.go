package app

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	mgrapp "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"
	mgrappcontrol "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/appuser-manager/api/v2/app"
	"github.com/NpoolPlatform/appuser-manager/api/v2/appcontrol"
)

func validate(info *npool.AppReq) error {
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

	return nil
}

func Validate(info *npool.AppReq) error {
	return validate(info)
}
