package recoverycode

import (
	"context"
	"encoding/json"
	"fmt"
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

	appusermwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/recoverycode"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/appuser-middleware/pkg/testinit"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	appusermwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/user"
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
	addressFields     = []string{uuid.NewString()}
	addressFieldsS, _ = json.Marshal(addressFields)
	appID             = uuid.NewString()
	ret               = appusermwpb.User{
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
		ActionCredits:       "0",
		Banned:              true,
		BanMessage:          uuid.NewString(),
	}
)

func setupRecoveryCode(t *testing.T) func(*testing.T) {
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

	id := ret.EntID
	appID = ret.AppID
	strVal := "AAA"
	user, err := appusermwcli.CreateUser(context.Background(), &appusermwpb.UserReq{
		EntID:              &id,
		AppID:              &appID,
		EmailAddress:       &ret.EmailAddress,
		PhoneNO:            &ret.PhoneNO,
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
		ThirdPartyUserID:   &strVal,
		ThirdPartyUsername: &strVal,
		ThirdPartyAvatar:   &strVal,
		Banned:             &ret.Banned,
		BanMessage:         &ret.BanMessage,
	})
	assert.Nil(t, err)
	assert.NotNil(t, user)
	return func(*testing.T) {
		_, _ = appmwcli.DeleteApp(context.Background(), app1.ID)
		_, _ = appusermwcli.DeleteUser(context.Background(), app1.EntID, user.ID)
	}
}

func generateRecoveryCodes(t *testing.T) {
	infos, err := GenerateRecoveryCodes(context.Background(), &npool.RecoveryCodeReq{
		AppID:  &ret.AppID,
		UserID: &ret.EntID,
	})
	assert.Nil(t, err)
	assert.Equal(t, 16, len(infos))
}

func getRecoveryCodes(t *testing.T) {
	conds := &npool.Conds{
		AppID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
	}
	infos, _, err := GetRecoveryCodes(context.Background(), conds, 0, 16)
	assert.Nil(t, err)
	assert.Equal(t, 16, len(infos))
}

func TestRecoveryCode(t *testing.T) {
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

	teardown := setupRecoveryCode(t)
	defer teardown(t)

	t.Run("generateRecoveryCodes", generateRecoveryCodes)
	t.Run("getRecoveryCodes", getRecoveryCodes)
}
