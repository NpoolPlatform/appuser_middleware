package user

import (
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entappuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuser"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID            *uuid.UUID
	AppID         *uuid.UUID
	EmailAddress  *string
	PhoneNO       *string
	ImportFromApp *uuid.UUID
	DeletedAt     *uint32
}

func CreateSet(c *ent.AppUserCreate, req *Req) *ent.AppUserCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.EmailAddress != nil {
		c.SetEmailAddress(*req.EmailAddress)
	}
	if req.PhoneNO != nil {
		c.SetPhoneNo(*req.PhoneNO)
	}
	if req.ImportFromApp != nil {
		c.SetImportFromApp(*req.ImportFromApp)
	}
	return c
}

func UpdateSet(u *ent.AppUserUpdateOne, req *Req) *ent.AppUserUpdateOne {
	if req.EmailAddress != nil {
		u.SetEmailAddress(*req.EmailAddress)
	}
	if req.PhoneNO != nil {
		u.SetPhoneNo(*req.PhoneNO)
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
	EmailAddress  *cruder.Cond
	PhoneNO       *cruder.Cond
	ImportFromApp *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.AppUserQuery, conds *Conds) (*ent.AppUserQuery, error) {
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
			q.Where(entappuser.ID(id))
		default:
			return nil, fmt.Errorf("invalid appuser field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		if len(ids) > 0 {
			switch conds.IDs.Op {
			case cruder.IN:
				q.Where(entappuser.IDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid appuser field")
			}
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappuser.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appuser field")
		}
	}
	if conds.EmailAddress != nil {
		addr, ok := conds.EmailAddress.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid email address")
		}
		switch conds.EmailAddress.Op {
		case cruder.EQ:
			q.Where(entappuser.EmailAddress(addr))
		default:
			return nil, fmt.Errorf("invalid appuser field")
		}
	}
	if conds.PhoneNO != nil {
		no, ok := conds.PhoneNO.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid phone no")
		}
		switch conds.PhoneNO.Op {
		case cruder.EQ:
			q.Where(entappuser.PhoneNo(no))
		default:
			return nil, fmt.Errorf("invalid appuser field")
		}
	}
	if conds.ImportFromApp != nil {
		id, ok := conds.ImportFromApp.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid import from app")
		}
		switch conds.ImportFromApp.Op {
		case cruder.EQ:
			q.Where(entappuser.ImportFromApp(id))
		default:
			return nil, fmt.Errorf("invalid appuser field")
		}
	}
	q.Where(entappuser.DeletedAt(0))
	return q, nil
}
