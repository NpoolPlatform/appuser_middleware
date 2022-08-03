package app

import (
	"context"
	"encoding/json"
	"fmt"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/appuser-middleware/pkg/testinit"
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
	appInfo       = App{
		ID:                 uuid.New(),
		CreatedBy:          uuid.New(),
		Name:               uuid.NewString(),
		Logo:               uuid.NewString(),
		Description:        uuid.NewString(),
		Banned:             false,
		SignupMethods:      string(uuidSliceS),
		ExtSigninMethods:   string(uuidSliceS),
		RecaptchaMethod:    uuid.NewString(),
		KycEnable:          1,
		SigninVerifyEnable: 1,
		InvitationCodeMust: 1,
	}
)

func creatApp(t *testing.T) {
	var (
		id        = appInfo.ID.String()
		createdBy = appInfo.CreatedBy.String()
		boolVal   = true
		appReq    = npool.AppReq{
			ID:                 &id,
			CreatedBy:          &createdBy,
			Name:               &appInfo.Name,
			Logo:               &appInfo.Logo,
			Description:        &appInfo.Description,
			Banned:             &appInfo.Banned,
			BanMessage:         &appInfo.BanMessage,
			SignupMethods:      uuidSlice,
			ExtSigninMethods:   uuidSlice,
			RecaptchaMethod:    &appInfo.RecaptchaMethod,
			KycEnable:          &boolVal,
			SigninVerifyEnable: &boolVal,
			InvitationCodeMust: &boolVal,
		}
	)
	info, err := CreateApp(context.Background(), &appReq)
	if assert.Nil(t, err) {
		info.CreatedAt = appInfo.CreatedAt
		assert.Equal(t, info, &appInfo)
	}
}

func updateApp(t *testing.T) {
	var (
		id        = appInfo.ID.String()
		createdBy = appInfo.CreatedBy.String()
		boolVal   = true
		appReq    = npool.AppReq{
			ID:                 &id,
			CreatedBy:          &createdBy,
			Name:               &appInfo.Name,
			Logo:               &appInfo.Logo,
			Description:        &appInfo.Description,
			Banned:             &appInfo.Banned,
			BanMessage:         &appInfo.BanMessage,
			SignupMethods:      uuidSlice,
			ExtSigninMethods:   uuidSlice,
			RecaptchaMethod:    &appInfo.RecaptchaMethod,
			KycEnable:          &boolVal,
			SigninVerifyEnable: &boolVal,
			InvitationCodeMust: &boolVal,
		}
	)
	info, err := UpdateApp(context.Background(), &appReq)
	if assert.Nil(t, err) {
		info.CreatedAt = appInfo.CreatedAt
		assert.Equal(t, info, &appInfo)
	}
}

func getApp(t *testing.T) {
	info, err := GetApp(context.Background(), appInfo.ID.String())
	if assert.Nil(t, err) {
		info.CreatedAt = appInfo.CreatedAt
		assert.Equal(t, info, &appInfo)
	}
}

func getApps(t *testing.T) {
	infos, err := GetApps(context.Background(), 0, 1)
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func getUserApps(t *testing.T) {
	infos, err := GetUserApps(context.Background(), appInfo.CreatedBy.String(), 0, 1)
	if !assert.Nil(t, err) {
		infos[0].CreatedAt = appInfo.CreatedAt
		assert.Equal(t, infos[0], &appInfo)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("createApp", creatApp)
	t.Run("updateApp", updateApp)
	t.Run("getApp", getApp)
	t.Run("getApps", getApps)
	t.Run("getUserApps", getUserApps)
}
