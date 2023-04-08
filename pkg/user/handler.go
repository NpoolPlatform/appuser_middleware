package user

import (
	"context"
	"fmt"
	"net/mail"
	"regexp"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/const"
	mgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID                 *string
	AppID              string
	PhoneNO            *string
	EmailAddress       *string
	ImportedFromAppID  *string
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
	ThirdPartyID       *string
	ThirdPartyUserID   *string
	ThirdPartyUsername *string
	ThirdPartyAvatar   *string
	Banned             *bool
	BanMessage         *string
	RoleIDs            []string
	Kol                *bool
	KolConfirmed       *bool
	ActionCredits      *string
	Account            *string
	AccountType        *basetypes.SignMethod
	Conds              *mgrpb.Conds
	Offset             *int32
	Limit              *int32
	IDs                []string
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

/// Here ID is UserID
func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.ID = id
		return nil
	}
}

func WithAppID(appID string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _, err := uuid.Parse(appID); err != nil {
			return err
		}
		// TODO: check app exist
		h.AppID = appID
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
		h.EmailAddress = emailAddress
		return nil
	}
}

func WithImportedFromAppID(appID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if appID == nil {
			return nil
		}
		if _, err := uuid.Parse(*appID); err != nil {
			return err
		}
		// TODO: check import from app exist
		h.ImportedFromAppID = appID
		return nil
	}
}

func WithUsername(username *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if username == nil {
			return nil
		}
		re := regexp.MustCompile("^[a-zA-Z0-9][[a-zA-Z0-9_-.]{3,32}$")
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

func WithThirdPartyID(thirdPartyID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if thirdPartyID == nil {
			return nil
		}
		if *thirdPartyID == "" {
			return fmt.Errorf("invalid thirdpartyid")
		}
		h.ThirdPartyID = thirdPartyID
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

func WithRoleIDs(roleIDs []string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, id := range roleIDs {
			if _, err := uuid.Parse(id); err != nil {
				return err
			}
		}
		h.RoleIDs = roleIDs
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
		if confirmed == nil {
			return nil
		}
		h.KolConfirmed = confirmed
		return nil
	}
}

func WithActionCredits(credits *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if credits == nil {
			return nil
		}
		if _, err := decimal.NewFromString(*credits); err != nil {
			return fmt.Errorf("invalid credits")
		}
		h.ActionCredits = credits
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

func WithConds(conds *mgrpb.Conds, offset, limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return fmt.Errorf("invalid conds")
		}

		if conds.ID != nil {
			if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
				return err
			}
		}
		if conds.AppID != nil {
			if _, err := uuid.Parse(conds.GetAppID().GetValue()); err != nil {
				return err
			}
		}

		h.Conds = conds
		h.Offset = &offset
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = &limit

		return nil
	}
}

func WithIDs(ids []string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, id := range ids {
			if _, err := uuid.Parse(id); err != nil {
				return err
			}
		}

		h.IDs = ids
		return nil
	}
}
