package user

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-middleware/pkg/encrypt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	roleusercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/role/user"
	subscribercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/subscriber"
	usercrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user"
	userthirdpartycrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user/3rdparty"
	userctrlcrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user/control"
	userextracrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user/extra"
	usersecretcrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/user/secret"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) account() (string, error) {
	switch *h.AccountType {
	case basetypes.SignMethod_Email:
		if h.EmailAddress == nil {
			return "", fmt.Errorf("invalid email address")
		}
		return *h.EmailAddress, nil
	case basetypes.SignMethod_Mobile:
		if h.PhoneNO == nil {
			return "", fmt.Errorf("invalid phone no")
		}
		return *h.PhoneNO, nil
	case basetypes.SignMethod_Twitter:
		fallthrough
	case basetypes.SignMethod_Github:
		fallthrough
	case basetypes.SignMethod_Facebook:
		fallthrough
	case basetypes.SignMethod_Linkedin:
		fallthrough
	case basetypes.SignMethod_Wechat:
		fallthrough
	case basetypes.SignMethod_Google:
		if h.ThirdPartyUserID == nil {
			return "", fmt.Errorf("invalid thirdpartyuserid")
		}
		return *h.ThirdPartyUserID, nil
	}
	return "", fmt.Errorf("invalid accounttype")
}

func (h *createHandler) createAppUser(ctx context.Context, tx *ent.Tx) error {
	if h.ID == nil {
		return fmt.Errorf("invalid id")
	}
	switch *h.AccountType {
	case basetypes.SignMethod_Twitter:
	case basetypes.SignMethod_Github:
	case basetypes.SignMethod_Facebook:
	case basetypes.SignMethod_Linkedin:
	case basetypes.SignMethod_Wechat:
	case basetypes.SignMethod_Google:
	default:
		if h.PhoneNO == nil && h.EmailAddress == nil {
			return fmt.Errorf("invalid account")
		}
	}

	if _, err := usercrud.CreateSet(
		tx.AppUser.Create(),
		&usercrud.Req{
			ID:            h.ID,
			AppID:         &h.AppID,
			PhoneNO:       h.PhoneNO,
			EmailAddress:  h.EmailAddress,
			ImportFromApp: h.ImportFromAppID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createAppUserExtra(ctx context.Context, tx *ent.Tx) error {
	if _, err := userextracrud.CreateSet(
		tx.AppUserExtra.Create(),
		&userextracrud.Req{
			AppID:  &h.AppID,
			UserID: h.ID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createAppUserControl(ctx context.Context, tx *ent.Tx) error {
	if _, err := userctrlcrud.CreateSet(
		tx.AppUserControl.Create(),
		&userctrlcrud.Req{
			AppID:  &h.AppID,
			UserID: h.ID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createAppUserSecret(ctx context.Context, tx *ent.Tx) error {
	if h.PasswordHash == nil {
		return fmt.Errorf("invalid password")
	}

	saltStr := encrypt.Salt()
	pwdStr, err := encrypt.EncryptWithSalt(*h.PasswordHash, saltStr)
	if err != nil {
		return err
	}

	if _, err := usersecretcrud.CreateSet(
		tx.AppUserSecret.Create(),
		&usersecretcrud.Req{
			AppID:        &h.AppID,
			UserID:       h.ID,
			PasswordHash: &pwdStr,
			Salt:         &saltStr,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *createHandler) createAppUserThirdParty(ctx context.Context, tx *ent.Tx) error {
	switch *h.AccountType {
	case basetypes.SignMethod_Twitter:
	case basetypes.SignMethod_Github:
	case basetypes.SignMethod_Facebook:
	case basetypes.SignMethod_Linkedin:
	case basetypes.SignMethod_Wechat:
	case basetypes.SignMethod_Google:
	default:
		return nil
	}
	if h.ThirdPartyID == nil {
		return fmt.Errorf("thirdPartyID is empry")
	}

	if _, err := userthirdpartycrud.CreateSet(
		tx.AppUserThirdParty.Create(),
		&userthirdpartycrud.Req{
			AppID:              &h.AppID,
			UserID:             h.ID,
			ThirdPartyID:       h.ThirdPartyID,
			ThirdPartyUserID:   h.ThirdPartyUserID,
			ThirdPartyUsername: h.ThirdPartyUsername,
			ThirdPartyAvatar:   h.ThirdPartyAvatar,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *createHandler) createAppRoleUser(ctx context.Context, tx *ent.Tx) error {
	if len(h.RoleIDs) == 0 {
		return nil
	}

	bulk := make([]*ent.AppRoleUserCreate, len(h.RoleIDs))
	for i, roleID := range h.RoleIDs {
		_roleID := roleID
		bulk[i] = roleusercrud.CreateSet(
			tx.AppRoleUser.Create(),
			&roleusercrud.Req{
				AppID:  &h.AppID,
				RoleID: &_roleID,
				UserID: h.ID,
			})
	}
	if _, err := tx.
		AppRoleUser.
		CreateBulk(bulk...).
		Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *createHandler) updateSubscriber(ctx context.Context, tx *ent.Tx) error {
	if h.EmailAddress == nil {
		return nil
	}

	stm, err := subscribercrud.SetQueryConds(
		tx.Subscriber.Query(),
		&subscribercrud.Conds{
			AppID:        &cruder.Cond{Op: cruder.EQ, Val: h.AppID},
			EmailAddress: &cruder.Cond{Op: cruder.EQ, Val: *h.EmailAddress},
		},
	)
	if err != nil {
		return err
	}
	sub, err := stm.Only(ctx)
	if ent.IsNotFound(err) {
		return nil
	}
	if err != nil {
		return err
	}

	if _, err := tx.
		Subscriber.
		UpdateOneID(sub.ID).
		SetRegistered(true).
		Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) CreateUser(ctx context.Context) (info *npool.User, err error) {
	handler := &createHandler{
		Handler: h,
	}

	account, err := handler.account()
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateUser, *h.AccountType, account)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	if err := h.checkAccountExist(ctx); err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createAppUser(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createAppUserExtra(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createAppUserControl(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createAppUserSecret(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createAppUserThirdParty(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createAppRoleUser(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updateSubscriber(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetUser(ctx)
}
