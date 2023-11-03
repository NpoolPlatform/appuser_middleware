package role

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/appuser-middleware/pkg/testinit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

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
		EntID:       uuid.NewString(),
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
		app.WithEntID(&ret.AppID, true),
		app.WithCreatedBy(&ret.EntID, true),
		app.WithName(&ret.AppName, true),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ah)
	app1, err := ah.CreateApp(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, app1)

	ah, err = app.NewHandler(
		context.Background(),
		app.WithID(&app1.ID, true),
	)
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = ah.DeleteApp(context.Background())
	}
}

func creatRole(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithCreatedBy(&ret.CreatedBy, true),
		WithRole(&ret.Role, true),
		WithDescription(&ret.Description, true),
		WithDefault(&ret.Default, true),
		WithGenesis(&ret.Genesis, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateRole(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateRole(t *testing.T) {
	ret.Role = uuid.NewString()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithRole(&ret.Role, true),
		WithDescription(&ret.Description, true),
		WithDefault(&ret.Default, true),
		WithGenesis(&ret.Genesis, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateRole(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getRole(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
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
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetRoles(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteRole(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteRole(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetRole(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestRole(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("creatRole", creatRole)
	t.Run("updateRole", updateRole)
	t.Run("getRole", getRole)
	t.Run("getRoles", getRoles)
	t.Run("deleteRole", deleteRole)
}
