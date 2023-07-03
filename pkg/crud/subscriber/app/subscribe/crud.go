package appsubscribe

import (
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entappsubscribe "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appsubscribe"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID             *uuid.UUID
	AppID          *uuid.UUID
	SubscribeAppID *uuid.UUID
	DeletedAt      *uint32
}

func CreateSet(c *ent.AppSubscribeCreate, req *Req) *ent.AppSubscribeCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.SubscribeAppID != nil {
		c.SetSubscribeAppID(*req.SubscribeAppID)
	}
	return c
}

func UpdateSet(u *ent.AppSubscribeUpdateOne, req *Req) *ent.AppSubscribeUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID             *cruder.Cond
	AppID          *cruder.Cond
	SubscribeAppID *cruder.Cond
}

//nolint:nolintlint,gocyclo
func SetQueryConds(q *ent.AppSubscribeQuery, conds *Conds) (*ent.AppSubscribeQuery, error) {
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
			q.Where(entappsubscribe.ID(id))
		default:
			return nil, fmt.Errorf("invalid appsubscribe field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappsubscribe.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appsubscribe field")
		}
	}
	if conds.SubscribeAppID != nil {
		id, ok := conds.SubscribeAppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid subscribe appid")
		}
		switch conds.SubscribeAppID.Op {
		case cruder.EQ:
			q.Where(entappsubscribe.SubscribeAppID(id))
		default:
			return nil, fmt.Errorf("invalid appsubscribe field")
		}
	}
	q.Where(entappsubscribe.DeletedAt(0))
	return q, nil
}
