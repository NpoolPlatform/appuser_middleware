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

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/appuser-middleware/pkg/testinit"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

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
	uuidSlice     = []string{uuid.NewString()}
	uuidSliceS, _ = json.Marshal(uuidSlice)
	appID         = uuid.NewString()
	ret           = npool.User{
		ID:                    uuid.NewString(),
		AppID:                 appID,
		EmailAddress:          "aaa@aaa.aaa",
		PhoneNO:               "+8613612203133",
		ImportedFromAppID:     uuid.NewString(),
		Username:              "amwnrekadsf.are-",
		AddressFieldsString:   string(uuidSliceS),
		AddressFields:         uuidSlice,
		Gender:                uuid.NewString(),
		PostalCode:            uuid.NewString(),
		Age:                   100,
		Birthday:              uint32(time.Now().Unix()),
		Avatar:                uuid.NewString(),
		Organization:          uuid.NewString(),
		FirstName:             uuid.NewString(),
		LastName:              uuid.NewString(),
		IDNumber:              uuid.NewString(),
		GoogleAuthVerifiedInt: 1,
		GoogleAuthVerified:    true,
		SigninVerifyType:      basetypes.SignMethod_Email,
		SigninVerifyTypeStr:   basetypes.SignMethod_Email.String(),
		GoogleSecret:          appID,
		HasGoogleSecret:       true,
		Roles:                 []string{""},
		ActionCredits:         "0",
	}
)

func setupUser(t *testing.T) func(*testing.T) {
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

	app2, err := appmwcli.CreateApp(
		context.Background(),
		&appmwpb.AppReq{
			ID:        &ret.ImportedFromAppID,
			CreatedBy: &ret.ID,
			Name:      &ret.ImportedFromAppID,
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, app2)

	return func(*testing.T) {
		_, _ = appmwcli.DeleteApp(context.Background(), ret.AppID)
		_, _ = appmwcli.DeleteApp(context.Background(), ret.ImportedFromAppID)
	}
}

func creatUser(t *testing.T) {
	ret.PhoneNO = fmt.Sprintf("+86%v", rand.Intn(100000000)+10000)           //nolint
	ret.EmailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+10000) //nolint
	ret.ImportedFromAppName = ret.ImportedFromAppID
	var (
		id                = ret.ID
		appID             = ret.AppID
		importedFromAppID = ret.ImportedFromAppID
		strVal            = "AAA"
		req               = npool.UserReq{
			ID:                 &id,
			AppID:              &appID,
			EmailAddress:       &ret.EmailAddress,
			PhoneNO:            &ret.PhoneNO,
			ImportedFromAppID:  &importedFromAppID,
			Username:           &ret.Username,
			AddressFields:      uuidSlice,
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
			ThirdPartyID:       &strVal,
			ThirdPartyUserID:   &strVal,
			ThirdPartyUsername: &strVal,
			ThirdPartyAvatar:   &strVal,
			Banned:             &ret.Banned,
			BanMessage:         &ret.BanMessage,
		}
		ret1 = npool.User{
			ID:                  ret.ID,
			AppID:               ret.AppID,
			EmailAddress:        ret.EmailAddress,
			PhoneNO:             ret.PhoneNO,
			ImportedFromAppID:   ret.ImportedFromAppID,
			ImportedFromAppName: ret.ImportedFromAppName,
			ActionCredits:       ret.ActionCredits,
			AddressFieldsString: "[]",
			AddressFields:       nil,
			SigninVerifyTypeStr: basetypes.SignMethod_Email.String(),
			SigninVerifyType:    basetypes.SignMethod_Email,
		}
	)

	info, err := CreateUser(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret1.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &ret1)
	}
}

func updateUser(t *testing.T) {
	ret.PhoneNO = fmt.Sprintf("+86%v", rand.Intn(100000000)+10000)           //nolint
	ret.EmailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+10000) //nolint
	var (
		appID   = ret.AppID
		strVal  = "AAA"
		credits = "1.234234"

		req = npool.UserReq{
			ID:                 &ret.ID,
			AppID:              &ret.AppID,
			EmailAddress:       &ret.EmailAddress,
			PhoneNO:            &ret.PhoneNO,
			ImportedFromAppID:  &ret.ImportedFromAppID,
			Username:           &ret.Username,
			AddressFields:      uuidSlice,
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
			ThirdPartyID:       &strVal,
			ThirdPartyUserID:   &strVal,
			ThirdPartyUsername: &strVal,
			ThirdPartyAvatar:   &strVal,
			Banned:             &ret.Banned,
			BanMessage:         &ret.BanMessage,
			ActionCredits:      &credits,
		}
	)

	ret.ActionCredits = credits

	info, err := UpdateUser(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.Roles = info.Roles
		assert.Equal(t, info, &ret)
	}
}

func getUser(t *testing.T) {
	info, err := GetUser(context.Background(), ret.AppID, ret.ID)
	if assert.Nil(t, err) {
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

	info, err = GetUser(context.Background(), ret.AppID, ret.ID)
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

	teardown := setupUser(t)
	defer teardown(t)

	t.Run("creatUser", creatUser)
	t.Run("updateUser", updateUser)
	t.Run("getUser", getUser)
	t.Run("getUsers", getUsers)
	t.Run("deleteUser", deleteUser)
}
