package history

import (
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entauthhistory "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/authhistory"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID       *uuid.UUID
	AppID    *uuid.UUID
	UserID   *uuid.UUID
	Resource *string
	Method   *string
	Allowed  *bool
}

func CreateSet(c *ent.AuthHistoryCreate, req *Req) *ent.AuthHistoryCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.Resource != nil {
		c.SetResource(*req.Resource)
	}
	if req.Method != nil {
		c.SetMethod(*req.Method)
	}
	if req.Allowed != nil {
		c.SetAllowed(*req.Allowed)
	}
	return c
}

func UpdateSet(u *ent.AuthHistoryUpdateOne, in *Req) *ent.AuthHistoryUpdateOne {
	return u
}

type Conds struct {
	ID       *cruder.Cond
	AppID    *cruder.Cond
	UserID   *cruder.Cond
	Resource *cruder.Cond
	Method   *cruder.Cond
	Allowed  *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.AuthHistoryQuery, conds *Conds) (*ent.AuthHistoryQuery, error) {
	if conds == nil {
		return stm, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			stm.Where(entauthhistory.ID(id))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			stm.Where(entauthhistory.AppID(id))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			stm.Where(entauthhistory.UserID(id))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.Resource != nil {
		res, ok := conds.Resource.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid resource")
		}
		switch conds.Resource.Op {
		case cruder.EQ:
			stm.Where(entauthhistory.Resource(res))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.Method != nil {
		method, ok := conds.Method.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid method")
		}
		switch conds.Method.Op {
		case cruder.EQ:
			stm.Where(entauthhistory.Method(method))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.Allowed != nil {
		allowed, ok := conds.Method.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid allowed")
		}
		switch conds.Allowed.Op {
		case cruder.EQ:
			stm.Where(entauthhistory.Allowed(allowed))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	return stm, nil
}
