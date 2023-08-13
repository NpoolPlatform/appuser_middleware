package role

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"
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
	ret = npool.Role{
		ID:          uuid.NewString(),
		AppID:       uuid.NewString(),
		CreatedBy:   uuid.NewString(),
		Role:        uuid.NewString(),
		Description: uuid.NewString(),
		Default:     false,
	}
)

func setupRole(t *testing.T) func(*testing.T) {
	ret.AppName = ret.AppID

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

	return func(*testing.T) {
		_, _ = ah.DeleteApp(context.Background())
	}
}

func createRole(t *testing.T) {
	req := npool.RoleReq{
		ID:          &ret.ID,
		AppID:       &ret.AppID,
		CreatedBy:   &ret.CreatedBy,
		Role:        &ret.Role,
		Description: &ret.Description,
		Default:     &ret.Default,
	}
	info, err := CreateRole(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateRole(t *testing.T) {
	ret.Default = true
	ret.Role = uuid.NewString()
	req := npool.RoleReq{
		ID:          &ret.ID,
		AppID:       &ret.AppID,
		CreatedBy:   &ret.CreatedBy,
		Role:        &ret.Role,
		Description: &ret.Description,
		Default:     &ret.Default,
	}
	info, err := UpdateRole(context.Background(), &req)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getRole(t *testing.T) {
	info, err := GetRole(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getRoles(t *testing.T) {
	_, total, err := GetRoles(context.Background(), &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, total, 0)
	}
}

func deleteRole(t *testing.T) {
	info, err := DeleteRole(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetRole(context.Background(), ret.ID)
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

	teardown := setupRole(t)
	defer teardown(t)

	t.Run("createRole", createRole)
	t.Run("updateRole", updateRole)
	t.Run("getRole", getRole)
	t.Run("getRoles", getRoles)
	t.Run("deleteRole", deleteRole)
}
