package subscriber

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/subscriber"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"

	mgrcli "github.com/NpoolPlatform/appuser-manager/pkg/client/subscriber"
)

func UpdateSubscriber(ctx context.Context, in *mgrpb.SubscriberReq) (*npool.Subscriber, error) {
	info, err := mgrcli.UpdateSubscriber(ctx, in)
	if err != nil {
		return nil, err
	}

	return GetSubscriber(ctx, info.ID)
}
