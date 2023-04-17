package history

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/authhistory"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

func CreateSet(c *ent.AuthHistoryCreate, info *npool.HistoryReq) *ent.AuthHistoryCreate {
	if info.ID != nil {
		c.SetID(uuid.MustParse(info.GetID()))
	}
	if info.AppID != nil {
		c.SetAppID(uuid.MustParse(info.GetAppID()))
	}
	if info.UserID != nil {
		c.SetUserID(uuid.MustParse(info.GetUserID()))
	}
	if info.Resource != nil {
		c.SetResource(info.GetResource())
	}
	if info.Method != nil {
		c.SetMethod(info.GetMethod())
	}
	if info.Allowed != nil {
		c.SetAllowed(info.GetAllowed())
	}
	return c
}

func UpdateSet(u *ent.AuthHistoryUpdateOne, in *npool.HistoryReq) *ent.AuthHistoryUpdateOne {
	return u
}

//nolint:nolintlint,gocyclo
func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AuthHistoryQuery, error) {
	stm := cli.AuthHistory.Query()

	if conds == nil {
		return stm, nil
	}

	if conds.ID != nil {
		id, err := uuid.Parse(conds.GetID().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(authhistory.ID(id))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}

	if conds.AppID != nil {
		appID, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(authhistory.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}

	if conds.UserID != nil {
		userID, err := uuid.Parse(conds.GetUserID().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(authhistory.UserID(userID))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}

	if conds.Resource != nil {
		switch conds.GetResource().GetOp() {
		case cruder.EQ:
			stm.Where(authhistory.Resource(conds.GetResource().GetValue()))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}

	if conds.Method != nil {
		switch conds.GetMethod().GetOp() {
		case cruder.EQ:
			stm.Where(authhistory.Method(conds.GetMethod().GetValue()))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}

	if conds.Allowed != nil {
		switch conds.GetAllowed().GetOp() {
		case cruder.EQ:
			stm.Where(authhistory.Allowed(conds.GetAllowed().GetValue()))
		default:
			return nil, fmt.Errorf("invalid app field")
		}
	}

	return stm, nil
}
