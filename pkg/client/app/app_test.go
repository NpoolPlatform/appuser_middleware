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

	ctrl "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
	rcpt "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/recaptcha"
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
	commitButton = uuid.NewString()
	uuidSlice    = []basetypes.SignMethod{basetypes.SignMethod_Email, basetypes.SignMethod_Mobile}
	uuidSliceS   = fmt.Sprintf(`["%v", "%v"]`, basetypes.SignMethod_Email.String(), basetypes.SignMethod_Mobile.String())
	appInfo      = npool.App{
		ID:                          uuid.NewString(),
		CreatedBy:                   uuid.NewString(),
		Name:                        uuid.NewString(),
		Logo:                        uuid.NewString(),
		Description:                 uuid.NewString(),
		Banned:                      false,
		SignupMethodsStr:            uuidSliceS,
		SignupMethods:               uuidSlice,
		ExtSigninMethodsStr:         uuidSliceS,
		ExtSigninMethods:            uuidSlice,
		RecaptchaMethodStr:          rcpt.RecaptchaType_GoogleRecaptchaV3.String(),
		RecaptchaMethod:             rcpt.RecaptchaType_GoogleRecaptchaV3,
		KycEnableInt:                1,
		KycEnable:                   true,
		SigninVerifyEnableInt:       1,
		SigninVerifyEnable:          true,
		InvitationCodeMustInt:       1,
		InvitationCodeMust:          true,
		CreatedAt:                   0,
		CreateInvitationCodeWhenStr: ctrl.CreateInvitationCodeWhen_Registration.String(),
		CreateInvitationCodeWhen:    ctrl.CreateInvitationCodeWhen_Registration,
		MaxTypedCouponsPerOrder:     1,
		Maintaining:                 true,
		CommitButtonTargetsStr:      fmt.Sprintf("[\"%v\"]", commitButton),
		CommitButtonTargets:         []string{commitButton},
	}
)

func creatApp(t *testing.T) {
	var (
		id        = appInfo.ID
		createdBy = appInfo.CreatedBy
		boolVal   = true
		appReq    = npool.AppReq{
			ID:                       &id,
			CreatedBy:                &createdBy,
			Name:                     &appInfo.Name,
			Logo:                     &appInfo.Logo,
			Description:              &appInfo.Description,
			Banned:                   &appInfo.Banned,
			BanMessage:               &appInfo.BanMessage,
			SignupMethods:            uuidSlice,
			ExtSigninMethods:         uuidSlice,
			RecaptchaMethod:          &appInfo.RecaptchaMethod,
			KycEnable:                &boolVal,
			SigninVerifyEnable:       &boolVal,
			InvitationCodeMust:       &boolVal,
			CreateInvitationCodeWhen: &appInfo.CreateInvitationCodeWhen,
			Maintaining:              &appInfo.Maintaining,
			CommitButtonTargets:      appInfo.CommitButtonTargets,
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
		boolVal                 = true
		createIvCodeWhen        = ctrl.CreateInvitationCodeWhen_SetToKol
		maxTypedCouponsPerOrder = uint32(5)
		appReq                  = npool.AppReq{
			ID:                       &appInfo.ID,
			CreatedBy:                &appInfo.Name,
			Name:                     &appInfo.Name,
			Logo:                     &appInfo.Logo,
			Description:              &appInfo.Description,
			Banned:                   &appInfo.Banned,
			BanMessage:               &appInfo.BanMessage,
			SignupMethods:            uuidSlice,
			ExtSigninMethods:         uuidSlice,
			RecaptchaMethod:          &appInfo.RecaptchaMethod,
			KycEnable:                &boolVal,
			SigninVerifyEnable:       &boolVal,
			InvitationCodeMust:       &boolVal,
			CreateInvitationCodeWhen: &createIvCodeWhen,
			MaxTypedCouponsPerOrder:  &maxTypedCouponsPerOrder,
		}
	)

	appInfo.MaxTypedCouponsPerOrder = maxTypedCouponsPerOrder
	appInfo.CreateInvitationCodeWhenStr = createIvCodeWhen.String()
	appInfo.CreateInvitationCodeWhen = createIvCodeWhen

	info, err := UpdateApp(context.Background(), &appReq)
	if assert.Nil(t, err) {
		info.CreatedAt = appInfo.CreatedAt
		assert.Equal(t, info, &appInfo)
	}
}

func getApp(t *testing.T) {
	info, err := GetApp(context.Background(), appInfo.ID)
	if assert.Nil(t, err) {
		info.CreatedAt = appInfo.CreatedAt
		assert.Equal(t, info, &appInfo)
	}
}

func getApps(t *testing.T) {
	infos, _, err := GetApps(context.Background(), 0, 1)
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func getUserApps(t *testing.T) {
	infos, _, err := GetUserApps(context.Background(), appInfo.CreatedBy, 0, 1)
	if !assert.Nil(t, err) {
		infos[0].CreatedAt = appInfo.CreatedAt
		assert.Equal(t, infos[0], &appInfo)
	}
}

func getManyApps(t *testing.T) {
	infos, _, err := GetManyApps(context.Background(), []string{appInfo.ID})
	if !assert.Nil(t, err) {
		infos[0].CreatedAt = appInfo.CreatedAt
		assert.Equal(t, infos[0], &appInfo)
	}
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
}
