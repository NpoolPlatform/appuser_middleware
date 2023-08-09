package user

import (
	"context"
	"fmt"
	"net/mail"
	"regexp"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/const"
	usercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user"
	app "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	oauththirdparty "github.com/NpoolPlatform/appuser-middleware/pkg/mw/oauth/oauththirdparty"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID                 *uuid.UUID
	AppID              uuid.UUID
	PhoneNO            *string
	EmailAddress       *string
	ImportFromAppID    *uuid.UUID
	Username           *string
	AddressFields      []string
	Gender             *string
	PostalCode         *string
	Age                *uint32
	Birthday           *uint32
	Avatar             *string
	Organization       *string
	FirstName          *string
	LastName           *string
	IDNumber           *string
	SigninVerifyType   *basetypes.SignMethod
	GoogleAuthVerified *bool
	PasswordHash       *string
	GoogleSecret       *string
	ThirdPartyID       *uuid.UUID
	ThirdPartyUserID   *string
	ThirdPartyUsername *string
	ThirdPartyAvatar   *string
	Banned             *bool
	BanMessage         *string
	RoleIDs            []uuid.UUID
	Kol                *bool
	KolConfirmed       *bool
	ActionCredits      *decimal.Decimal
	SelectedLangID     *uuid.UUID
	Account            *string
	AccountType        *basetypes.SignMethod
	Conds              *usercrud.Conds
	Offset             int32
	Limit              int32
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

// Here ID is UserID
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

func validateEmailAddress(emailAddress string) error {
	if _, err := mail.ParseAddress(emailAddress); err != nil {
		return err
	}
	return nil
}

func validatePhoneNO(phoneNO string) error {
	re := regexp.MustCompile(
		`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[` +
			`\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?)` +
			`{0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)` +
			`[\-\.\ \\\/]?(\d+))?$`,
	)
	if !re.MatchString(phoneNO) {
		return fmt.Errorf("invalid phone no")
	}

	return nil
}

func WithPhoneNO(phoneNO *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if phoneNO == nil {
			return nil
		}
		if err := validatePhoneNO(*phoneNO); err != nil {
			return err
		}
		accountType := basetypes.SignMethod_Mobile
		h.Account = phoneNO
		h.AccountType = &accountType
		h.PhoneNO = phoneNO
		return nil
	}
}

func WithEmailAddress(emailAddress *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if emailAddress == nil {
			return nil
		}
		if err := validateEmailAddress(*emailAddress); err != nil {
			return err
		}

		accountType := basetypes.SignMethod_Email
		h.Account = emailAddress
		h.AccountType = &accountType
		h.EmailAddress = emailAddress
		return nil
	}
}

func WithImportFromAppID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		handler, err := app.NewHandler(
			ctx,
			app.WithID(id),
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
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ImportFromAppID = &_id
		return nil
	}
}

func WithUsername(username *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if username == nil {
			return nil
		}
		re := regexp.MustCompile("^[a-zA-Z0-9\u3040-\u31ff][[a-zA-Z0-9_\\-\\.\u3040-\u31ff]{3,32}$") //nolint
		if !re.MatchString(*username) {
			return fmt.Errorf("invalid username")
		}
		h.Username = username
		return nil
	}
}

func WithAddressFields(addressFields []string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AddressFields = addressFields
		return nil
	}
}

func WithGender(gender *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if gender == nil {
			return nil
		}
		if *gender == "" {
			return fmt.Errorf("invalid gender")
		}
		h.Gender = gender
		return nil
	}
}

func WithPostalCode(postalCode *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if postalCode == nil {
			return nil
		}
		if *postalCode == "" {
			return fmt.Errorf("invalid postalCode")
		}
		h.PostalCode = postalCode
		return nil
	}
}

func WithAge(age *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if age == nil {
			return nil
		}
		h.Age = age
		return nil
	}
}

func WithBirthday(birthday *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if birthday == nil {
			return nil
		}
		h.Birthday = birthday
		return nil
	}
}

func WithAvatar(avatar *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if avatar == nil {
			return nil
		}
		if *avatar == "" {
			return fmt.Errorf("invalid avatar")
		}
		h.Avatar = avatar
		return nil
	}
}

func WithOrganization(organization *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if organization == nil {
			return nil
		}
		if *organization == "" {
			return fmt.Errorf("invalid organization")
		}
		h.Organization = organization
		return nil
	}
}

func WithFirstName(firstName *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if firstName == nil {
			return nil
		}
		if *firstName == "" {
			return fmt.Errorf("invalid firstname")
		}
		h.FirstName = firstName
		return nil
	}
}

func WithLastName(lastName *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if lastName == nil {
			return nil
		}
		if *lastName == "" {
			return fmt.Errorf("invalid lastname")
		}
		h.LastName = lastName
		return nil
	}
}

func WithIDNumber(idNumber *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if idNumber == nil {
			return nil
		}
		if *idNumber == "" {
			return fmt.Errorf("invalid idnumber")
		}
		h.IDNumber = idNumber
		return nil
	}
}

func WithSigninVerifyType(verifyType *basetypes.SignMethod) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if verifyType == nil {
			return nil
		}
		switch *verifyType {
		case basetypes.SignMethod_Email:
		case basetypes.SignMethod_Mobile:
		case basetypes.SignMethod_Google:
		default:
			return fmt.Errorf("invalid sign verify type")
		}
		h.SigninVerifyType = verifyType
		return nil
	}
}

func WithGoogleAuthVerified(verified *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if verified == nil {
			return nil
		}
		h.GoogleAuthVerified = verified
		return nil
	}
}

func WithPasswordHash(passwordHash *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if passwordHash == nil {
			return nil
		}
		if *passwordHash == "" {
			return fmt.Errorf("invalid passwordhash")
		}
		h.PasswordHash = passwordHash
		return nil
	}
}

func WithGoogleSecret(secret *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if secret == nil {
			return nil
		}
		if *secret == "" {
			return fmt.Errorf("invalid google secret")
		}
		h.GoogleSecret = secret
		return nil
	}
}

func WithThirdPartyID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		handler, err := oauththirdparty.NewHandler(
			ctx,
			oauththirdparty.WithID(id),
		)
		if err != nil {
			return err
		}
		thirdParty, err := handler.GetOAuthThirdParty(ctx)
		if err != nil {
			return err
		}
		if thirdParty == nil {
			return fmt.Errorf("invalid oauththirdparty")
		}
		accountType := basetypes.SignMethod(basetypes.SignMethod_value[thirdParty.ClientNameStr])
		h.AccountType = &accountType

		h.ThirdPartyID = &_id
		return nil
	}
}

func WithThirdPartyUserID(userID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if userID == nil {
			return nil
		}
		if *userID == "" {
			return fmt.Errorf("invalid thirdpartyuserid")
		}
		h.Account = userID
		h.ThirdPartyUserID = userID
		return nil
	}
}

func WithThirdPartyUsername(username *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if username == nil {
			return nil
		}
		if *username == "" {
			return fmt.Errorf("invalid thirdpartyusername")
		}
		h.ThirdPartyUsername = username
		return nil
	}
}

func WithThirdPartyAvatar(avatar *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if avatar == nil {
			return nil
		}
		if *avatar == "" {
			return fmt.Errorf("invalid avatar")
		}
		h.ThirdPartyAvatar = avatar
		return nil
	}
}

func WithBanned(banned *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if banned == nil {
			return nil
		}
		h.Banned = banned
		return nil
	}
}

func WithBanMessage(message *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if message == nil {
			return nil
		}
		if *message == "" {
			return fmt.Errorf("invalid message")
		}
		h.BanMessage = message
		return nil
	}
}

func WithRoleIDs(ids []string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if len(ids) == 0 {
			return nil
		}
		_ids := []uuid.UUID{}
		for _, id := range ids {
			_id, err := uuid.Parse(id)
			if err != nil {
				return err
			}
			_ids = append(_ids, _id)
		}
		h.RoleIDs = _ids
		return nil
	}
}

func WithKol(kol *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if kol == nil {
			return nil
		}
		h.Kol = kol
		return nil
	}
}

func WithKolConfirmed(confirmed *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.KolConfirmed = confirmed
		return nil
	}
}

func WithActionCredits(credits *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if credits == nil {
			return nil
		}
		_credits, err := decimal.NewFromString(*credits)
		if err != nil {
			return err
		}
		h.ActionCredits = &_credits
		return nil
	}
}

func WithSelectedLangID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.SelectedLangID = &_id
		return nil
	}
}

func WithAccount(account string, accountType basetypes.SignMethod) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if account == "" {
			return fmt.Errorf("invalid account")
		}

		var err error

		switch accountType {
		case basetypes.SignMethod_Mobile:
			h.PhoneNO = &account
			err = validatePhoneNO(account)
		case basetypes.SignMethod_Email:
			h.EmailAddress = &account
			err = validateEmailAddress(account)
		default:
			return fmt.Errorf("invalid account type")
		}

		if err != nil {
			return err
		}

		h.AccountType = &accountType
		h.Account = &account
		return nil
	}
}

//nolint:gocyclo
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
		if conds.ThirdPartyID != nil {
			id, err := uuid.Parse(conds.GetThirdPartyID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ThirdPartyID = &cruder.Cond{Op: conds.GetThirdPartyID().GetOp(), Val: id}
		}
		if conds.ThirdPartyUserID != nil {
			h.Conds.ThirdPartyUserID = &cruder.Cond{Op: conds.GetThirdPartyUserID().GetOp(), Val: conds.GetThirdPartyUserID().GetValue()}
		}
		return nil
	}
}

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
