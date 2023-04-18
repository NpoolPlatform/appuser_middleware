package role

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	// "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/appuser-middleware/pkg/testinit"
	// basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	app "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
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
	ret = npool.Role{
		ID:          uuid.NewString(),
		AppID:       uuid.NewString(),
		AppName:     uuid.NewString(),
		CreatedBy:   uuid.NewString(),
		Role:        uuid.NewString(),
		Description: uuid.NewString(),
		Default:     true,
		Genesis:     false,
	}
)

func setup(t *testing.T) func(*testing.T) {
	ah, err := app.NewHandler(
		context.Background(),
		app.WithID(&ret.AppID),
		app.WithCreatedBy(ret.ID),
		app.WithName(&ret.AppName),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ah)
	app1, err := ah.CreateApp(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, app1)

	return func(*testing.T) {
		ah.DeleteApp(context.Background())
	}
}

func creatRole(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(ret.AppID),
		WithCreatedBy(&ret.CreatedBy),
		WithRole(&ret.Role),
		WithDescription(&ret.Description),
		WithDefault(&ret.Default),
		WithGenesis(&ret.Genesis),
	)
	assert.Nil(t, err)

	info, err := handler.CreateRole(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

/*
func updateRole(t *testing.T) {
	ret.PhoneNO = fmt.Sprintf("+86%v", rand.Intn(100000000)+10000)           //nolint
	ret.EmailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+10000) //nolint
	var (
		appID        = ret.AppID
		strVal       = "AAA"
		kol          = true
		kolConfirmed = true
		credits      = "1.2342"
		req          = npool.RoleReq{
			ID:                 &ret.ID,
			AppID:              &ret.AppID,
			EmailAddress:       &ret.EmailAddress,
			PhoneNO:            &ret.PhoneNO,
			ImportedFromAppID:  &ret.ImportedFromAppID,
			Rolename:           &ret.Rolename,
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
			ThirdPartyRoleID:   &strVal,
			ThirdPartyRolename: &strVal,
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
		WithRolename(req.Rolename),
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

	info, err := handler.UpdateRole(context.Background())
	if assert.Nil(t, err) {
		ret.Roles = info.Roles
		assert.Equal(t, info, &ret)
	}
}

func getRole(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithAppID(ret.AppID),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetRole(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getRoles(t *testing.T) {
	conds := &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds, 0, 1),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetRoles(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func getManyRoles(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithIDs([]string{ret.ID}),
	)
	assert.Nil(t, err)

	infos, err := handler.GetManyRoles(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, infos[0], &ret)
	}
}
*/

func TestRole(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("creatRole", creatRole)
	// t.Run("updateRole", updateRole)
	// t.Run("getRole", getRole)
	// t.Run("getRoles", getRoles)
	// t.Run("getManyRoles", getManyRoles)
}
