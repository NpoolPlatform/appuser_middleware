package auth

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	handler "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/handler"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/auth"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/appuser-middleware/pkg/testinit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	app "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	role "github.com/NpoolPlatform/appuser-middleware/pkg/mw/role"
	user "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user"
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
	ret = npool.Auth{
		ID:       uuid.NewString(),
		AppID:    uuid.NewString(),
		RoleID:   uuid.NewString(),
		UserID:   uuid.NewString(),
		Resource: uuid.NewString(),
		Method:   "POST",
	}
)

func setupAuth(t *testing.T) func(*testing.T) {
	ret.AppName = ret.AppID
	ret.RoleName = ret.RoleID

	ah, err := app.NewHandler(
		context.Background(),
		app.WithID(&ret.AppID),
		app.WithCreatedBy(ret.UserID),
		app.WithName(&ret.AppID),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ah)
	app1, err := ah.CreateApp(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, app1)

	rh, err := role.NewHandler(
		context.Background(),
		role.WithID(&ret.RoleID),
		role.WithAppID(ret.GetAppID()),
		role.WithCreatedBy(&ret.UserID),
		role.WithRole(&ret.RoleID),
	)
	assert.Nil(t, err)
	assert.NotNil(t, rh)
	role1, err := rh.CreateRole(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, role1)

	ret.PhoneNO = fmt.Sprintf("+86%v", rand.Intn(100000000)+1000000)           //nolint
	ret.EmailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+1000000) //nolint
	passwordHash := uuid.NewString()

	ret.AppName = ret.AppID

	uh, err := user.NewHandler(
		context.Background(),
		user.WithID(&ret.UserID),
		user.WithAppID(ret.GetAppID()),
		user.WithPhoneNO(&ret.PhoneNO),
		user.WithEmailAddress(&ret.EmailAddress),
		user.WithPasswordHash(&passwordHash),
	)
	assert.Nil(t, err)
	assert.NotNil(t, uh)
	user1, err := uh.CreateUser(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, user1)

	return func(*testing.T) {
		_, _ = ah.DeleteApp(context.Background())
		_, _ = rh.DeleteRole(context.Background())
		_, _ = uh.DeleteUser(context.Background())
	}
}

func creatAuth(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		handler.WithID(&ret.ID),
		handler.WithAppID(ret.AppID),
		handler.WithRoleID(&ret.RoleID),
		handler.WithUserID(&ret.UserID),
		handler.WithResource(&ret.Resource),
		handler.WithMethod(&ret.Method),
	)
	assert.Nil(t, err)

	info, err := handler.CreateAuth(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateAuth(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		handler.WithID(&ret.ID),
		handler.WithAppID(ret.AppID),
		handler.WithRoleID(&ret.RoleID),
		handler.WithUserID(&ret.UserID),
		handler.WithResource(&ret.Resource),
		handler.WithMethod(&ret.Method),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateAuth(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAuth(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		handler.WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetAuth(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAuths(t *testing.T) {
	conds := &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		handler.WithOffset(0),
		handler.WithLimit(1),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetAuths(context.Background())
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteAuth(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		handler.WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteAuth(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetAuth(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestAuth(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupAuth(t)
	defer teardown(t)

	t.Run("creatAuth", creatAuth)
	t.Run("updateAuth", updateAuth)
	t.Run("getAuth", getAuth)
	t.Run("getAuths", getAuths)
	// t.Run("existAuth", existAuth)
	t.Run("deleteAuth", deleteAuth)
}
