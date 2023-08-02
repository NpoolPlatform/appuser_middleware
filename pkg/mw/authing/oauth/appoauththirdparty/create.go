package appoauththirdparty

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/aes"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"

	appoauththirdpartycrud "github.com/NpoolPlatform/appuser-middleware/pkg/crud/authing/oauth/appoauththirdparty"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/oauth/appoauththirdparty"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validtate() error {
	if h.ThirdPartyID == nil {
		return fmt.Errorf("invalid clientsecret")
	}
	if h.ClientID == nil {
		return fmt.Errorf("invalid clientid")
	}
	if h.ClientSecret == nil {
		return fmt.Errorf("invalid clientsecret")
	}
	if h.CallbackURL == nil {
		return fmt.Errorf("invalid callbackurl")
	}
	return nil
}

func (h *Handler) CreateOAuthThirdParty(ctx context.Context) (*npool.OAuthThirdParty, error) {
	handler := &createHandler{
		Handler: h,
	}
	if err := handler.validtate(); err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateAppOAuthThirdParty, h.AppID, h.ThirdPartyID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	oauthHandler, err := NewHandler(
		ctx,
		WithConds(&npool.Conds{
			AppID:        &basetypes.StringVal{Op: cruder.EQ, Value: h.AppID.String()},
			ThirdPartyID: &basetypes.StringVal{Op: cruder.EQ, Value: h.ThirdPartyID.String()},
		}),
	)
	if err != nil {
		return nil, err
	}
	exist, err := oauthHandler.ExistOAuthThirdPartyConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("oauththirdparty exist")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}
	salt, err := aes.NewAesKey(aes.AES256)
	if err != nil {
		return nil, fmt.Errorf("get salt failed")
	}
	clientSecret, err := aes.AesEncrypt([]byte(salt), []byte(*h.ClientSecret))
	if err != nil {
		return nil, fmt.Errorf("encrypt clientSecret failed")
	}
	clientSecretStr := hex.EncodeToString(clientSecret)
	h.ClientSecret = &clientSecretStr

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := appoauththirdpartycrud.CreateSet(
			tx.AppOAuthThirdParty.Create(),
			&appoauththirdpartycrud.Req{
				ID:           h.ID,
				AppID:        &h.AppID,
				ThirdPartyID: h.ThirdPartyID,
				ClientID:     h.ClientID,
				ClientSecret: h.ClientSecret,
				CallbackURL:  h.CallbackURL,
				Salt:         &salt,
			},
		).Save(ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOAuthThirdParty(ctx)
}
