package app

import (
	"context"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-middleware/pkg/tracer/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"

	appmgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"
	appctrlmgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"

	appmgrcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/app"
	appctrlmgrcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appcontrol"
)

func CreateApp(ctx context.Context, in *npool.AppReq) (*npool.App, error) {
	var id string
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateApp")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	span = commontracer.TraceInvoker(span, "app", "db", "CreateTx")

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info, err := appmgrcrud.CreateSet(tx.App.Create(), &appmgrpb.AppReq{
			ID:          in.ID,
			CreatedBy:   in.CreatedBy,
			Name:        in.Name,
			Logo:        in.Logo,
			Description: in.Description,
		}).Save(ctx)
		if err != nil {
			logger.Sugar().Errorw("CreateApp", "error", err)
			return err
		}

		id = info.ID.String()

		if _, err := appctrlmgrcrud.CreateSet(tx.AppControl.Create(), &appctrlmgrpb.AppControlReq{
			AppID:                    &id,
			SignupMethods:            in.SignupMethods,
			ExtSigninMethods:         in.ExtSigninMethods,
			RecaptchaMethod:          in.RecaptchaMethod,
			KycEnable:                in.KycEnable,
			SigninVerifyEnable:       in.SigninVerifyEnable,
			InvitationCodeMust:       in.InvitationCodeMust,
			CreateInvitationCodeWhen: in.CreateInvitationCodeWhen,
			MaxTypedCouponsPerOrder:  in.MaxTypedCouponsPerOrder,
			Maintaining:              in.Maintaining,
			CommitButtonTargets:      in.CommitButtonTargets,
		}).Save(ctx); err != nil {
			logger.Sugar().Errorw("CreateApp", "error", err)
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetApp(ctx, id)
}
