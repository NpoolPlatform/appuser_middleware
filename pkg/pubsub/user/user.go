package user

import (
	"context"
	"encoding/json"
	"fmt"

	user1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user"
	eventmwpb "github.com/NpoolPlatform/message/npool/inspire/mw/v1/event"
)

func Prepare(body string) (interface{}, error) {
	req := []*eventmwpb.Credit{}
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return nil, err
	}
	return req, nil
}

func Apply(ctx context.Context, req interface{}) error {
	credits, ok := req.([]*eventmwpb.Credit)
	if !ok {
		return fmt.Errorf("invalid request")
	}

	// TODO: here we should run in transaction
	for _, credit := range credits {
		handler, err := user1.NewHandler(
			ctx,
			user1.WithAppID(credit.GetAppID()),
			user1.WithID(&credit.UserID),
			user1.WithActionCredits(&credit.Credits),
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
