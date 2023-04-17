package app

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

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
	signupMethods       = []basetypes.SignMethod{basetypes.SignMethod_Email, basetypes.SignMethod_Mobile}
	signupMethodsStr    = fmt.Sprintf(`["%v", "%v"]`, basetypes.SignMethod_Email.String(), basetypes.SignMethod_Mobile.String())
	extSignupMethods    = []basetypes.SignMethod{}
	extSignupMethodsStr = fmt.Sprintf(`[]`)
	rec                 = basetypes.RecaptchaMethod_GoogleRecaptchaV3
	commitButton        = uuid.NewString()
	ret                 = npool.App{
		ID:                          uuid.NewString(),
		CreatedBy:                   uuid.NewString(),
		Name:                        uuid.NewString(),
		Logo:                        uuid.NewString(),
		Description:                 uuid.NewString(),
		Banned:                      false,
		SignupMethodsStr:            signupMethodsStr,
		SignupMethods:               signupMethods,
		ExtSigninMethodsStr:         extSignupMethodsStr,
		ExtSigninMethods:            extSignupMethods,
		RecaptchaMethodStr:          basetypes.RecaptchaMethod_GoogleRecaptchaV3.String(),
		RecaptchaMethod:             basetypes.RecaptchaMethod_GoogleRecaptchaV3,
		KycEnable:                   true,
		SigninVerifyEnable:          true,
		InvitationCodeMust:          true,
		CreateInvitationCodeWhenStr: basetypes.CreateInvitationCodeWhen_Registration.String(),
		CreateInvitationCodeWhen:    basetypes.CreateInvitationCodeWhen_Registration,
		MaxTypedCouponsPerOrder:     1,
		Maintaining:                 true,
		CommitButtonTargetsStr:      fmt.Sprintf("[\"%v\"]", commitButton),
		CommitButtonTargets:         []string{commitButton},
	}
)

func creatApp(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithCreatedBy(ret.GetCreatedBy()),
		WithName(&ret.Name),
		WithLogo(&ret.Logo),
		WithDescription(&ret.Description),
		WithSignupMethods(ret.GetSignupMethods()),
		WithExtSigninMethods(ret.GetExtSigninMethods()),
		WithRecaptchaMethod(&ret.RecaptchaMethod),
		WithKycEnable(&ret.KycEnable),
		WithSigninVerifyEnable(&ret.SigninVerifyEnable),
		WithInvitationCodeMust(&ret.InvitationCodeMust),
		WithCreateInvitationCodeWhen(&ret.CreateInvitationCodeWhen),
		WithMaxTypedCouponsPerOrder(&ret.MaxTypedCouponsPerOrder),
		WithMaintaining(&ret.Maintaining),
		WithCommitButtonTargets(ret.GetCommitButtonTargets()),
	)
	assert.Nil(t, err)
	info, err := handler.CreateApp(context.Background())
	if assert.Nil(t, err) {
		info.CreatedAt = ret.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

/*
func updateApp(t *testing.T) {
	var (
		boolVal                 = true
		createIvCodeWhen        = ctrl.CreateInvitationCodeWhen_SetToKol
		maxTypedCouponsPerOrder = uint32(5)

		req = npool.AppReq{
			ID:                       &ret.ID,
			CreatedBy:                &ret.Name,
			Name:                     &ret.Name,
			Logo:                     &ret.Logo,
			Description:              &ret.Description,
			Banned:                   &ret.Banned,
			BanMessage:               &ret.BanMessage,
			SignupMethods:            uuidSlice,
			ExtSigninMethods:         uuidSlice,
			RecaptchaMethod:          &rec,
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
	infos, err := GetApps(context.Background(), 0, 1)
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func getUserApps(t *testing.T) {
	infos, _, err := GetUserApps(context.Background(), ret.CreatedBy, 0, 1)
	if !assert.Nil(t, err) {
		infos[0].CreatedAt = ret.CreatedAt
		assert.Equal(t, infos[0], &ret)
	}
}
*/

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("createApp", creatApp)
	// t.Run("updateApp", updateApp)
	// t.Run("getApp", getApp)
	// t.Run("getApps", getApps)
	// t.Run("getUserApps", getUserApps)
}
