package history

import (
	"context"
	"encoding/json"
	"fmt"

	history1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user/login/history"
	pubsubnotif "github.com/NpoolPlatform/appuser-middleware/pkg/pubsub/user/login/notif"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	historymwpb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/login/history"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/go-resty/resty/v2"
)

func getIPLocation(ip string) (string, error) {
	type resp struct {
		Error   bool   `json:"error"`
		City    string `json:"city"`
		Country string `json:"country_name"`
		IP      string `json:"ip"`
		Reason  string `json:"reason"`
	}

	r, err := resty.
		New().
		R().
		SetResult(&resp{}).
		Get(fmt.Sprintf("https://ipapi.co/%v/json", ip))
	if err != nil {
		return "", err
	}

	rc, ok := r.Result().(*resp)
	if rc.Error {
		return "", fmt.Errorf("%v", rc.Reason)
	}
	if !ok {
		return "", fmt.Errorf("invalid response")
	}
	return fmt.Sprintf("%v, %v", rc.City, rc.Country), nil
}

func createHistory(ctx context.Context, req *historymwpb.HistoryReq) error {
	if req.ClientIP == nil {
		return fmt.Errorf("invalid client ip")
	}
	conds := &historymwpb.Conds{
		Location: &basetypes.StringVal{Op: cruder.NEQ, Value: ""},
		ClientIP: &basetypes.StringVal{Op: cruder.EQ, Value: *req.ClientIP},
	}
	handler, err := history1.NewHandler(
		ctx,
		history1.WithConds(conds),
		history1.WithOffset(0),
		history1.WithLimit(1),
	)
	if err != nil {
		return err
	}
	infos, _, err := handler.GetHistories(ctx)
	if err != nil {
		return err
	}

	if len(infos) > 0 && infos[0].CreatedAt >= uint32(1685946444) { //nolint
		req.Location = &infos[0].Location
	} else if loc, err := getIPLocation(*req.ClientIP); err == nil {
		req.Location = &loc
		// send notif when new location detected
		pubsubnotif.NotifyNewDevice(req)
	}

	handler, err = history1.NewHandler(
		ctx,
		history1.WithID(req.ID),
		history1.WithAppID(req.GetAppID()),
		history1.WithUserID(req.GetUserID()),
		history1.WithClientIP(req.ClientIP),
		history1.WithUserAgent(req.UserAgent),
		history1.WithLocation(req.Location),
		history1.WithLoginType(req.LoginType),
	)
	if err != nil {
		return err
	}
	if _, err := handler.CreateHistory(ctx); err != nil {
		return err
	}

	return nil
}

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

	go func() {
		if err := createHistory(ctx, _req); err != nil {
			logger.Sugar().Errorw(
				"Apply",
				"Req", _req,
				"Error", err,
			)
		}
	}()

	return nil
}
