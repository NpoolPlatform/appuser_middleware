package app

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
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
	commitButton        = uuid.NewString()
	signupMethods       = []basetypes.SignMethod{basetypes.SignMethod_Email, basetypes.SignMethod_Mobile}
	signupMethodsStr    = fmt.Sprintf(`["%v", "%v"]`, basetypes.SignMethod_Email.String(), basetypes.SignMethod_Mobile.String())
	extSigninMethods    = []basetypes.SignMethod{}
	extSigninMethodsStr = fmt.Sprintf(`[]`)
	ret                 = npool.App{
		ID:                          uuid.NewString(),
		CreatedBy:                   uuid.NewString(),
		Name:                        uuid.NewString(),
		Logo:                        uuid.NewString(),
		Description:                 uuid.NewString(),
		Banned:                      false,
		SignupMethodsStr:            signupMethodsStr,
		SignupMethods:               signupMethods,
		ExtSigninMethodsStr:         extSigninMethodsStr,
		ExtSigninMethods:            extSigninMethods,
		RecaptchaMethodStr:          basetypes.RecaptchaMethod_GoogleRecaptchaV3.String(),
		RecaptchaMethod:             basetypes.RecaptchaMethod_GoogleRecaptchaV3,
		KycEnable:                   true,
		SigninVerifyEnable:          true,
		InvitationCodeMust:          true,
		CreatedAt:                   0,
		CreateInvitationCodeWhenStr: basetypes.CreateInvitationCodeWhen_Registration.String(),
		CreateInvitationCodeWhen:    basetypes.CreateInvitationCodeWhen_Registration,
		MaxTypedCouponsPerOrder:     1,
		Maintaining:                 true,
		CommitButtonTargetsStr:      fmt.Sprintf("[\"%v\"]", commitButton),
		CommitButtonTargets:         []string{commitButton},
	}
)

func creatApp(t *testing.T) {
	var (
		id        = ret.ID
		createdBy = ret.CreatedBy
		boolVal   = true
		req       = npool.AppReq{
			ID:                       &id,
			CreatedBy:                &createdBy,
			Name:                     &ret.Name,
			Logo:                     &ret.Logo,
			Description:              &ret.Description,
			Banned:                   &ret.Banned,
			BanMessage:               &ret.BanMessage,
			SignupMethods:            signupMethods,
			ExtSigninMethods:         extSigninMethods,
			RecaptchaMethod:          &ret.RecaptchaMethod,
			KycEnable:                &boolVal,
			SigninVerifyEnable:       &boolVal,
			InvitationCodeMust:       &boolVal,
			CreateInvitationCodeWhen: &ret.CreateInvitationCodeWhen,
			Maintaining:              &ret.Maintaining,
			CommitButtonTargets:      ret.CommitButtonTargets,
		}
	)
	info, err := CreateApp(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.ExtSigninMethods = info.ExtSigninMethods
		assert.Equal(t, info, &ret)
	}
}

func updateApp(t *testing.T) {
	var (
		boolVal                 = true
		createIvCodeWhen        = basetypes.CreateInvitationCodeWhen_SetToKol
		maxTypedCouponsPerOrder = uint32(5)
		req                     = npool.AppReq{
			ID:                       &ret.ID,
			CreatedBy:                &ret.Name,
			Name:                     &ret.Name,
			Logo:                     &ret.Logo,
			Description:              &ret.Description,
			Banned:                   &ret.Banned,
			BanMessage:               &ret.BanMessage,
			SignupMethods:            signupMethods,
			ExtSigninMethods:         extSigninMethods,
			RecaptchaMethod:          &ret.RecaptchaMethod,
			KycEnable:                &boolVal,
			SigninVerifyEnable:       &boolVal,
			InvitationCodeMust:       &boolVal,
			CreateInvitationCodeWhen: &createIvCodeWhen,
			MaxTypedCouponsPerOrder:  &maxTypedCouponsPerOrder,
		}
	)

	ret.MaxTypedCouponsPerOrder = maxTypedCouponsPerOrder
	ret.CreateInvitationCodeWhenStr = createIvCodeWhen.String()
	ret.CreateInvitationCodeWhen = createIvCodeWhen

	info, err := UpdateApp(context.Background(), &req)
	if assert.Nil(t, err) {
		info.CreatedAt = ret.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

func getApp(t *testing.T) {
	info, err := GetApp(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		info.CreatedAt = ret.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

func getApps(t *testing.T) {
	infos, _, err := GetApps(context.Background(), 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func getUserApps(t *testing.T) {
	infos, _, err := GetUserApps(context.Background(), ret.CreatedBy, 0, 1)
	if assert.Nil(t, err) {
		infos[0].CreatedAt = ret.CreatedAt
		assert.Equal(t, infos[0], &ret)
	}
}

func getManyApps(t *testing.T) {
	infos, _, err := GetManyApps(context.Background(), []string{ret.ID})
	if assert.Nil(t, err) {
		infos[0].CreatedAt = ret.CreatedAt
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteApp(t *testing.T) {
	info, err := DeleteApp(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetApp(context.Background(), ret.ID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	t.Run("createApp", creatApp)
	t.Run("updateApp", updateApp)
	t.Run("getApp", getApp)
	t.Run("getApps", getApps)
	t.Run("getUserApps", getUserApps)
	t.Run("getManyApps", getManyApps)
	t.Run("deleteApp", deleteApp)
}
