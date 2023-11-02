package common

import (
	"context"
	"fmt"

	user "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user"
)

func ValidateUser(ctx context.Context, appID, userID string) error {
	handler, err := user.NewHandler(
		ctx,
		user.WithAppID(&appID, true),
		user.WithEntID(&userID, true),
	)
	if err != nil {
		return err
	}
	exist, err := handler.ExistUserConds(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("user not exists")
	}
	return nil
}
