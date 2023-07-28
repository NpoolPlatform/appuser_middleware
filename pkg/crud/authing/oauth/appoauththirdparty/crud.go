package auth

import (
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entappoauththirdparty "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appoauththirdparty"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID           *uuid.UUID
	AppID        *uuid.UUID
	ThirdPartyID *uuid.UUID
	DeletedAt    *uint32
}

func CreateSet(c *ent.AppOAuthThirdPartyCreate, req *Req) *ent.AppOAuthThirdPartyCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.ThirdPartyID != nil {
		c.SetThirdPartyID(*req.ThirdPartyID)
	}
	return c
}

func UpdateSet(u *ent.AppOAuthThirdPartyUpdateOne, req *Req) *ent.AppOAuthThirdPartyUpdateOne {
	if req.ClientID != nil {
		u.SetClientID(*req.ClientID)
	}
	if req.ClientSecret != nil {
		u.SetClientSecret(*req.ClientSecret)
	}
	if req.CallbackURL != nil {
		u.SetCallbackURL(*req.CallbackURL)
	}
	if req.Salt != nil {
		u.SetSalt(*req.Salt)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID            *cruder.Cond
	IDs           *cruder.Cond
	AppID         *cruder.Cond
	ThirdPartyID  *cruder.Cond
	ThirdPartyIDs *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.AppOAuthThirdPartyQuery, conds *Conds) (*ent.AppOAuthThirdPartyQuery, error) {
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
			q.Where(entappoauththirdparty.ID(id))
		default:
			return nil, fmt.Errorf("invalid oauth field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entappoauththirdparty.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid oauth field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappoauththirdparty.AppID(id))
		default:
			return nil, fmt.Errorf("invalid oauth field")
		}
	}
	if conds.ThirdPartyID != nil {
		id, ok := conds.ThirdPartyID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid thirdpartyid")
		}
		switch conds.ThirdPartyID.Op {
		case cruder.EQ:
			q.Where(entappoauththirdparty.ThirdPartyID(id))
		default:
			return nil, fmt.Errorf("invalid oauth field")
		}
	}
	if conds.ThirdPartyIDs != nil {
		ids, ok := conds.ThirdPartyIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid thirdpartyids")
		}
		switch conds.ThirdPartyIDs.Op {
		case cruder.IN:
			q.Where(entappoauththirdparty.ThirdPartyIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid oauth field")
		}
	}
	q.Where(entappoauththirdparty.DeletedAt(0))
	return q, nil
}
