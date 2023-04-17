package kyc

import (
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/kyc"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/kyc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

func CreateSet(c *ent.KycCreate, info *npool.KycReq) *ent.KycCreate {
	if info.ID != nil {
		c.SetID(uuid.MustParse(info.GetID()))
	}
	if info.AppID != nil {
		c.SetAppID(uuid.MustParse(info.GetAppID()))
	}
	if info.UserID != nil {
		c.SetUserID(uuid.MustParse(info.GetUserID()))
	}
	if info.DocumentType != nil {
		c.SetDocumentType(info.GetDocumentType().String())
	}
	if info.IDNumber != nil {
		c.SetIDNumber(info.GetIDNumber())
	}
	if info.FrontImg != nil {
		c.SetFrontImg(info.GetFrontImg())
	}
	if info.BackImg != nil {
		c.SetBackImg(info.GetBackImg())
	}
	if info.SelfieImg != nil {
		c.SetSelfieImg(info.GetSelfieImg())
	}
	if info.EntityType != nil {
		c.SetEntityType(info.GetEntityType().String())
	}
	if info.ReviewID != nil {
		c.SetReviewID(uuid.MustParse(info.GetReviewID()))
	}
	if info.State != nil {
		c.SetState(info.GetState().String())
	}
	return c
}

func UpdateSet(info *ent.Kyc, in *npool.KycReq) *ent.KycUpdateOne {
	u := info.Update()

	if in.DocumentType != nil {
		u.SetDocumentType(in.GetDocumentType().String())
	}
	if in.IDNumber != nil {
		u.SetIDNumber(in.GetIDNumber())
	}
	if in.FrontImg != nil {
		u.SetFrontImg(in.GetFrontImg())
	}
	if in.BackImg != nil {
		u.SetBackImg(in.GetBackImg())
	}
	if in.SelfieImg != nil {
		u.SetSelfieImg(in.GetSelfieImg())
	}
	if in.EntityType != nil {
		u.SetEntityType(in.GetEntityType().String())
	}
	if in.ReviewID != nil {
		u.SetReviewID(uuid.MustParse(in.GetReviewID()))
	}
	if in.State != nil {
		u.SetState(in.GetState().String())
	}
	return u
}

//nolint
func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.KycQuery, error) {
	stm := cli.Kyc.Query()

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
			stm.Where(kyc.ID(id))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}

	if conds.AppID != nil {
		appID, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return nil, err
		}

		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(kyc.AppID(appID))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}

	if conds.UserID != nil {
		userID, err := uuid.Parse(conds.GetUserID().GetValue())
		if err != nil {
			return nil, err
		}

		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(kyc.UserID(userID))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}

	if conds.ReviewID != nil {
		reviewID, err := uuid.Parse(conds.GetReviewID().GetValue())
		if err != nil {
			return nil, err
		}
		switch conds.GetReviewID().GetOp() {
		case cruder.EQ:
			stm.Where(kyc.ReviewID(reviewID))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}

	if conds.IDNumber != nil {
		switch conds.GetIDNumber().GetOp() {
		case cruder.EQ:
			stm.Where(kyc.IDNumber(conds.GetIDNumber().GetValue()))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}

	if conds.DocumentType != nil {
		switch conds.GetDocumentType().GetOp() {
		case cruder.EQ:
			stm.Where(kyc.DocumentType(conds.GetDocumentType().GetValue()))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}

	if conds.EntityType != nil {
		switch conds.GetEntityType().GetOp() {
		case cruder.EQ:
			stm.Where(kyc.EntityType(conds.GetEntityType().GetValue()))
		default:
			return nil, fmt.Errorf("invalid kyc field")
		}
	}
	return stm, nil
}
