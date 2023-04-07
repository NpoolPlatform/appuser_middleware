package app

import (
	"context"

	"github.com/google/uuid"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	entappcontrol "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appcontrol"
	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-middleware/pkg/tracer/app"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"

	appmgrcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/app"
	appctrlmgrcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appcontrol"
	appmgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"
	appctrlmgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
)

func UpdateApp(ctx context.Context, in *npool.AppReq) (*npool.App, error) {
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
		_, err := appmgrcrud.UpdateSet(
			tx.App.
				UpdateOneID(
					uuid.MustParse(in.GetID()),
				),
			&appmgrpb.AppReq{
				ID:          in.ID,
				Name:        in.Name,
				Logo:        in.Logo,
				Description: in.Description,
			}).Save(ctx)
		if err != nil {
			logger.Sugar().Errorw("UpdateApp", "error", err)
			return err
		}

		info, err := tx.
			AppControl.
			Query().
			Where(
				entappcontrol.AppID(uuid.MustParse(in.GetID())),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		if _, err = appctrlmgrcrud.UpdateSet(
			info,
			&appctrlmgrpb.AppControlReq{
				AppID:                    in.ID,
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
