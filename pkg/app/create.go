package app

import (
	"context"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"

	appmgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"
	appctrlmgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"

	appmgrcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/app"
	appctrlmgrcrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/appcontrol"
)

func CreateApp(ctx context.Context, in *npool.AppReq) (*App, error) {
	var id string

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info, err := appmgrcrud.CreateAppTx(ctx, tx, &appmgrpb.AppReq{
			ID:          in.ID,
			CreatedBy:   in.CreatedBy,
			Name:        in.Name,
			Logo:        in.Logo,
			Description: in.Description,
		}).Save(ctx)
		if err != nil {
			return err
		}

		id = info.ID.String()

		if _, err := appctrlmgrcrud.CreateAppControlTx(ctx, tx, &appctrlmgrpb.AppControlReq{
			AppID:               &id,
			SignupMethods:       in.SignupMethods,
			ExternSigninMethods: in.ExtSigninMethods,
			RecaptchaMethod:     in.RecaptchaMethod,
			KycEnable:           in.KycEnable,
			SigninVerifyEnable:  in.SigninVerifyEnable,
			InvitationCodeMust:  in.InvitationCodeMust,
		}).Save(ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetApp(ctx, id)
}
