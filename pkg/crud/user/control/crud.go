package control

import (
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entappuserctrl "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appusercontrol"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/google/uuid"
)

type Req struct {
	ID                 *uuid.UUID
	AppID              *uuid.UUID
	UserID             *uuid.UUID
	GoogleAuthVerified *bool
	SigninVerifyType   *basetypes.SignMethod
	Kol                *bool
	KolConfirmed       *bool
	SelectedLangID     *uuid.UUID
}

func CreateSet(c *ent.AppUserControlCreate, req *Req) *ent.AppUserControlCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.GoogleAuthVerified != nil {
		c.SetGoogleAuthenticationVerified(*req.GoogleAuthVerified)
	}
	if req.SigninVerifyType != nil {
		c.SetSigninVerifyType(req.SigninVerifyType.String())
	}
	if req.Kol != nil {
		c.SetKol(*req.Kol)
	}
	if req.KolConfirmed != nil {
		c.SetKolConfirmed(*req.KolConfirmed)
	}
	if req.SelectedLangID != nil {
		c.SetSelectedLangID(*req.SelectedLangID)
	}
	return c
}

func UpdateSet(u *ent.AppUserControlUpdateOne, req *Req) *ent.AppUserControlUpdateOne {
	if req.GoogleAuthVerified != nil {
		u.SetGoogleAuthenticationVerified(*req.GoogleAuthVerified)
	}
	if req.SigninVerifyType != nil {
		u.SetSigninVerifyType(req.SigninVerifyType.String())
	}
	if req.Kol != nil {
		u.SetKol(*req.Kol)
	}
	if req.KolConfirmed != nil {
		u.SetKolConfirmed(*req.KolConfirmed)
	}
	if req.SelectedLangID != nil {
		u.SetSelectedLangID(*req.SelectedLangID)
	}
	return u
}

type Conds struct {
	ID           *cruder.Cond
	AppID        *cruder.Cond
	UserID       *cruder.Cond
	Kol          *cruder.Cond
	KolConfirmed *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.AppUserControlQuery, conds *Conds) (*ent.AppUserControlQuery, error) {
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
			q.Where(entappuserctrl.ID(id))
		default:
			return nil, fmt.Errorf("invalid appusercontrol field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappuserctrl.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appusercontrol field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entappuserctrl.UserID(id))
		default:
			return nil, fmt.Errorf("invalid appusercontrol field")
		}
	}
	if conds.Kol != nil {
		kol, ok := conds.Kol.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid kol")
		}
		switch conds.Kol.Op {
		case cruder.EQ:
			q.Where(entappuserctrl.Kol(kol))
		case cruder.NEQ:
			q.Where(entappuserctrl.KolNEQ(kol))
		default:
			return nil, fmt.Errorf("invalid appusercontrol field")
		}
	}
	if conds.KolConfirmed != nil {
		confirmed, ok := conds.KolConfirmed.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid kol confirmed")
		}
		switch conds.KolConfirmed.Op {
		case cruder.EQ:
			q.Where(entappuserctrl.KolConfirmed(confirmed))
		case cruder.NEQ:
			q.Where(entappuserctrl.KolConfirmedNEQ(confirmed))
		default:
			return nil, fmt.Errorf("invalid appusercontrol field")
		}
	}
	return q, nil
}
