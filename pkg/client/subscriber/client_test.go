package subscriber

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	mgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/subscriber"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"

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
	ret = npool.Subscriber{
		ID:           uuid.NewString(),
		AppID:        uuid.NewString(),
		EmailAddress: uuid.NewString(),
	}

	req = &mgrpb.SubscriberReq{
		ID:           &ret.ID,
		AppID:        &ret.AppID,
		EmailAddress: &ret.EmailAddress,
	}
)

func createSubscriber(t *testing.T) {
	info, err := CreateSubscriber(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateSubscriber(t *testing.T) {
	registered := true

	req.Registered = &registered
	ret.Registered = registered

	info, err := UpdateSubscriber(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getSubscriber(t *testing.T) {
	info, err := GetSubscriber(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getSubscriberes(t *testing.T) {
	infos, _, err := GetSubscriberes(context.Background(), nil, 0, 1)
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteSubscriber(t *testing.T) {
	info, err := DeleteSubscriber(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	_, err = GetSubscriber(context.Background(), ret.ID)
	assert.NotNil(t, err)
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	t.Run("createSubscriber", createSubscriber)
	t.Run("updateSubscriber", updateSubscriber)
	t.Run("getSubscriber", getSubscriber)
	t.Run("getSubscriberes", getSubscriberes)
	t.Run("deleteSubscriber", deleteSubscriber)
}
