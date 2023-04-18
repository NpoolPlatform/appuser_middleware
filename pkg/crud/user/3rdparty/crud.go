package thirdparty

import (
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entappuserthirdparty "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuserthirdparty"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID                 *uuid.UUID
	AppID              *uuid.UUID
	UserID             *uuid.UUID
	ThirdPartyID       *uuid.UUID
	ThirdPartyUserID   *string
	ThirdPartyUsername *string
	ThirdPartyAvatar   *string
}

func CreateSet(c *ent.AppUserThirdPartyCreate, req *Req) *ent.AppUserThirdPartyCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.ThirdPartyUserID != nil {
		c.SetThirdPartyUserID(*req.ThirdPartyUserID)
	}
	if req.ThirdPartyID != nil {
		c.SetThirdPartyID(*req.ThirdPartyID)
	}
	if req.ThirdPartyUsername != nil {
		c.SetThirdPartyUsername(*req.ThirdPartyUsername)
	}
	if req.ThirdPartyAvatar != nil {
		c.SetThirdPartyAvatar(*req.ThirdPartyAvatar)
	}
	return c
}

func UpdateSet(u *ent.AppUserThirdPartyUpdateOne, req *Req) *ent.AppUserThirdPartyUpdateOne {
	if req.ThirdPartyUsername != nil {
		u.SetThirdPartyUsername(*req.ThirdPartyUsername)
	}
	if req.ThirdPartyAvatar != nil {
		u.SetThirdPartyAvatar(*req.ThirdPartyAvatar)
	}
	return u
}

type Conds struct {
	ID               *cruder.Cond
	AppID            *cruder.Cond
	UserID           *cruder.Cond
	ThirdPartyUserID *cruder.Cond
	ThirdPartyID     *cruder.Cond
}

//nolint:nolintlint,gocyclo
func SetQueryConds(q *ent.AppUserThirdPartyQuery, conds *Conds) (*ent.AppUserThirdPartyQuery, error) {
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
			q.Where(entappuserthirdparty.ID(id))

		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappuserthirdparty.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entappuserthirdparty.UserID(id))
		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	if conds.ThirdPartyUserID != nil {
		id, ok := conds.UserID.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid 3rd userid")
		}
		switch conds.ThirdPartyUserID.Op {
		case cruder.EQ:
			q.Where(entappuserthirdparty.ThirdPartyUserID(id))
		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	if conds.ThirdPartyID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid 3rd id")
		}
		switch conds.ThirdPartyID.Op {
		case cruder.EQ:
			q.Where(entappuserthirdparty.ThirdPartyID(id))
		default:
			return nil, fmt.Errorf("invalid appuserthirdparty field")
		}
	}
	return q, nil
}
