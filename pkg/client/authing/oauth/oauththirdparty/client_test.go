package oauththirdparty

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

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/oauth/oauththirdparty"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/appuser-middleware/pkg/testinit"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
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
	ret = npool.OAuthThirdParty{
		ID:             uuid.NewString(),
		ClientName:     basetypes.SignMethod_Wechat,
		ClientNameStr:  basetypes.SignMethod_Wechat.String(),
		ClientTag:      "Wechat",
		ClientOAuthURL: "https://wechat.com/login/oauth/authorize",
		ClientLogoURL:  "wechat",
		ResponseType:   "code",
		Scope:          "user:email",
	}
)

func createOAuthThirdParty(t *testing.T) {
	var (
		id  = ret.ID
		req = npool.OAuthThirdPartyReq{
			ID:             &id,
			ClientName:     &ret.ClientName,
			ClientTag:      &ret.ClientTag,
			ClientLogoURL:  &ret.ClientLogoURL,
			ClientOAuthURL: &ret.ClientOAuthURL,
			ResponseType:   &ret.ResponseType,
			Scope:          &ret.Scope,
		}
	)

	info, err := CreateOAuthThirdParty(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateOAuthThirdParty(t *testing.T) {
	var (
		clientName     = basetypes.SignMethod_Linkedin
		clientNameStr  = basetypes.SignMethod_Linkedin.String()
		clientTag      = "linkedin"
		clientLogoURL  = "linkedin"
		clientOAuthURL = "https://accounts.linkedin.com/o/oauth2/v2/auth"
		scope          = "email"

		req = npool.OAuthThirdPartyReq{
			ID:             &ret.ID,
			ClientName:     &clientName,
			ClientTag:      &clientTag,
			ClientLogoURL:  &clientLogoURL,
			ClientOAuthURL: &clientOAuthURL,
			ResponseType:   &ret.ResponseType,
			Scope:          &scope,
		}
	)

	info, err := UpdateOAuthThirdParty(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.ClientName = clientName
		ret.ClientNameStr = clientNameStr
		ret.ClientTag = clientTag
		ret.ClientLogoURL = clientLogoURL
		ret.ClientOAuthURL = clientOAuthURL
		ret.Scope = scope
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
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
		ClientTag: &basetypes.StringVal{Op: cruder.EQ, Value: ret.ClientTag},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteOAuthThirdParty(t *testing.T) {
	info, err := DeleteOAuthThirdParty(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetOAuthThirdParty(context.Background(), ret.ID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestOAuthThirdParty(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createOAuthThirdParty", createOAuthThirdParty)
	t.Run("updateOAuthThirdParty", updateOAuthThirdParty)
	t.Run("getOAuthThirdParty", getOAuthThirdParty)
	t.Run("getOAuthThirdParties", getOAuthThirdParties)
	t.Run("deleteOAuthThirdParty", deleteOAuthThirdParty)
}
