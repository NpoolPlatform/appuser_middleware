package kyc

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/appuser-middleware/pkg/testinit"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	app "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	user "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user"
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

	return func(*testing.T) {
		_, _ = ah.DeleteApp(context.Background())
		_, _ = uh.DeleteUser(context.Background())
	}
}

func createKyc(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithAppID(ret.GetAppID()),
		WithUserID(ret.GetUserID()),
		WithDocumentType(&ret.DocumentType),
		WithIDNumber(&ret.IDNumber),
		WithFrontImg(&ret.FrontImg),
		WithBackImg(&ret.BackImg),
		WithSelfieImg(&ret.SelfieImg),
		WithEntityType(&ret.EntityType),
		WithReviewID(&ret.ReviewID),
	)
	assert.Nil(t, err)

	info, err := handler.CreateKyc(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateKyc(t *testing.T) {
	ret.State = basetypes.KycState_Approved
	ret.StateStr = basetypes.KycState_Approved.String()
	ret.FrontImg = uuid.NewString()
	ret.BackImg = uuid.NewString()
	ret.SelfieImg = uuid.NewString()

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithDocumentType(&ret.DocumentType),
		WithIDNumber(&ret.IDNumber),
		WithFrontImg(&ret.FrontImg),
		WithBackImg(&ret.BackImg),
		WithSelfieImg(&ret.SelfieImg),
		WithEntityType(&ret.EntityType),
		WithReviewID(&ret.ReviewID),
		WithState(&ret.State),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateKyc(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getKyc(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetKyc(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getKycs(t *testing.T) {
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

	infos, _, err := handler.GetKycs(context.Background())
	if !assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteKyc(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteKyc(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetKyc(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestKyc(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setupKyc(t)
	defer teardown(t)

	t.Run("createKyc", createKyc)
	t.Run("updateKyc", updateKyc)
	t.Run("getKyc", getKyc)
	t.Run("getKycs", getKycs)
	t.Run("deleteKyc", deleteKyc)
}
