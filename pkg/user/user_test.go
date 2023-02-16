package user

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
	mgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
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
	uuidSlice     = []string{uuid.NewString()}
	uuidSliceS, _ = json.Marshal(uuidSlice)
	signType      = basetypes.SignMethod_Email
	appID         = uuid.NewString()
	userInfo      = npool.User{
		ID:                          uuid.NewString(),
		AppID:                       appID,
		EmailAddress:                uuid.NewString(),
		PhoneNO:                     uuid.NewString(),
		ImportedFromAppID:           uuid.NewString(),
		Username:                    uuid.NewString(),
		AddressFieldsString:         string(uuidSliceS),
		Gender:                      uuid.NewString(),
		PostalCode:                  uuid.NewString(),
		Age:                         0,
		Birthday:                    0,
		Avatar:                      uuid.NewString(),
		Organization:                uuid.NewString(),
		FirstName:                   uuid.NewString(),
		LastName:                    uuid.NewString(),
		IDNumber:                    uuid.NewString(),
		SigninVerifyByGoogleAuthInt: 0,
		SigninVerifyTypeStr:         signType.String(),
		GoogleAuthVerifiedInt:       0,
		GoogleSecret:                appID,
		HasGoogleSecret:             true,
		Roles:                       []string{""},
	}
)

func creatUser(t *testing.T) {
	var (
		id                = userInfo.ID
		appID             = userInfo.AppID
		importedFromAppID = userInfo.ImportedFromAppID
		strVal            = "AAA"
		userReq           = npool.UserReq{
			ID:                 &id,
			AppID:              &appID,
			EmailAddress:       &userInfo.EmailAddress,
			PhoneNO:            &userInfo.PhoneNO,
			ImportedFromAppID:  &importedFromAppID,
			Username:           &userInfo.Username,
			AddressFields:      uuidSlice,
			Gender:             &userInfo.Gender,
			PostalCode:         &userInfo.PostalCode,
			Age:                &userInfo.Age,
			Birthday:           &userInfo.Birthday,
			Avatar:             &userInfo.Avatar,
			Organization:       &userInfo.Organization,
			FirstName:          &userInfo.FirstName,
			LastName:           &userInfo.LastName,
			IDNumber:           &userInfo.IDNumber,
			GoogleAuthVerified: &userInfo.GoogleAuthVerified,
			SigninVerifyType:   &signType,
			PasswordHash:       &strVal,
			GoogleSecret:       &appID,
			ThirdPartyID:       &strVal,
			ThirdPartyUserID:   &strVal,
			ThirdPartyUsername: &strVal,
			ThirdPartyAvatar:   &strVal,
			Banned:             &userInfo.Banned,
			BanMessage:         &userInfo.BanMessage,
		}
	)
	info, err := CreateUser(context.Background(), &userReq)
	if assert.Nil(t, err) {
		userInfo.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &userInfo)
	}
}

func updateUser(t *testing.T) {
	var (
		appID        = userInfo.AppID
		strVal       = "AAA"
		kol          = true
		kolConfirmed = true
		userReq      = npool.UserReq{
			ID:                 &userInfo.ID,
			AppID:              &userInfo.AppID,
			EmailAddress:       &userInfo.EmailAddress,
			PhoneNO:            &userInfo.PhoneNO,
			ImportedFromAppID:  &userInfo.ImportedFromAppID,
			Username:           &userInfo.Username,
			AddressFields:      uuidSlice,
			Gender:             &userInfo.Gender,
			PostalCode:         &userInfo.PostalCode,
			Age:                &userInfo.Age,
			Birthday:           &userInfo.Birthday,
			Avatar:             &userInfo.Avatar,
			Organization:       &userInfo.Organization,
			FirstName:          &userInfo.FirstName,
			LastName:           &userInfo.LastName,
			IDNumber:           &userInfo.IDNumber,
			GoogleAuthVerified: &userInfo.GoogleAuthVerified,
			SigninVerifyType:   &signType,
			PasswordHash:       &strVal,
			GoogleSecret:       &appID,
			ThirdPartyID:       &strVal,
			ThirdPartyUserID:   &strVal,
			ThirdPartyUsername: &strVal,
			ThirdPartyAvatar:   &strVal,
			Banned:             &userInfo.Banned,
			BanMessage:         &userInfo.BanMessage,
			Kol:                &kol,
			KolConfirmed:       &kolConfirmed,
		}
	)

	userInfo.Kol = true
	userInfo.KolConfirmed = true

	info, err := UpdateUser(context.Background(), &userReq)
	if assert.Nil(t, err) {
		info.Roles = userInfo.Roles
		assert.Equal(t, info, &userInfo)
	}
}

func getUser(t *testing.T) {
	info, err := GetUser(context.Background(), userInfo.AppID, userInfo.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &userInfo)
	}
}

func getUsers(t *testing.T) {
	infos, _, err := GetUsers(context.Background(), &mgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: userInfo.AppID,
		},
	}, 0, 1)
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func getManyUsers(t *testing.T) {
	infos, _, err := GetManyUsers(context.Background(), []string{userInfo.ID})
	if !assert.Nil(t, err) {
		assert.Equal(t, infos[0], &userInfo)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("creatUser", creatUser)
	t.Run("updateUser", updateUser)
	t.Run("getUser", getUser)
	t.Run("getUsers", getUsers)
	t.Run("getManyUsers", getManyUsers)
}
