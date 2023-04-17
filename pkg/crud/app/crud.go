package app

import (
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entapp "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/app"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID          *uuid.UUID
	CreatedBy   *uuid.UUID
	Name        *string
	Logo        *string
	Description *string
}

func CreateSet(c *ent.AppCreate, req *Req) *ent.AppCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.CreatedBy != nil {
		c.SetCreatedBy(*req.CreatedBy)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Logo != nil {
		c.SetLogo(*req.Logo)
	}
	if req.Description != nil {
		c.SetDescription(*req.Description)
	}
	return c
}

func UpdateSet(u *ent.AppUpdateOne, req *Req) *ent.AppUpdateOne {
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.Logo != nil {
		u.SetLogo(*req.Logo)
	}
	if req.Description != nil {
		u.SetDescription(*req.Description)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	IDs       *cruder.Cond
	CreatedBy *cruder.Cond
	Name      *cruder.Cond
}

func SetQueryConds(q *ent.AppQuery, conds *Conds) (*ent.AppQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid app id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entapp.ID(id))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid app ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entapp.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.CreatedBy != nil {
		createdBy, ok := conds.CreatedBy.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid creator")
		}
		switch conds.CreatedBy.Op {
		case cruder.EQ:
			q.Where(entapp.CreatedBy(createdBy))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(entapp.Name(name))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}
	return q, nil
}
