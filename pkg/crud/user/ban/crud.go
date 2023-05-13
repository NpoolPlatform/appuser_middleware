package ban

import (
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entbanappuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/banappuser"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	Message   *string
	DeletedAt *uint32
}

func CreateSet(c *ent.BanAppUserCreate, req *Req) *ent.BanAppUserCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	return c
}

func UpdateSet(u *ent.BanAppUserUpdateOne, req *Req) *ent.BanAppUserUpdateOne {
	if req.Message != nil {
		u.SetMessage(*req.Message)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID     *cruder.Cond
	AppID  *cruder.Cond
	UserID *cruder.Cond
}

func SetQueryConds(q *ent.BanAppUserQuery, conds *Conds) (*ent.BanAppUserQuery, error) {
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
			q.Where(entbanappuser.ID(id))
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
			q.Where(entbanappuser.AppID(id))
		default:
			return nil, fmt.Errorf("invalid banapp field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entbanappuser.UserID(id))
		default:
			return nil, fmt.Errorf("invalid banapp field")
		}
	}
	q.Where(entbanappuser.DeletedAt(0))
	return q, nil
}
