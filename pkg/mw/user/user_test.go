package user

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

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
	ret           = npool.User{
		ID:                          uuid.NewString(),
		AppID:                       appID,
		EmailAddress:                "aaa@hhh.ccc",
		PhoneNO:                     "+8613612203166",
		ImportedFromAppID:           uuid.NewString(),
		Username:                    "adfjskajfdl.afd-",
		AddressFieldsString:         string(uuidSliceS),
		AddressFields:               uuidSlice,
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
		SigninVerifyType:            signType,
		GoogleAuthVerifiedInt:       0,
		GoogleSecret:                appID,
		HasGoogleSecret:             true,
		Roles:                       []string{""},
		ActionCredits:               "0",
	}
)

func creatUser(t *testing.T) {
	ret.PhoneNO = fmt.Sprintf("+86%v", rand.Intn(100000000)+10000)           //nolint
	ret.EmailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+10000) //nolint
	var (
		strVal = "AAA"
		req    = npool.UserReq{
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
			SigninVerifyType:   &signType,
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
			ActionCredits:       ret.ActionCredits,
			AddressFieldsString: "[]",
			AddressFields:       []string{},
			SigninVerifyTypeStr: basetypes.SignMethod_Email.String(),
			SigninVerifyType:    basetypes.SignMethod_Email,
		}
	)

	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID),
		WithAppID(req.GetAppID()),
		WithPhoneNO(req.PhoneNO),
		WithEmailAddress(req.EmailAddress),
		WithImportedFromAppID(req.ImportedFromAppID),
		WithPasswordHash(req.PasswordHash),
	)
	assert.Nil(t, err)

	info, err := handler.CreateUser(context.Background())
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
		appID        = ret.AppID
		strVal       = "AAA"
		kol          = true
		kolConfirmed = true
		credits      = "1.2342"
		req          = npool.UserReq{
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
			SigninVerifyType:   &signType,
			PasswordHash:       &strVal,
			GoogleSecret:       &appID,
			ThirdPartyID:       &strVal,
			ThirdPartyUserID:   &strVal,
			ThirdPartyUsername: &strVal,
			ThirdPartyAvatar:   &strVal,
			Banned:             &ret.Banned,
			BanMessage:         &ret.BanMessage,
			Kol:                &kol,
			KolConfirmed:       &kolConfirmed,
			ActionCredits:      &credits,
		}
	)

	ret.Kol = true
	ret.KolConfirmed = true
	ret.ActionCredits = credits

	handler, err := NewHandler(
		context.Background(),
		WithID(req.ID),
		WithAppID(req.GetAppID()),
		WithPhoneNO(req.PhoneNO),
		WithEmailAddress(req.EmailAddress),
		WithImportedFromAppID(req.ImportedFromAppID),
		WithPasswordHash(req.PasswordHash),
		WithFirstName(req.FirstName),
		WithLastName(req.LastName),
		WithBirthday(req.Birthday),
		WithGender(req.Gender),
		WithAvatar(req.Avatar),
		WithUsername(req.Username),
		WithPostalCode(req.PostalCode),
		WithAge(req.Age),
		WithOrganization(req.Organization),
		WithIDNumber(req.IDNumber),
		WithAddressFields(req.AddressFields),
		WithGoogleSecret(req.GoogleSecret),
		WithGoogleAuthVerified(req.GoogleAuthVerified),
		WithKol(req.Kol),
		WithKolConfirmed(req.KolConfirmed),
		WithActionCredits(req.ActionCredits),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateUser(context.Background())
	if assert.Nil(t, err) {
		ret.Roles = info.Roles
		assert.Equal(t, info, &ret)
	}
}

func getUser(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithAppID(ret.AppID),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetUser(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getUsers(t *testing.T) {
	conds := &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds, 0, 1),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetUsers(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func getManyUsers(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithIDs([]string{ret.ID}),
	)
	assert.Nil(t, err)

	infos, err := handler.GetManyUsers(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, infos[0], &ret)
	}
}

func TestUser(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("creatUser", creatUser)
	t.Run("updateUser", updateUser)
	t.Run("getUser", getUser)
	t.Run("getUsers", getUsers)
	t.Run("getManyUsers", getManyUsers)
}
