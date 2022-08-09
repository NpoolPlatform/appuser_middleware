////nolint:dupl
package admin

//
//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"os"
//	"strconv"
//	"testing"
//
//	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
//	"github.com/google/uuid"
//	"github.com/stretchr/testify/assert"
//
//	"github.com/NpoolPlatform/appuser-middleware/pkg/testinit"
//)
//
//func init() {
//	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
//		return
//	}
//	if err := testinit.Init(); err != nil {
//		fmt.Printf("cannot init test stub: %v\n", err)
//	}
//}
//
//var (
//	uuidSlice     = []string{uuid.NewString()}
//	uuidSliceS, _ = json.Marshal(uuidSlice)
//	userInfo      = User{
//		ID:                                 uuid.New(),
//		AppID:                              uuid.New(),
//		EmailAddress:                       uuid.NewString(),
//		PhoneNO:                            uuid.NewString(),
//		ImportedFromAppID:                  uuid.New(),
//		CreatedAt:                          0,
//		Username:                           uuid.NewString(),
//		AddressFields:                      string(uuidSliceS),
//		Gender:                             uuid.NewString(),
//		PostalCode:                         uuid.NewString(),
//		Age:                                0,
//		Birthday:                           0,
//		Avatar:                             uuid.NewString(),
//		Organization:                       uuid.NewString(),
//		FirstName:                          uuid.NewString(),
//		LastName:                           uuid.NewString(),
//		IDNumber:                           uuid.NewString(),
//		SigninVerifyByGoogleAuthentication: 1,
//		GoogleAuthenticationVerified:       1,
//		Banned:                             false,
//		HasGoogleSecret:                    true,
//		Roles:                              []string{""},
//	}
//)
//
//func creatUser(t *testing.T) {
//	var (
//		id                = userInfo.ID.String()
//		appID             = userInfo.AppID.String()
//		importedFromAppID = userInfo.ImportedFromAppID.String()
//		boolVal           = true
//		strVal            = "AAA"
//		userReq           = npool.UserReq{
//			ID:                           &id,
//			AppID:                        &appID,
//			EmailAddress:                 &userInfo.EmailAddress,
//			PhoneNO:                      &userInfo.PhoneNO,
//			ImportedFromAppID:            &importedFromAppID,
//			Username:                     &userInfo.Username,
//			AddressFields:                uuidSlice,
//			Gender:                       &userInfo.Gender,
//			PostalCode:                   &userInfo.PostalCode,
//			Age:                          &userInfo.Age,
//			Birthday:                     &userInfo.Birthday,
//			Avatar:                       &userInfo.Avatar,
//			Organization:                 &userInfo.Organization,
//			FirstName:                    &userInfo.FirstName,
//			LastName:                     &userInfo.LastName,
//			IDNumber:                     &userInfo.IDNumber,
//			SigninVerifyByGoogleAuth:     &boolVal,
//			GoogleAuthenticationVerified: &boolVal,
//			PasswordHash:                 &strVal,
//			GoogleSecret:                 &appID,
//			ThirdPartyID:                 &strVal,
//			ThirdPartyUserID:             &strVal,
//			ThirdPartyUsername:           &strVal,
//			ThirdPartyUserAvatar:         &strVal,
//			Banned:                       &userInfo.Banned,
//			BanMessage:                   &userInfo.BanMessage,
//		}
//	)
//	info, err := CreateUser(context.Background(), &userReq)
//	if assert.Nil(t, err) {
//		info.CreatedAt = userInfo.CreatedAt
//		assert.Equal(t, info, &userInfo)
//	}
//}
//
//func updateUser(t *testing.T) {
//	var (
//		id                = userInfo.ID.String()
//		appID             = userInfo.AppID.String()
//		importedFromAppID = userInfo.ImportedFromAppID.String()
//		boolVal           = true
//		strVal            = "AAA"
//		userReq           = npool.UserReq{
//			ID:                           &id,
//			AppID:                        &appID,
//			EmailAddress:                 &userInfo.EmailAddress,
//			PhoneNO:                      &userInfo.PhoneNO,
//			ImportedFromAppID:            &importedFromAppID,
//			Username:                     &userInfo.Username,
//			AddressFields:                uuidSlice,
//			Gender:                       &userInfo.Gender,
//			PostalCode:                   &userInfo.PostalCode,
//			Age:                          &userInfo.Age,
//			Birthday:                     &userInfo.Birthday,
//			Avatar:                       &userInfo.Avatar,
//			Organization:                 &userInfo.Organization,
//			FirstName:                    &userInfo.FirstName,
//			LastName:                     &userInfo.LastName,
//			IDNumber:                     &userInfo.IDNumber,
//			SigninVerifyByGoogleAuth:     &boolVal,
//			GoogleAuthenticationVerified: &boolVal,
//			PasswordHash:                 &strVal,
//			GoogleSecret:                 &appID,
//			ThirdPartyID:                 &strVal,
//			ThirdPartyUserID:             &strVal,
//			ThirdPartyUsername:           &strVal,
//			ThirdPartyUserAvatar:         &strVal,
//			Banned:                       &userInfo.Banned,
//			BanMessage:                   &userInfo.BanMessage,
//		}
//	)
//	info, err := UpdateUser(context.Background(), &userReq)
//	if assert.Nil(t, err) {
//		info.CreatedAt = userInfo.CreatedAt
//		info.Roles = userInfo.Roles
//		assert.Equal(t, info, &userInfo)
//	}
//}
//
//func getUser(t *testing.T) {
//	info, err := GetUser(context.Background(), userInfo.AppID.String(), userInfo.ID.String())
//	if assert.Nil(t, err) {
//		info.CreatedAt = userInfo.CreatedAt
//		assert.Equal(t, info, &userInfo)
//	}
//}
//
//func getUsers(t *testing.T) {
//	infos, err := GetUsers(context.Background(), userInfo.AppID.String(), 0, 1)
//	if !assert.Nil(t, err) {
//		assert.NotEqual(t, len(infos), 0)
//	}
//}
//
//func getManyUsers(t *testing.T) {
//	infos, err := GetManyUsers(context.Background(), []string{userInfo.ID.String()})
//	if !assert.Nil(t, err) {
//		infos[0].CreatedAt = userInfo.CreatedAt
//		assert.Equal(t, infos[0], &userInfo)
//	}
//}
//
//func TestMainOrder(t *testing.T) {
//	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
//		return
//	}
//	t.Run("creatUser", creatUser)
//	t.Run("updateUser", updateUser)
//	t.Run("getUser", getUser)
//	t.Run("getUsers", getUsers)
//	t.Run("getManyUsers", getManyUsers)
//}
