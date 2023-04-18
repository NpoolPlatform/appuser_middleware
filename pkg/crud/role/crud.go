package role

import (
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entapprole "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approle"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID          *uuid.UUID
	AppID       *uuid.UUID
	CreatedBy   *uuid.UUID
	Role        *string
	Description *string
	Default     *bool
	Genesis     *bool
	DeletedAt   *uint32
}

func CreateSet(c *ent.AppRoleCreate, req *Req) *ent.AppRoleCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.CreatedBy != nil {
		c.SetCreatedBy(*req.CreatedBy)
	}
	if req.Role != nil {
		c.SetRole(*req.Role)
	}
	if req.Description != nil {
		c.SetDescription(*req.Description)
	}
	if req.Default != nil {
		c.SetDefault(*req.Default)
	}
	if req.Genesis != nil {
		c.SetGenesis(*req.Genesis)
	}
	return c
}

func UpdateSet(u *ent.AppRoleUpdateOne, req *Req) *ent.AppRoleUpdateOne {
	if req.Role != nil {
		u.SetRole(*req.Role)
	}
	if req.Description != nil {
		u.SetDescription(*req.Description)
	}
	if req.Default != nil {
		u.SetDefault(*req.Default)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	AppID     *cruder.Cond
	CreatedBy *cruder.Cond
	Role      *cruder.Cond
	Default   *cruder.Cond
	Genesis   *cruder.Cond
	Roles     *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.AppRoleQuery, conds *Conds) (*ent.AppRoleQuery, error) {
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
			q.Where(entapprole.ID(id))
		default:
			return nil, fmt.Errorf("invalid approle field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entapprole.AppID(id))
		default:
			return nil, fmt.Errorf("invalid approle field")
		}
	}
	if conds.CreatedBy != nil {
		id, ok := conds.CreatedBy.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.CreatedBy.Op {
		case cruder.EQ:
			q.Where(entapprole.CreatedBy(id))
		default:
			return nil, fmt.Errorf("invalid approle field")
		}
	}
	if conds.Role != nil {
		role, ok := conds.Role.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid role")
		}
		switch conds.Role.Op {
		case cruder.EQ:
			q.Where(entapprole.Role(role))
		default:
			return nil, fmt.Errorf("invalid approle field")
		}
	}
	if conds.Default != nil {
		defautl, ok := conds.Default.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid default")
		}
		switch conds.Default.Op {
		case cruder.EQ:
			q.Where(entapprole.Default(defautl))
		default:
			return nil, fmt.Errorf("invalid approle field")
		}
	}
	if conds.Genesis != nil {
		genesis, ok := conds.Genesis.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid default")
		}
		switch conds.Genesis.Op {
		case cruder.EQ:
			q.Where(entapprole.Genesis(genesis))
		default:
			return nil, fmt.Errorf("invalid approle field")
		}
	}
	if conds.Roles != nil {
		roles, ok := conds.Roles.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid roles")
		}
		switch conds.Roles.Op {
		case cruder.IN:
			q.Where(entapprole.RoleIn(roles...))
		default:
			return nil, fmt.Errorf("invalid approle field")
		}
	}
	return q, nil
}
