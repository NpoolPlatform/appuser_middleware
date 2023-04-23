package history

import (
	"context"
	"encoding/json"
	"fmt"

	handler "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/handler"
	history1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/history"
	historymwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/history"
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

	_handler, err := history1.NewHandler(
		ctx,
		handler.WithID(_req.ID),
		handler.WithAppID(_req.GetAppID()),
		handler.WithUserID(_req.UserID),
		handler.WithResource(_req.Resource),
		handler.WithMethod(_req.Method),
		history1.WithAllowed(_req.Allowed),
	)
	if err != nil {
		return err
	}
	if _, err := _handler.CreateHistory(ctx); err != nil {
		return err
	}

	return nil
}
