package extra

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	entappuserextra "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID            *uuid.UUID
	AppID         *uuid.UUID
	UserID        *uuid.UUID
	FirstName     *string
	LastName      *string
	Organization  *string
	IDNumber      *string
	PostalCode    *string
	Age           *uint32
	Birthday      *uint32
	Avatar        *string
	Username      *string
	Gender        *string
	AddressFields []string
	ActionCredits *decimal.Decimal
}

func CreateSet(c *ent.AppUserExtraCreate, req *Req) *ent.AppUserExtraCreate { //nolint
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.FirstName != nil {
		c.SetFirstName(*req.FirstName)
	}
	if req.LastName != nil {
		c.SetLastName(*req.LastName)
	}
	if req.Organization != nil {
		c.SetOrganization(*req.Organization)
	}
	if req.IDNumber != nil {
		c.SetIDNumber(*req.IDNumber)
	}
	if req.PostalCode != nil {
		c.SetPostalCode(*req.PostalCode)
	}
	if req.Age != nil {
		c.SetAge(*req.Age)
	}
	if req.Birthday != nil {
		c.SetBirthday(*req.Birthday)
	}
	if req.Avatar != nil {
		c.SetAvatar(*req.Avatar)
	}
	if req.Username != nil {
		c.SetUsername(*req.Username)
	}
	if req.Gender != nil {
		c.SetGender(*req.Gender)
	}
	if req.AddressFields != nil {
		c.SetAddressFields(req.AddressFields)
	}
	if req.ActionCredits != nil {
		c.SetActionCredits(*req.ActionCredits)
	}
	return c
}

func UpdateSet(ctx context.Context, u *ent.AppUserExtraUpdateOne, req *Req) (*ent.AppUserExtraUpdateOne, error) {
	if req.Username != nil {
		u.SetUsername(*req.Username)
	}
	if req.FirstName != nil {
		u.SetFirstName(*req.FirstName)
	}
	if req.LastName != nil {
		u.SetLastName(*req.LastName)
	}
	if req.AddressFields != nil {
		u.SetAddressFields(req.AddressFields)
	}
	if req.Gender != nil {
		u.SetGender(*req.Gender)
	}
	if req.PostalCode != nil {
		u.SetPostalCode(*req.PostalCode)
	}
	if req.IDNumber != nil {
		u.SetIDNumber(*req.IDNumber)
	}
	if req.Organization != nil {
		u.SetOrganization(*req.Organization)
	}
	if req.Age != nil {
		u.SetAge(*req.Age)
	}
	if req.Birthday != nil {
		u.SetBirthday(*req.Birthday)
	}
	if req.Avatar != nil {
		u.SetAvatar(*req.Avatar)
	}
	if req.LastName != nil {
		u.SetLastName(*req.LastName)
	}
	if req.ActionCredits != nil {
		oldCredits, err := u.Mutation().OldActionCredits(ctx)
		if err != nil {
			return nil, err
		}
		credits := oldCredits.Add(*req.ActionCredits)
		u.SetActionCredits(credits)
	}

	return u, nil
}

type Conds struct {
	ID       *cruder.Cond
	AppID    *cruder.Cond
	UserID   *cruder.Cond
	IDNumber *cruder.Cond
}

//nolint:nolintlint,gocyclo
func SetQueryConds(q *ent.AppUserExtraQuery, conds *Conds) (*ent.AppUserExtraQuery, error) {
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
			q.Where(entappuserextra.ID(id))
		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappuserextra.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entappuserextra.UserID(id))
		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	if conds.IDNumber != nil {
		idNumber, ok := conds.IDNumber.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid id number")
		}
		switch conds.IDNumber.Op {
		case cruder.EQ:
			q.Where(entappuserextra.IDNumber(idNumber))
		default:
			return nil, fmt.Errorf("invalid appuserextra field")
		}
	}
	return q, nil
}
