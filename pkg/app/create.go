package app

import (
	"context"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"time"

	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"

	"github.com/google/uuid"
)

func CreateApp(ctx context.Context, in *npool.AppReq) (*AppCreateResp, error) {
	var err error
	info := AppCreateResp{}

	if in.ID == nil || in.GetID() == "" {
		id := uuid.NewString()
		in.ID = &id
	}
	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info.App, err = appCreate(ctx, tx, in)
		if err != nil {
			return err
		}

		info.BanApp, err = banAppCreate(ctx, tx, in)
		if err != nil {
			return err
		}

		info.AppControl, err = appControlCreate(ctx, tx, in)
		if err != nil {
			return err
		}
		return nil
	})

	return &info, nil
}

func appCreate(ctx context.Context, tx *ent.Tx, in *npool.AppReq) (*ent.App, error) {
	cTx := tx.App.Create()
	if in.ID != nil {
		cTx.SetID(uuid.MustParse(in.GetID()))
	}
	if in.CreatedBy != nil {
		cTx.SetCreatedBy(uuid.MustParse(in.GetCreatedBy()))
	}
	if in.Name != nil {
		cTx.SetName(in.GetName())
	}
	if in.Logo != nil {
		cTx.SetLogo(in.GetLogo())
	}
	if in.Description != nil {
		cTx.SetDescription(in.GetDescription())
	}
	info, err := cTx.Save(ctx)
	if err != nil {
		logger.Sugar().Errorw("fail create app: %v", err)
		return nil, err
	}
	return info, nil
}

func banAppCreate(ctx context.Context, tx *ent.Tx, in *npool.AppReq) (*ent.BanApp, error) {
	cTx := tx.BanApp.Create()

	cTx.SetAppID(uuid.MustParse(in.GetID()))

	if in.BanMessage != nil {
		cTx.SetMessage(in.GetBanMessage())
	}
	if in.Banned != nil && in.GetBanned() {
		cTx.SetDeletedAt(uint32(time.Now().Unix()))
	}
	info, err := cTx.Save(ctx)
	if err != nil {
		logger.Sugar().Errorw("fail create ban app: %v", err)
		return nil, err
	}
	return info, nil
}

func appControlCreate(ctx context.Context, tx *ent.Tx, in *npool.AppReq) (*ent.AppControl, error) {
	cTx := tx.AppControl.Create()

	cTx.SetAppID(uuid.MustParse(in.GetID()))

	if in.SignupMethods != nil {
		cTx.SetSignupMethods(in.GetSignupMethods())
	}
	if in.ExtSigninMethods != nil {
		cTx.SetExternSigninMethods(in.GetExtSigninMethods())
	}
	if in.RecaptchaMethod != nil {
		cTx.SetRecaptchaMethod(in.GetRecaptchaMethod())
	}
	if in.KycEnable != nil {
		cTx.SetKycEnable(in.GetKycEnable())
	}
	if in.SigninVerifyEnable != nil {
		cTx.SetSigninVerifyEnable(in.GetSigninVerifyEnable())
	}
	if in.InvitationCodeMust != nil {
		cTx.SetInvitationCodeMust(in.GetInvitationCodeMust())
	}
	info, err := cTx.Save(ctx)
	if err != nil {
		logger.Sugar().Errorw("fail create app control: %v", err)
		return nil, err
	}
	return info, nil
}
