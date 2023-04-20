package history

import (
	"context"
	"encoding/json"
	"fmt"

	history1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user/login/history"
	historymwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/login/history"
)

func Prepare(body string) (interface{}, error) {
	req := historymwpb.HistoryReq{}
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return nil, err
	}
	return &req, nil
}

func Apply(ctx context.Context, req interface{}) error {
	_req, ok := req.(*historymwpb.HistoryReq)
	if !ok {
		return fmt.Errorf("invalid request")
	}
	handler, err := history1.NewHandler(
		ctx,
		history1.WithID(_req.ID),
		history1.WithAppID(_req.GetAppID()),
		history1.WithUserID(_req.GetUserID()),
		history1.WithClientIP(_req.ClientIP),
		history1.WithUserAgent(_req.UserAgent),
		history1.WithLocation(_req.Location),
		history1.WithLoginType(_req.LoginType),
	)
	if err != nil {
		return err
	}
	if _, err := handler.CreateHistory(ctx); err != nil {
		return err
	}
	return nil
}
