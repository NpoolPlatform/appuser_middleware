package kyc

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/const"
	usercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user"
	app "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID           *uuid.UUID
	AppID        uuid.UUID
	UserID       uuid.UUID
	DocumentType *basetypes.KycDocumentType
	IDNumber     *string
	FromImg      *string
	BackImg      *string
	SelfieImg    *string
	EntityType   *basetypes.KycEntityType
	ReviewID     *uuid.UUID
	State        *basetypes.KycState
	// Conds        *kyccrud.Conds
	Offset int32
	Limit  int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithAppID(id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := app.NewHandler(
			ctx,
			app.WithID(&id),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistApp(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid app")
		}
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.AppID = _id
		return nil
	}
}

func WithUserID(id string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		// TODO: check user exist
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.UserID = _id
		return nil
	}
}

func WithDocumentType(docType *basetypes.KycDocumentType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if docType == nil {
			return nil
		}
		switch *docType {
		case basetypes.KycDocumentType_IDCard:
		case basetypes.KycDocumentType_DriverLicense:
		case basetypes.KycDocumentType_Passport:
		default:
			return fmt.Errorf("invalid document type")
		}
		h.DocumentType = docType
		return nil
	}
}

func WithIDNumber(idNumber *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if idNumber == nil {
			return nil
		}
		const leastIDNumberLen = 8
		if len(*idNumber) < leastIDNumberLen {
			return fmt.Errorf("invalid id number")
		}
		h.IDNumber = idNumber
		return nil
	}
}

func WithFrontImg(img *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.FrontImg = img
		return nil
	}
}

func WithBackImg(img *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.BackImg = img
		return nil
	}
}

func WithSelfieImg(img *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.SelfieImg = img
		return nil
	}
}

func WithEntityType(entType *basetypes.KycEntityType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if entType == nil {
			return nil
		}
		switch *entType {
		case basetypes.KycEntityType_Individual:
		case basetypes.KycEntityType_Organization:
		default:
			return fmt.Errorf("invalid entity type")
		}
		h.EntityType = entType
		return nil
	}
}

func WithReviewID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ReviewID = &_id
		return nil
	}
}

/*
func WithState(state *basetypes.KycState) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			return nil
		}
		switch *state {
		case basetypes.KycStaet_Approved:
		case basetypes.KycStaet_Reviewing:
		case basetypes.KycStaet_Rejected:
		default:
			return fmt.Errorf("invalid state")
		}
		h.State = state
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &usercrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: id}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.PhoneNO != nil {
			h.Conds.PhoneNO = &cruder.Cond{
				Op:  conds.GetPhoneNO().GetOp(),
				Val: conds.GetPhoneNO().GetValue(),
			}
		}
		if conds.EmailAddress != nil {
			h.Conds.EmailAddress = &cruder.Cond{
				Op:  conds.GetEmailAddress().GetOp(),
				Val: conds.GetEmailAddress().GetValue(),
			}
		}
		if conds.ImportFromApp != nil {
			id, err := uuid.Parse(conds.GetImportFromApp().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ImportFromApp = &cruder.Cond{
				Op:  conds.GetImportFromApp().GetOp(),
				Val: id,
			}
		}
		if len(conds.GetIDs().GetValue()) > 0 {
			ids := []uuid.UUID{}
			for _, id := range conds.GetIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.IDs = &cruder.Cond{Op: conds.GetIDs().GetOp(), Val: ids}
		}
		return nil
	}
}
*/

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
