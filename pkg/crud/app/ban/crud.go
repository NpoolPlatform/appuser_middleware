package ban

import (
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entbanapp "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/banapp"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID      *uuid.UUID
	AppID   *uuid.UUID
	Message *string
}

func CreateSet(c *ent.BanAppCreate, req *Req) *ent.BanAppCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	return c
}

func UpdateSet(u *ent.BanAppUpdateOne, req *Req) *ent.BanAppUpdateOne {
	if req.Message != nil {
		u.SetMessage(*req.Message)
	}
	return u
}

type Conds struct {
	ID    *cruder.Cond
	AppID *cruder.Cond
}

func SetQueryConds(q *ent.BanAppQuery, conds *Conds) (*ent.BanAppQuery, error) {
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
			q.Where(entbanapp.ID(id))
		default:
			return nil, fmt.Errorf("invalid banapp field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entbanapp.AppID(id))
		default:
			return nil, fmt.Errorf("invalid banapp field")
		}
	}
	q.Where(entbanapp.DeletedAt(0))
	return q, nil
}
