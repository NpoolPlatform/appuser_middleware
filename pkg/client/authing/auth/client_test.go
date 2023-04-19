package auth

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/auth"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	app "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	role "github.com/NpoolPlatform/appuser-middleware/pkg/mw/role"
	roleuser "github.com/NpoolPlatform/appuser-middleware/pkg/mw/role/user"
	user "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user"
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
	ret = npool.Auth{
		ID:       uuid.NewString(),
		AppID:    uuid.NewString(),
		Resource: uuid.NewString(),
		Method:   "POST",
	}
	roleID = uuid.NewString()
	userID = uuid.NewString()
)

func setupAuth(t *testing.T) func(*testing.T) {
	ret.AppName = ret.AppID
	ret.UserID = userID
	ret.RoleID = uuid.UUID{}.String()

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
		role.WithID(&roleID),
		role.WithAppID(ret.GetAppID()),
		role.WithCreatedBy(&ret.UserID),
		role.WithRole(&roleID),
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

	ruh, err := roleuser.NewHandler(
		context.Background(),
		roleuser.WithAppID(ret.GetAppID()),
		roleuser.WithRoleID(&roleID),
		roleuser.WithUserID(&userID),
	)
	assert.Nil(t, err)
	assert.NotNil(t, ruh)
	roleuser1, err := ruh.CreateUser(context.Background())
	assert.Nil(t, err)
	assert.NotNil(t, roleuser1)

	return func(*testing.T) {
		_, _ = ah.DeleteApp(context.Background())
		_, _ = rh.DeleteRole(context.Background())
		_, _ = uh.DeleteUser(context.Background())
		_, _ = ruh.DeleteUser(context.Background())
	}
}

func existUserFalseAuth(t *testing.T) {
	exist, err := ExistAuth(context.Background(), ret.AppID, &ret.UserID, ret.Resource, ret.Method)
	assert.Nil(t, err)
	assert.Equal(t, false, exist)
}

func existAppTrueAuth(t *testing.T) {
	exist, err := ExistAuth(context.Background(), ret.AppID, nil, ret.Resource, ret.Method)
	assert.Nil(t, err)
	assert.Equal(t, true, exist)
}

func existAppFalseAuth(t *testing.T) {
	exist, err := ExistAuth(context.Background(), ret.AppID, nil, ret.Resource, ret.Method)
	assert.Nil(t, err)
	assert.Equal(t, false, exist)
}

func createUserAuth(t *testing.T) {
	req := &npool.AuthReq{
		ID:       &ret.ID,
		AppID:    &ret.AppID,
		UserID:   &ret.UserID,
		Resource: &ret.Resource,
		Method:   &ret.Method,
	}
	info, err := CreateAuth(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, &ret, info)
	}
}

func getAuth(t *testing.T) {
	info, err := GetAuth(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getAuths(t *testing.T) {
	infos, _, err := GetAuths(context.Background(), &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}, 0, int32(1))
	if assert.Nil(t, err) {
		assert.Equal(t, 1, len(infos))
		assert.Equal(t, &ret, infos[0])
	}
}

func existUserTrueAuth(t *testing.T) {
	exist, err := ExistAuth(context.Background(), ret.AppID, &userID, ret.Resource, ret.Method)
	assert.Nil(t, err)
	assert.Equal(t, true, exist)
}

func deleteAuth(t *testing.T) {
	info, err := DeleteAuth(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}

	info, err = GetAuth(context.Background(), ret.ID)
	assert.Nil(t, err)
	assert.Nil(t, info)

	exist, err := ExistAuth(context.Background(), ret.AppID, &userID, ret.Resource, ret.Method)
	assert.Nil(t, err)
	assert.Equal(t, false, exist)

	exist, err = ExistAuth(context.Background(), ret.AppID, nil, ret.Resource, ret.Method)
	assert.Nil(t, err)
	assert.Equal(t, false, exist)
}

func createRoleAuth(t *testing.T) {
	ret.ID = uuid.NewString()
	ret.Resource = uuid.NewString()
	ret.RoleID = roleID
	ret.RoleName = roleID
	ret.UserID = uuid.UUID{}.String()
	ret.EmailAddress = ""
	ret.PhoneNO = ""

	req := &npool.AuthReq{
		ID:       &ret.ID,
		AppID:    &ret.AppID,
		RoleID:   &roleID,
		Resource: &ret.Resource,
		Method:   &ret.Method,
	}
	info, err := CreateAuth(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, &ret, info)
	}
}

func createAppAuth(t *testing.T) {
	ret.ID = uuid.NewString()
	ret.Resource = uuid.NewString()
	ret.RoleID = uuid.UUID{}.String()
	ret.RoleName = ""
	ret.UserID = uuid.UUID{}.String()
	ret.EmailAddress = ""
	ret.PhoneNO = ""

	req := &npool.AuthReq{
		ID:       &ret.ID,
		AppID:    &ret.AppID,
		Resource: &ret.Resource,
		Method:   &ret.Method,
	}
	info, err := CreateAuth(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		assert.Equal(t, &ret, info)
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

	teardown := setupAuth(t)
	defer teardown(t)

	t.Run("existUserFalseAuth", existUserFalseAuth)
	t.Run("existAppFalseAuth", existAppFalseAuth)

	t.Run("createUserAuth", createUserAuth)
	t.Run("getAuth", getAuth)
	t.Run("getAuths", getAuths)
	t.Run("existUserTrueAuth", existUserTrueAuth)
	t.Run("existAppFalseAuth", existAppFalseAuth)
	t.Run("deleteAuth", deleteAuth)

	t.Run("createRoleAuth", createRoleAuth)
	t.Run("getAuth", getAuth)
	t.Run("getAuths", getAuths)
	t.Run("existUserTrueAuth", existUserTrueAuth)
	t.Run("existAppFalseAuth", existAppFalseAuth)
	ret.UserID = uuid.UUID{}.String()
	t.Run("deleteAuth", deleteAuth)

	t.Run("createAppAuth", createAppAuth)
	t.Run("getAuth", getAuth)
	t.Run("getAuths", getAuths)
	t.Run("existUserTrueAuth", existUserTrueAuth)
	t.Run("existAppTrueAuth", existAppTrueAuth)
	ret.UserID = uuid.UUID{}.String()
	t.Run("deleteAuth", deleteAuth)
}
