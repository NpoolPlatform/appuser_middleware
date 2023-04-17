package control

import (
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entappctrl "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appcontrol"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

type Req struct {
	ID                       *uuid.UUID
	AppID                    *uuid.UUID
	SignupMethods            []basetypes.SignMethod
	ExtSigninMethods         []basetypes.SignMethod
	RecaptchaMethod          *basetypes.RecaptchaMethod
	KycEnable                *bool
	SigninVerifyEnable       *bool
	InvitationCodeMust       *bool
	CreateInvitationCodeWhen *basetypes.CreateInvitationCodeWhen
	MaxTypedCouponsPerOrder  *uint32
	Maintaining              *bool
	CommitButtonTargets      []string
}

func CreateSet(c *ent.AppControlCreate, req *Req) *ent.AppControlCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if len(req.SignupMethods) > 0 {
		methods := []string{}
		for _, m := range req.SignupMethods {
			methods = append(methods, m.String())
		}
		c.SetSignupMethods(methods)
	}
	if len(req.ExtSigninMethods) > 0 {
		methods := []string{}
		for _, m := range req.ExtSigninMethods {
			methods = append(methods, m.String())
		}
		c.SetExternSigninMethods(methods)
	}
	if req.RecaptchaMethod != nil {
		c.SetRecaptchaMethod(req.RecaptchaMethod.String())
	}
	if req.KycEnable != nil {
		c.SetKycEnable(*req.KycEnable)
	}
	if req.SigninVerifyEnable != nil {
		c.SetSigninVerifyEnable(*req.SigninVerifyEnable)
	}
	if req.InvitationCodeMust != nil {
		c.SetInvitationCodeMust(*req.InvitationCodeMust)
	}
	if req.CreateInvitationCodeWhen != nil {
		c.SetCreateInvitationCodeWhen(req.CreateInvitationCodeWhen.String())
	}
	if req.MaxTypedCouponsPerOrder != nil {
		c.SetMaxTypedCouponsPerOrder(*req.MaxTypedCouponsPerOrder)
	}
	if req.Maintaining != nil {
		c.SetMaintaining(*req.Maintaining)
	}
	if len(req.CommitButtonTargets) > 0 {
		c.SetCommitButtonTargets(req.CommitButtonTargets)
	}
	return c
}

func UpdateSet(u *ent.AppControlUpdateOne, req *Req) *ent.AppControlUpdateOne {
	if len(req.SignupMethods) > 0 {
		methods := []string{}
		for _, m := range req.SignupMethods {
			methods = append(methods, m.String())
		}
		u.SetSignupMethods(methods)
	}
	if len(req.ExtSigninMethods) > 0 {
		methods := []string{}
		for _, m := range req.ExtSigninMethods {
			methods = append(methods, m.String())
		}
		u.SetExternSigninMethods(methods)
	}
	if req.RecaptchaMethod != nil {
		u.SetRecaptchaMethod(req.RecaptchaMethod.String())
	}
	if req.KycEnable != nil {
		u.SetKycEnable(*req.KycEnable)
	}
	if req.SigninVerifyEnable != nil {
		u.SetSigninVerifyEnable(*req.SigninVerifyEnable)
	}
	if req.InvitationCodeMust != nil {
		u.SetInvitationCodeMust(*req.InvitationCodeMust)
	}
	if req.CreateInvitationCodeWhen != nil {
		u.SetCreateInvitationCodeWhen(req.CreateInvitationCodeWhen.String())
	}
	if req.MaxTypedCouponsPerOrder != nil {
		u.SetMaxTypedCouponsPerOrder(*req.MaxTypedCouponsPerOrder)
	}
	if req.Maintaining != nil {
		u.SetMaintaining(*req.Maintaining)
	}
	if len(req.CommitButtonTargets) > 0 {
		u.SetCommitButtonTargets(req.CommitButtonTargets)
	}
	return u
}

type Conds struct {
	ID    *cruder.Cond
	AppID *cruder.Cond
}

func SetQueryConds(q *ent.AppControlQuery, conds *Conds) (*ent.AppControlQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entappctrl.ID(id))
		default:
			return nil, fmt.Errorf("invalid appcontrol field")
		}
	}

	if conds.AppID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappctrl.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appcontrol field")
		}
	}

	return q, nil
}
