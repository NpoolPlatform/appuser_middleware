package authing

import (
	"context"
	"fmt"

	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	authcli "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/authing/auth"
	authhistorycli "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/authing/history"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/authing/auth"
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/authing/history"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/appuser-middleware/pkg/testinit"
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
	authInfo = auth.Auth{
		ID:       uuid.NewString(),
		AppID:    uuid.NewString(),
		RoleID:   uuid.NewString(),
		UserID:   uuid.NewString(),
		Resource: uuid.NewString(),
		Method:   uuid.NewString(),
	}

	historyInfo = history.History{
		ID:        uuid.NewString(),
		AppID:     uuid.NewString(),
		UserID:    uuid.NewString(),
		Resource:  uuid.NewString(),
		Method:    uuid.NewString(),
		Allowed:   true,
		CreatedAt: 0,
	}
)

func create(t *testing.T) {
	appAuthReq := auth.AuthReq{
		ID:       &authInfo.ID,
		AppID:    &authInfo.AppID,
		RoleID:   &authInfo.RoleID,
		UserID:   &authInfo.UserID,
		Resource: &authInfo.Resource,
		Method:   &authInfo.Method,
	}
	_, err := authcli.Create(context.Background(), &appAuthReq)
	if !assert.Nil(t, err) {
		return
	}
}

func createH(t *testing.T) {
	historyReq := history.HistoryReq{
		ID:       &historyInfo.ID,
		AppID:    &historyInfo.AppID,
		UserID:   &historyInfo.UserID,
		Resource: &historyInfo.Resource,
		Method:   &historyInfo.Method,
		Allowed:  &historyInfo.Allowed,
	}
	_, err := authhistorycli.Create(context.Background(), &historyReq)
	if !assert.Nil(t, err) {
		return
	}
}

func existAuth(t *testing.T) {
	exist, err := ExistAuth(context.Background(), authInfo.GetAppID(), authInfo.GetUserID(), authInfo.GetResource(), authInfo.Method)
	if assert.Nil(t, err) {
		assert.NotEqual(t, exist, true)
	}
}

func getAuths(t *testing.T) {
	_, total, err := GetAuths(context.Background(), authInfo.GetAppID(), 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, total, 0)
	}
}

func getHistories(t *testing.T) {
	_, total, err := GetHistories(context.Background(), authInfo.GetAppID(), 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, total, 0)
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

	t.Run("create", create)
	t.Run("createH", createH)
	t.Run("existAuth", existAuth)
	t.Run("getAuths", getAuths)
	t.Run("getHistories", getHistories)
}
