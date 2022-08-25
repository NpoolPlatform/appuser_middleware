package kyc

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool"

	mgr "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/kyc"

	"bou.ke/monkey"
	kyccli "github.com/NpoolPlatform/appuser-manager/pkg/crud/v2/kyc"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"
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
	kycInfo = kyc.Kyc{
		ID:              uuid.NewString(),
		AppID:           uuid.NewString(),
		UserID:          uuid.NewString(),
		DocumentType:    mgr.KycDocumentType_IDCard,
		DocumentTypeStr: mgr.KycDocumentType_IDCard.String(),
		IDNumber:        uuid.NewString(),
		FrontImg:        uuid.NewString(),
		BackImg:         uuid.NewString(),
		SelfieImg:       uuid.NewString(),
		EntityType:      mgr.KycEntityType_Individual,
		EntityTypeStr:   mgr.KycEntityType_Individual.String(),
		ReviewID:        uuid.NewString(),
		StateStr:        mgr.KycState_Wait.String(),
	}
)

func create(t *testing.T) {
	appRoleReq := mgr.KycReq{
		ID:           &kycInfo.ID,
		AppID:        &kycInfo.AppID,
		UserID:       &kycInfo.UserID,
		DocumentType: &kycInfo.DocumentType,
		IDNumber:     &kycInfo.IDNumber,
		FrontImg:     &kycInfo.FrontImg,
		BackImg:      &kycInfo.BackImg,
		SelfieImg:    &kycInfo.SelfieImg,
		EntityType:   &kycInfo.EntityType,
		ReviewID:     &kycInfo.ReviewID,
	}
	_, err := kyccli.Create(context.Background(), &appRoleReq)
	if !assert.Nil(t, err) {
		return
	}
}

func getKyc(t *testing.T) {
	info, err := GetKyc(context.Background(), kycInfo.GetID())
	if assert.Nil(t, err) {
		kycInfo.CreatedAt = info.CreatedAt
		kycInfo.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &kycInfo)
	}
}

func getKycs(t *testing.T) {
	infos, _, err := GetKycs(context.Background(), &kyc.Conds{
		Conds: &mgr.Conds{
			ID: &npool.StringVal{
				Op:    cruder.EQ,
				Value: kycInfo.ID,
			},
		}}, 0, 1)
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

	t.Run("create", create)
	t.Run("getKyc", getKyc)
	t.Run("getKycs", getKycs)
}
