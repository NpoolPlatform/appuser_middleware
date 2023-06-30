package kyc

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
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	app "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	user "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user"
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
	ret = npool.Kyc{
		ID:              uuid.NewString(),
		AppID:           uuid.NewString(),
		UserID:          uuid.NewString(),
		DocumentType:    basetypes.KycDocumentType_IDCard,
		DocumentTypeStr: basetypes.KycDocumentType_IDCard.String(),
		IDNumber:        uuid.NewString(),
		FrontImg:        uuid.NewString(),
		BackImg:         uuid.NewString(),
		SelfieImg:       uuid.NewString(),
		EntityType:      basetypes.KycEntityType_Individual,
		EntityTypeStr:   basetypes.KycEntityType_Individual.String(),
		ReviewID:        uuid.NewString(),
		State:           basetypes.KycState_Reviewing,
		StateStr:        basetypes.KycState_Reviewing.String(),
	}
)

func setupKyc(t *testing.T) func(*testing.T) {
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

	ret.PhoneNO = fmt.Sprintf("+86%v", rand.Intn(100000000)+1000000)           //nolint
	ret.EmailAddress = fmt.Sprintf("%v@hhh.ccc", rand.Intn(100000000)+2000000) //nolint
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

	return func(*testing.T) {
		_, _ = ah.DeleteApp(context.Background())
		_, _ = uh.DeleteUser(context.Background())
	}
}

func createKyc(t *testing.T) {
	req := npool.KycReq{
		ID:           &ret.ID,
		AppID:        &ret.AppID,
		UserID:       &ret.UserID,
		DocumentType: &ret.DocumentType,
		IDNumber:     &ret.IDNumber,
		FrontImg:     &ret.FrontImg,
		BackImg:      &ret.BackImg,
		SelfieImg:    &ret.SelfieImg,
		EntityType:   &ret.EntityType,
		ReviewID:     &ret.ReviewID,
		State:        &ret.State,
	}
	info, err := CreateKyc(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateKyc(t *testing.T) {
	ret.IDNumber = uuid.NewString()
	ret.FrontImg = uuid.NewString()
	ret.State = basetypes.KycState_Rejected
	ret.StateStr = basetypes.KycState_Rejected.String()

	req := npool.KycReq{
		ID:           &ret.ID,
		DocumentType: &ret.DocumentType,
		IDNumber:     &ret.IDNumber,
		FrontImg:     &ret.FrontImg,
		BackImg:      &ret.BackImg,
		SelfieImg:    &ret.SelfieImg,
		EntityType:   &ret.EntityType,
		ReviewID:     &ret.ReviewID,
		State:        &ret.State,
	}
	info, err := UpdateKyc(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getKyc(t *testing.T) {
	info, err := GetKyc(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getKycs(t *testing.T) {
	_, total, err := GetKycs(context.Background(), &npool.Conds{
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, total, 0)
	}
}

func deleteKyc(t *testing.T) {
	info, err := DeleteKyc(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetKyc(context.Background(), ret.ID)
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

	teardown := setupKyc(t)
	defer teardown(t)

	t.Run("createKyc", createKyc)
	t.Run("updateKyc", updateKyc)
	t.Run("getKyc", getKyc)
	t.Run("getKycs", getKycs)
	t.Run("deleteKyc", deleteKyc)
}
