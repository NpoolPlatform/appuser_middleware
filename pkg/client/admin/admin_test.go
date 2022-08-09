package admin

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	approlecrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/approle"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	approlemgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approle"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"

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

func createGenesisUser(t *testing.T) {
	appID := uuid.NewString()
	role := uuid.NewString()
	userID := uuid.NewString()
	emailAddress := uuid.NewString()
	passwordHash := uuid.NewString()
	description := uuid.NewString()
	defaultVal := false
	_, err := approlecrud.Create(context.Background(), &approlemgrpb.AppRoleReq{
		AppID:       &appID,
		CreatedBy:   &userID,
		Role:        &role,
		Description: &description,
		Default:     &defaultVal,
	})
	if !assert.Nil(t, err) {
		return
	}

	user, err := CreateGenesisUser(context.Background(), appID, userID, role, emailAddress, passwordHash)
	if assert.Nil(t, err) {
		assert.Equal(t, user.ID, userID)
		assert.Equal(t, user.AppID, appID)
		assert.Equal(t, user.Roles[0], role)
		assert.Equal(t, user.EmailAddress, emailAddress)
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

	t.Run("createGenesisUser", createGenesisUser)
}
