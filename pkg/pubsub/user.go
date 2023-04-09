package pubsub

import (
	"context"
	"encoding/json"

	user "github.com/NpoolPlatform/appuser-middleware/pkg/user"
	eventmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"
)

func prepareIncreaseUserActionCredits(body string) (interface{}, error) {
	req := []*eventmwpb.Credit{}
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return nil, err
	}
	return req, nil
}

func handleIncreaseUserActionCredits(ctx context.Context, req interface{}) error {
	credits := req.([]*eventmwpb.Credit)

	// TODO: here we should run in transaction
	for _, credit := range credits {
		handler, err := user.NewHandler(
			ctx,
			user.WithAppID(credit.GetAppID()),
			user.WithID(&credit.UserID),
			user.WithActionCredits(&credit.Credits),
		)
		if err != nil {
			return err
		}
		if _, err := handler.UpdateUser(ctx); err != nil {
			return err
		}
	}

	return nil
}
