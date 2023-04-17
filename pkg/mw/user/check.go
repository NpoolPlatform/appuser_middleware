package user

import (
	"context"
	"fmt"

	appusermgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"

	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuser"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
)

func (h *Handler) checkAccountExist(ctx context.Context) error {
	if h.PhoneNO == nil && h.EmailAddress == nil {
		return fmt.Errorf("invalid account")
	}

	conds := &appusermgrpb.Conds{
		AppID: &commonpb.StringVal{Op: cruder.EQ, Value: h.AppID},
	}
	if h.EmailAddress != nil {
		conds.EmailAddress = &commonpb.StringVal{Op: cruder.EQ, Value: *h.EmailAddress}
	}
	if h.PhoneNO != nil {
		conds.PhoneNO = &commonpb.StringVal{Op: cruder.EQ, Value: *h.PhoneNO}
	}

	exist, err := appusercrud.ExistConds(ctx, conds)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("user already exist")
	}

	return nil
}
