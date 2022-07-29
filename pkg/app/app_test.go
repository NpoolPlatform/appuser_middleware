package app

import (
	"context"
	"encoding/json"
	"fmt"
	testinit "github.com/NpoolPlatform/appuser-manager/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

//}
//var entApp = ent.App{
//	ID:          uuid.New(),
//	CreatedBy:   uuid.New(),
//	Name:        uuid.New().String(),
//	Description: uuid.New().String(),
//	Logo:        uuid.New().String(),
//}
//
//var (
//	id        = entApp.ID.String()
//	createdBy = entApp.CreatedBy.String()
//	appInfo   = npool.AppReq{
//		ID:          &id,
//		CreatedBy:   &createdBy,
//		Name:        &entApp.Name,
//		Description: &entApp.Description,
//		Logo:        &entApp.Logo,
//	}
//)
//
//var info *ent.App
//
//func rowToObject(row *ent.App) *ent.App {
//	return &ent.App{
//		ID:          row.ID,
//		CreatedBy:   row.CreatedBy,
//		Name:        row.Name,
//		Logo:        row.Logo,
//		Description: row.Description,
//		CreatedAt:   row.CreatedAt,
//	}
//}

var (
	appInfo = npool.App{
		ID:                 uuid.NewString(),
		CreatedBy:          uuid.NewString(),
		Name:               uuid.NewString(),
		Logo:               uuid.NewString(),
		Description:        uuid.NewString(),
		Banned:             false,
		BanMessage:         uuid.NewString(),
		SignupMethods:      []string{uuid.NewString(), uuid.NewString()},
		ExtSigninMethods:   []string{uuid.NewString(), uuid.NewString()},
		RecaptchaMethod:    uuid.NewString(),
		KycEnable:          false,
		SigninVerifyEnable: false,
		InvitationCodeMust: false,
	}
)

var (
	appReq = npool.AppReq{
		ID:                 &appInfo.ID,
		CreatedBy:          &appInfo.CreatedBy,
		Name:               &appInfo.Name,
		Logo:               &appInfo.Logo,
		Description:        &appInfo.Description,
		Banned:             &appInfo.Banned,
		BanMessage:         &appInfo.BanMessage,
		SignupMethods:      appInfo.SignupMethods,
		ExtSigninMethods:   appInfo.ExtSigninMethods,
		RecaptchaMethod:    &appInfo.RecaptchaMethod,
		KycEnable:          &appInfo.KycEnable,
		SigninVerifyEnable: &appInfo.SigninVerifyEnable,
		InvitationCodeMust: &appInfo.InvitationCodeMust,
	}
)

func QueryEnt2Grpc(row *AppQueryResp) (*npool.App, error) {
	if row == nil {
		return nil, nil
	}

	methods := []string{}
	if row.SignupMethods != "" {
		err := json.Unmarshal([]byte(row.SignupMethods), &methods)
		if err != nil {
			return nil, err
		}
	}

	emethods := []string{}
	if row.ExtSigninMethods != "" {
		err := json.Unmarshal([]byte(row.ExtSigninMethods), &emethods)
		if err != nil {
			return nil, err
		}
	}

	return &npool.App{
		ID:                 row.ID.String(),
		CreatedBy:          row.CreatedBy.String(),
		Name:               row.Name,
		Logo:               row.Logo,
		Description:        row.Description,
		Banned:             row.Banned,
		BanMessage:         row.BanMessage,
		SignupMethods:      methods,
		ExtSigninMethods:   emethods,
		RecaptchaMethod:    row.RecaptchaMethod,
		KycEnable:          row.KycEnable != 0,
		SigninVerifyEnable: row.SigninVerifyEnable != 0,
		InvitationCodeMust: row.InvitationCodeMust != 0,
		CreatedAt:          row.CreatedAt,
	}, nil
}

func CreateEnt2Grpc(row *AppCreateResp) (*npool.App, error) {
	banned := false
	bannedMsg := ""
	if row.BanApp != nil {
		banned = true
		bannedMsg = row.BanApp.Message
	}

	return &npool.App{
		ID:          row.App.ID.String(),
		CreatedBy:   row.App.CreatedBy.String(),
		Name:        row.App.Name,
		Logo:        row.App.Logo,
		Description: row.App.Description,
		Banned:      banned,
		BanMessage:  bannedMsg,

		SignupMethods:    row.AppControl.SignupMethods,
		ExtSigninMethods: row.AppControl.ExternSigninMethods,

		RecaptchaMethod:    row.AppControl.RecaptchaMethod,
		KycEnable:          row.AppControl.KycEnable,
		SigninVerifyEnable: row.AppControl.SigninVerifyEnable,
		InvitationCodeMust: row.AppControl.InvitationCodeMust,

		CreatedAt: row.App.CreatedAt,
	}, nil
}
func TestApp(t *testing.T) {
	cinfo, err := CreateApp(context.Background(), &appReq)
	if !assert.Nil(t, err) {
		return
	}
	qinfo, err := GetApp(context.Background(), appInfo.ID)
	if !assert.Nil(t, err) {
		return
	}

	c, err := CreateEnt2Grpc(cinfo)
	if !assert.Nil(t, err) {
		return
	}
	q, err := QueryEnt2Grpc(qinfo)
	if !assert.Nil(t, err) {
		return
	}

	assert.Equal(t, c, q)
}

//
//func getUser(t *testing.T) {
//	var err error
//	info, err = GetUser(context.Background(), &appRoleInfo)
//	if assert.Nil(t, err) {
//		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
//			entAppRole.ID = info.ID
//		}
//		assert.Equal(t, rowToObject(info), &entAppRole)
//	}
//}
