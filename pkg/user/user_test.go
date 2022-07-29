package user

import (
	"fmt"
	testinit "github.com/NpoolPlatform/appuser-manager/pkg/testinit"
	"os"
	"strconv"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

//}
//var entApp = ent.App{
//	ID:          uuid.New(),
//	CreatedBy:   uuid.New(),
//	Name:        uuid.New().String(),
//	Description: uuid.New().String(),
//	Logo:        uuid.New().String(),
//}
//
//var (
//	id        = entApp.ID.String()
//	createdBy = entApp.CreatedBy.String()
//	appInfo   = npool.AppReq{
//		ID:          &id,
//		CreatedBy:   &createdBy,
//		Name:        &entApp.Name,
//		Description: &entApp.Description,
//		Logo:        &entApp.Logo,
//	}
//)
//
//var info *ent.App
//
//func rowToObject(row *ent.App) *ent.App {
//	return &ent.App{
//		ID:          row.ID,
//		CreatedBy:   row.CreatedBy,
//		Name:        row.Name,
//		Logo:        row.Logo,
//		Description: row.Description,
//		CreatedAt:   row.CreatedAt,
//	}
//}

//func getUser(t *testing.T) {
//	var err error
//	info, err = GetUser(context.Background(), &appRoleInfo)
//	if assert.Nil(t, err) {
//		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
//			entAppRole.ID = info.ID
//		}
//		assert.Equal(t, rowToObject(info), &entAppRole)
//	}
//}
//
//func getUser(t *testing.T) {
//	var err error
//	info, err = GetUser(context.Background(), &appRoleInfo)
//	if assert.Nil(t, err) {
//		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
//			entAppRole.ID = info.ID
//		}
//		assert.Equal(t, rowToObject(info), &entAppRole)
//	}
//}
