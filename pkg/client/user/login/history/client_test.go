package history

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/login/history"
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
	ret = npool.History{
		ID:           uuid.NewString(),
		AppID:        uuid.NewString(),
		UserID:       uuid.NewString(),
		ClientIP:     "192.168.1.2",
		UserAgent:    uuid.NewString(),
		Location:     "Shanghai",
		LoginType:    basetypes.LoginType_FreshLogin,
		LoginTypeStr: basetypes.LoginType_FreshLogin.String(),
	}
)

func createHistory(t *testing.T) {
	req := &npool.HistoryReq{
		ID:        &ret.ID,
		AppID:     &ret.AppID,
		UserID:    &ret.UserID,
		ClientIP:  &ret.ClientIP,
		UserAgent: &ret.UserAgent,
		Location:  &ret.Location,
		LoginType: &ret.LoginType,
	}

	info, err := CreateHistory(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		info.AppName = ret.AppName
		info.AppLogo = ret.AppLogo
		assert.Equal(t, info, &ret)
	}
}

func getHistory(t *testing.T) {
	info, err := GetHistory(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getHistories(t *testing.T) {
	infos, _, err := GetHistories(context.Background(), &npool.Conds{
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		ClientIP:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.ClientIP},
		UserAgent: &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserAgent},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
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
	t.Run("createHistory", createHistory)
	t.Run("getHistory", getHistory)
	t.Run("getHistories", getHistories)
}
