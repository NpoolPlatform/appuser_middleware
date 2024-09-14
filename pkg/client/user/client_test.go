package user

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	appoauththirdpartymwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/oauth/appoauththirdparty"
	appoauththirdpartymwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/oauth/appoauththirdparty"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/appuser-middleware/pkg/testinit"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	appmwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"

	oauththirdpartymwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/oauth/oauththirdparty"
	oauththirdpartymwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/oauth/oauththirdparty"
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
	addressFields     = []string{uuid.NewString()}
	addressFieldsS, _ = json.Marshal(addressFields)
	appID             = uuid.NewString()
	thirdPartyID      = uuid.NewString()
	ret               = npool.User{
		EntID:               uuid.NewString(),
		AppID:               appID,
		EmailAddress:        "aaa@aaa.aaa",
		PhoneNO:             "+8613612203133",
		ImportedFromAppID:   uuid.NewString(),
		Username:            "amwnrekadsf.are-",
		AddressFieldsString: string(addressFieldsS),
		AddressFields:       addressFields,
		Gender:              uuid.NewString(),
		PostalCode:          uuid.NewString(),
		Age:                 100,
		Birthday:            uint32(time.Now().Unix()),
		Avatar:              uuid.NewString(),
		Organization:        uuid.NewString(),
		FirstName:           uuid.NewString(),
		LastName:            uuid.NewString(),
		IDNumber:            uuid.NewString(),
		GoogleAuthVerified:  true,
		SigninVerifyType:    basetypes.SignMethod_Email,
		SigninVerifyTypeStr: basetypes.SignMethod_Email.String(),
		GoogleSecret:        appID,
		HasGoogleSecret:     true,
		Roles:               []string{""},
		Banned:              true,
		BanMessage:          uuid.NewString(),
	}

	thirdRet = appoauththirdpartymwpb.OAuthThirdParty{
		EntID:          uuid.NewString(),
		AppID:          appID,
		ThirdPartyID:   thirdPartyID,
		ClientID:       "123123123123",
		ClientSecret:   uuid.NewString(),
		CallbackURL:    "http://localhost:8011/oauth/callback",
		ClientName:     basetypes.SignMethod_Twitter,
		ClientNameStr:  basetypes.SignMethod_Twitter.String(),
		ClientTag:      "twitter",
		ClientOAuthURL: "https://twitter.com/login/oauth/authorize",
		ClientLogoURL:  "twitter",
		ResponseType:   "code",
		Scope:          "user:email",
	}
)

func setupUser(t *testing.T) func(*testing.T) {
	app1, err := appmwcli.CreateApp(
		context.Background(),
		&appmwpb.AppReq{
			EntID:     &ret.AppID,
			CreatedBy: &ret.EntID,
			Name:      &ret.AppID,
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, app1)

	app2, err := appmwcli.CreateApp(
		context.Background(),
		&appmwpb.AppReq{
			EntID:     &ret.ImportedFromAppID,
			CreatedBy: &ret.EntID,
			Name:      &ret.ImportedFromAppID,
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, app2)

	oauth1, err := oauththirdpartymwcli.CreateOAuthThirdParty(
		context.Background(),
		&oauththirdpartymwpb.OAuthThirdPartyReq{
			EntID:          &thirdRet.ThirdPartyID,
			ClientName:     &thirdRet.ClientName,
			ClientTag:      &thirdRet.ClientTag,
			ClientLogoURL:  &thirdRet.ClientLogoURL,
			ClientOAuthURL: &thirdRet.ClientOAuthURL,
			ResponseType:   &thirdRet.ResponseType,
			Scope:          &thirdRet.Scope,
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, oauth1)
	thirdRet.ThirdPartyID = oauth1.EntID

	appoauth1, err := appoauththirdpartymwcli.CreateOAuthThirdParty(
		context.Background(),
		&appoauththirdpartymwpb.OAuthThirdPartyReq{
			EntID:        &thirdRet.EntID,
			AppID:        &thirdRet.AppID,
			ClientID:     &thirdRet.ClientID,
			ClientSecret: &thirdRet.ClientSecret,
			CallbackURL:  &thirdRet.CallbackURL,
			ThirdPartyID: &thirdRet.ThirdPartyID,
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, appoauth1)

	return func(*testing.T) {
		_, _ = appmwcli.DeleteApp(context.Background(), app1.ID)
		_, _ = appmwcli.DeleteApp(context.Background(), app2.ID)
		if thirdPartyID == oauth1.EntID {
			_, _ = oauththirdpartymwcli.DeleteOAuthThirdParty(context.Background(), oauth1.ID)
		}
		_, _ = appoauththirdpartymwcli.DeleteOAuthThirdParty(context.Background(), appoauth1.ID)
	}
}

func creatUser(t *testing.T) {
	ret.PhoneNO = fmt.Sprintf("+86%v", rand.Intn(100000000)+10000)           //nolint
	ret.EmailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+30000) //nolint
	ret.ImportedFromAppName = ret.ImportedFromAppID
	var (
		id                = ret.EntID
		appID             = ret.AppID
		importedFromAppID = ret.ImportedFromAppID
		strVal            = "AAA"
		req               = npool.UserReq{
			EntID:              &id,
			AppID:              &appID,
			EmailAddress:       &ret.EmailAddress,
			PhoneNO:            &ret.PhoneNO,
			ImportedFromAppID:  &importedFromAppID,
			Username:           &ret.Username,
			AddressFields:      addressFields,
			Gender:             &ret.Gender,
			PostalCode:         &ret.PostalCode,
			Age:                &ret.Age,
			Birthday:           &ret.Birthday,
			Avatar:             &ret.Avatar,
			Organization:       &ret.Organization,
			FirstName:          &ret.FirstName,
			LastName:           &ret.LastName,
			IDNumber:           &ret.IDNumber,
			GoogleAuthVerified: &ret.GoogleAuthVerified,
			SigninVerifyType:   &ret.SigninVerifyType,
			PasswordHash:       &strVal,
			GoogleSecret:       &appID,
			ThirdPartyID:       &thirdRet.ThirdPartyID,
			ThirdPartyUserID:   &strVal,
			ThirdPartyUsername: &strVal,
			ThirdPartyAvatar:   &strVal,
			Banned:             &ret.Banned,
			BanMessage:         &ret.BanMessage,
		}
		ret1 = npool.User{
			EntID:               ret.EntID,
			AppID:               ret.AppID,
			EmailAddress:        ret.EmailAddress,
			PhoneNO:             ret.PhoneNO,
			ImportedFromAppID:   ret.ImportedFromAppID,
			ImportedFromAppName: ret.ImportedFromAppName,
			AddressFieldsString: "[]",
			AddressFields:       nil,
			SigninVerifyTypeStr: basetypes.SignMethod_Email.String(),
			SigninVerifyType:    basetypes.SignMethod_Email,
		}
	)

	info, err := CreateUser(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.OAuthThirdParties = info.OAuthThirdParties
		ret.CreatedAt = info.CreatedAt
		ret1.CreatedAt = info.CreatedAt
		ret1.OAuthThirdParties = info.OAuthThirdParties
		ret.ID = info.ID
		ret1.ID = info.ID
		assert.Equal(t, info, &ret1)
	}
}

func updateUser(t *testing.T) {
	ret.PhoneNO = fmt.Sprintf("+86%v", rand.Intn(100000000)+10000)           //nolint
	ret.EmailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+40000) //nolint
	ret.BanAppUserID = ret.EntID
	var (
		appID  = ret.AppID
		strVal = "AAA"

		req = npool.UserReq{
			ID:                 &ret.ID,
			AppID:              &ret.AppID,
			EmailAddress:       &ret.EmailAddress,
			PhoneNO:            &ret.PhoneNO,
			ImportedFromAppID:  &ret.ImportedFromAppID,
			Username:           &ret.Username,
			AddressFields:      addressFields,
			Gender:             &ret.Gender,
			PostalCode:         &ret.PostalCode,
			Age:                &ret.Age,
			Birthday:           &ret.Birthday,
			Avatar:             &ret.Avatar,
			Organization:       &ret.Organization,
			FirstName:          &ret.FirstName,
			LastName:           &ret.LastName,
			IDNumber:           &ret.IDNumber,
			GoogleAuthVerified: &ret.GoogleAuthVerified,
			SigninVerifyType:   &ret.SigninVerifyType,
			PasswordHash:       &strVal,
			GoogleSecret:       &appID,
			ThirdPartyID:       &thirdRet.ThirdPartyID,
			ThirdPartyUserID:   &strVal,
			ThirdPartyUsername: &strVal,
			ThirdPartyAvatar:   &strVal,
			Banned:             &ret.Banned,
			BanMessage:         &ret.BanMessage,
		}
	)

	info, err := UpdateUser(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.OAuthThirdParties = info.OAuthThirdParties
		ret.Roles = info.Roles
		assert.Equal(t, info, &ret)
	}
}

func getUser(t *testing.T) {
	info, err := GetUser(context.Background(), ret.AppID, ret.EntID)
	if assert.Nil(t, err) {
		ret.OAuthThirdParties = info.OAuthThirdParties
		assert.Equal(t, info, &ret)
	}
}

func getUsers(t *testing.T) {
	infos, _, err := GetUsers(context.Background(), &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteUser(t *testing.T) {
	info, err := DeleteUser(context.Background(), ret.AppID, ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetUser(context.Background(), ret.AppID, ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestUser(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	teardown := setupUser(t)
	defer teardown(t)

	t.Run("creatUser", creatUser)
	t.Run("updateUser", updateUser)
	t.Run("getUser", getUser)
	t.Run("getUsers", getUsers)
	t.Run("deleteUser", deleteUser)
}
