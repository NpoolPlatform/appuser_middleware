package app

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-middleware/pkg/tracer/app"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"

	appmgrcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/app"
	appctrlmgrcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appcontrol"
	appmgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"
	appctrlmgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
)

func UpdateApp(ctx context.Context, in *npool.AppReq) (*App, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateApp")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	span = commontracer.TraceInvoker(span, "app", "db", "UpdateTx")

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		_, err := appmgrcrud.UpdateTx(tx, &appmgrpb.AppReq{
			ID:          in.ID,
			Name:        in.Name,
			Logo:        in.Logo,
			Description: in.Description,
		}).Save(ctx)
		if err != nil {
			logger.Sugar().Errorw("UpdateApp", "error", err)
			return err
		}

		if _, err = appctrlmgrcrud.UpdateTx(tx, &appctrlmgrpb.AppControlReq{
			AppID:               in.ID,
			SignupMethods:       in.SignupMethods,
			ExternSigninMethods: in.ExtSigninMethods,
			RecaptchaMethod:     in.RecaptchaMethod,
			KycEnable:           in.KycEnable,
			SigninVerifyEnable:  in.SigninVerifyEnable,
			InvitationCodeMust:  in.InvitationCodeMust,
		}).Save(ctx); err != nil {
			logger.Sugar().Errorw("UpdateApp", "error", err)
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return GetApp(ctx, in.GetID())
}
