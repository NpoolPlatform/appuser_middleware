package appsubscribe

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber/app/subscribe"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	app "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	"github.com/NpoolPlatform/appuser-middleware/pkg/testinit"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
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
	ret.AppName = ret.AppID
	ret.SubscribeAppName = ret.SubscribeAppID

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

	return func(*testing.T) {
		_, _ = ah.DeleteApp(context.Background())
		_, _ = ah1.DeleteApp(context.Background())
	}
}

func createAppSubscribe(t *testing.T) {
	req := npool.AppSubscribeReq{
		ID:             &ret.ID,
		AppID:          &ret.AppID,
		SubscribeAppID: &ret.SubscribeAppID,
	}
	info, err := CreateAppSubscribe(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getAppSubscribe(t *testing.T) {
	info, err := GetAppSubscribe(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAppSubscribes(t *testing.T) {
	_, total, err := GetAppSubscribes(context.Background(), &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, total, 0)
	}
}

func existAppSubscribe(t *testing.T) {
	exist, err := ExistAppSubscribeConds(context.Background(), &npool.Conds{
		ID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteAppSubscribe(t *testing.T) {
	info, err := DeleteAppSubscribe(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetAppSubscribe(context.Background(), ret.ID)
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
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	teardown := setupAppSubscribe(t)
	defer teardown(t)

	t.Run("createAppSubscribe", createAppSubscribe)
	t.Run("getAppSubscribe", getAppSubscribe)
	t.Run("getAppSubscribes", getAppSubscribes)
	t.Run("existAppSubscribe", existAppSubscribe)
	t.Run("deleteAppSubscribe", deleteAppSubscribe)
}
