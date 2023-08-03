package appoauththirdparty

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/oauth/appoauththirdparty"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/appuser-middleware/pkg/testinit"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	oauththirdpartymwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/authing/oauth/oauththirdparty"
	oauththirdpartymwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/oauth/oauththirdparty"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	id           = uuid.NewString()
	appID        = uuid.NewString()
	thirdPartyID = uuid.NewString()
	ret          = npool.OAuthThirdParty{
		ID:             id,
		AppID:          appID,
		ThirdPartyID:   thirdPartyID,
		ClientID:       "123123123123",
		ClientSecret:   uuid.NewString(),
		CallbackURL:    "http://localhost:8011/oauth/callback",
		ClientName:     basetypes.SignMethod_Linkedin,
		ClientNameStr:  basetypes.SignMethod_Linkedin.String(),
		ClientTag:      "linkedin",
		ClientOAuthURL: "https://linkedin.com/login/oauth/authorize",
		ClientLogoURL:  "linkedin",
		ResponseType:   "code",
		Scope:          "user:email",
	}
)

func setupOAuthThirdParty(t *testing.T) func(*testing.T) {
	app1, err := appmwcli.CreateApp(
		context.Background(),
		&appmwpb.AppReq{
			ID:        &ret.AppID,
			CreatedBy: &ret.ID,
			Name:      &ret.AppID,
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, app1)

	oauth1, err := oauththirdpartymwcli.CreateOAuthThirdParty(
		context.Background(),
		&oauththirdpartymwpb.OAuthThirdPartyReq{
			ID:             &ret.ThirdPartyID,
			ClientName:     &ret.ClientName,
			ClientTag:      &ret.ClientTag,
			ClientLogoURL:  &ret.ClientLogoURL,
			ClientOAuthURL: &ret.ClientOAuthURL,
			ResponseType:   &ret.ResponseType,
			Scope:          &ret.Scope,
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, oauth1)
	ret.ThirdPartyID = oauth1.ID

	return func(*testing.T) {
		_, _ = appmwcli.DeleteApp(context.Background(), ret.AppID)
		if thirdPartyID == oauth1.ID {
			_, _ = oauththirdpartymwcli.DeleteOAuthThirdParty(context.Background(), ret.ThirdPartyID)
		}
	}
}

func createOAuthThirdParty(t *testing.T) {
	var (
		req = npool.OAuthThirdPartyReq{
			ID:           &ret.ID,
			AppID:        &ret.AppID,
			ThirdPartyID: &ret.ThirdPartyID,
			ClientID:     &ret.ClientID,
			ClientSecret: &ret.ClientSecret,
			CallbackURL:  &ret.CallbackURL,
		}
	)

	info, err := CreateOAuthThirdParty(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.ThirdPartyID = info.ThirdPartyID
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ClientSecret = info.ClientSecret
		assert.Equal(t, info, &ret)
	}
}

func updateOAuthThirdParty(t *testing.T) {
	var (
		clientID     = "147147147"
		clientSecret = uuid.NewString()
		callbackURL  = "http://localhost:8091"

		req = npool.OAuthThirdPartyReq{
			ID:           &ret.ID,
			AppID:        &ret.AppID,
			ClientID:     &clientID,
			ClientSecret: &clientSecret,
			CallbackURL:  &callbackURL,
		}
	)

	if ret.ID == id {
		info, err := UpdateOAuthThirdParty(context.Background(), &req)
		if assert.Nil(t, err) {
			ret.ClientID = clientID
			ret.CallbackURL = callbackURL
			ret.ClientSecret = info.ClientSecret
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getOAuthThirdParty(t *testing.T) {
	info, err := GetOAuthThirdParty(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getOAuthThirdParties(t *testing.T) {
	infos, _, err := GetOAuthThirdParties(context.Background(), &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteOAuthThirdParty(t *testing.T) {
	if ret.ID == id {
		info, err := DeleteOAuthThirdParty(context.Background(), ret.ID)
		if assert.Nil(t, err) {
			assert.Equal(t, info, &ret)
		}

		info, err = GetOAuthThirdParty(context.Background(), ret.ID)
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestUser(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	teardown := setupOAuthThirdParty(t)
	defer teardown(t)

	t.Run("createOAuthThirdParty", createOAuthThirdParty)
	t.Run("updateOAuthThirdParty", updateOAuthThirdParty)
	t.Run("getOAuthThirdParty", getOAuthThirdParty)
	t.Run("getOAuthThirdParties", getOAuthThirdParties)
	t.Run("deleteOAuthThirdParty", deleteOAuthThirdParty)
}
