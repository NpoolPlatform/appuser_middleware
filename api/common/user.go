package common

import (
	"context"

	user "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user"
)

func ExistUser(ctx context.Context, appID, userID string) (bool, error) {
	handler, err := user.NewHandler(
		ctx,
		user.WithAppID(appID),
		user.WithID(&userID),
	)
	if err != nil {
		return false, err
	}
	return handler.ExistUserConds(ctx)
}
