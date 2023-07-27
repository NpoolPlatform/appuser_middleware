package notif

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/pubsub"
	loginhispb "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/login/history"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func NotifyNewDevice(in *loginhispb.HistoryReq) {
	if err := pubsub.WithPublisher(func(publisher *pubsub.Publisher) error {
		return publisher.Update(
			basetypes.MsgID_CreateNewLoginReq.String(),
			nil,
			nil,
			nil,
			in,
		)
	}); err != nil {
		logger.Sugar().Errorw(
			"notifyNewDevice",
			"AppID", in.AppID,
			"UserID", in.UserID,
			"ClientIP", in.ClientIP,
			"UserAgent", in.UserAgent,
			"Location", in.Location,
			"Error", err,
		)
	}
}
