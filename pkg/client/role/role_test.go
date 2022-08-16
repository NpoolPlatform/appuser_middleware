package role

import (
	"bou.ke/monkey"
	"context"
	"fmt"
	approlecli "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/approle"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approle"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"strconv"
	"testing"

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
	roleInfo = approle.AppRole{
		ID:          uuid.NewString(),
		AppID:       uuid.NewString(),
		CreatedBy:   uuid.NewString(),
		Role:        uuid.NewString(),
		Description: uuid.NewString(),
		Default:     false,
	}
)

func create(t *testing.T) {
	appRoleReq := approle.AppRoleReq{
		ID:          &roleInfo.ID,
		AppID:       &roleInfo.AppID,
		CreatedBy:   &roleInfo.CreatedBy,
		Role:        &roleInfo.Role,
		Description: &roleInfo.Description,
		Default:     &roleInfo.Default,
	}
	_, err := approlecli.Create(context.Background(), &appRoleReq)
	if !assert.Nil(t, err) {
		return
	}
}
func getRoles(t *testing.T) {
	_, total, err := GetRoles(context.Background(), roleInfo.GetAppID(), 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, total, 0)
	}
}

func getManyRoles(t *testing.T) {
	_, total, err := GetManyRoles(context.Background(), []string{roleInfo.GetAppID()})
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
	t.Run("getRoles", getRoles)
	t.Run("getManyRoles", getManyRoles)
}
