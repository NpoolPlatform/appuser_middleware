package history

import (
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entloginhistory "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/loginhistory"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	ClientIP  *string
	UserAgent *string
	Location  *string
	LoginType *basetypes.LoginType
}

func CreateSet(c *ent.LoginHistoryCreate, req *Req) *ent.LoginHistoryCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.ClientIP != nil {
		c.SetClientIP(*req.ClientIP)
	}
	if req.UserAgent != nil {
		c.SetUserAgent(*req.UserAgent)
	}
	if req.Location != nil {
		c.SetLocation(*req.Location)
	}
	if req.LoginType != nil {
		c.SetLoginType(req.LoginType.String())
	}
	return c
}

func UpdateSet(u *ent.LoginHistoryUpdateOne, req *Req) *ent.LoginHistoryUpdateOne {
	if req.Location != nil {
		u.SetLocation(*req.Location)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	AppID     *cruder.Cond
	UserID    *cruder.Cond
	LoginType *cruder.Cond
	ClientIP  *cruder.Cond
	Location  *cruder.Cond
}

func SetQueryConds(q *ent.LoginHistoryQuery, conds *Conds) (*ent.LoginHistoryQuery, error) { //nolint
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
			q.Where(entloginhistory.ID(id))
		default:
			return nil, fmt.Errorf("invalid login history field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entloginhistory.AppID(id))
		default:
			return nil, fmt.Errorf("invalid login history field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entloginhistory.UserID(id))
		default:
			return nil, fmt.Errorf("invalid login history field")
		}
	}
	if conds.LoginType != nil {
		loginType, ok := conds.LoginType.Val.(basetypes.LoginType)
		if !ok {
			return nil, fmt.Errorf("invalid login type")
		}
		switch conds.LoginType.Op {
		case cruder.EQ:
			q.Where(entloginhistory.LoginType(loginType.String()))
		default:
			return nil, fmt.Errorf("invalid login history field")
		}
	}
	if conds.ClientIP != nil {
		ip, ok := conds.ClientIP.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid client ip")
		}
		switch conds.ClientIP.Op {
		case cruder.EQ:
			q.Where(entloginhistory.ClientIP(ip))
		default:
			return nil, fmt.Errorf("invalid login history field")
		}
	}
	if conds.Location != nil {
		loc, ok := conds.Location.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid location")
		}
		switch conds.Location.Op {
		case cruder.EQ:
			q.Where(entloginhistory.Location(loc))
		default:
			return nil, fmt.Errorf("invalid login history field")
		}
	}
	return q, nil
}
