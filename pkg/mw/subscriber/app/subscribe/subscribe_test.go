package appsubscribe

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber/app/subscribe"
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
	ret = npool.AppSubscribe{
		ID:             uuid.NewString(),
		AppID:          uuid.NewString(),
		SubscribeAppID: uuid.NewString(),
	}
)

func setupAppSubscribe(t *testing.T) func(*testing.T) {
	ah, err := app.NewHandler(
		context.Background(),
		app.WithID(&ret.AppID),
		app.WithCreatedBy(uuid.NewString()),
		app.WithName(&ret.AppID),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ah)
	app1, err := ah.CreateApp(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, app1)

	ah1, err := app.NewHandler(
		context.Background(),
		app.WithID(&ret.SubscribeAppID),
		app.WithCreatedBy(uuid.NewString()),
		app.WithName(&ret.SubscribeAppID),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ah)
	app2, err := ah1.CreateApp(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, app2)

	ret.AppName = ret.AppID
	ret.SubscribeAppName = ret.SubscribeAppID

	return func(*testing.T) {
		_, _ = ah.DeleteApp(context.Background())
		_, _ = ah1.DeleteApp(context.Background())
	}
}

func createAppSubscribe(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(ret.GetAppID()),
		WithSubscribeAppID(ret.SubscribeAppID),
	)
	assert.Nil(t, err)

	info, err := handler.CreateAppSubscribe(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getAppSubscribe(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetAppSubscribe(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAppSubscribes(t *testing.T) {
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

	infos, _, err := handler.GetAppSubscribes(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func existAppSubscribe(t *testing.T) {
	conds := &npool.Conds{
		ID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
	)
	assert.Nil(t, err)

	exist, err := handler.ExistAppSubscribeConds(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteAppSubscribe(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteAppSubscribe(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetAppSubscribe(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestAppSubscribe(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupAppSubscribe(t)
	defer teardown(t)

	t.Run("createAppSubscribe", createAppSubscribe)
	t.Run("getAppSubscribe", getAppSubscribe)
	t.Run("getAppSubscribes", getAppSubscribes)
	t.Run("existAppSubscribs", existAppSubscribe)
	t.Run("deleteAppSubscribe", deleteAppSubscribe)
}
