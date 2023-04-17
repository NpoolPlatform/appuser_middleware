package subscriber

import (
	"context"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"

	mgrapi "github.com/NpoolPlatform/appuser-manager/api/subscriber"
	mgrcli "github.com/NpoolPlatform/appuser-manager/pkg/client/subscriber"
	mgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/subscriber"

	subscriber1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/subscriber"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateSubscriber(ctx context.Context, in *npool.CreateSubscriberRequest) (*npool.CreateSubscriberResponse, error) {
	if err := mgrapi.Validate(in.GetInfo()); err != nil {
		return &npool.CreateSubscriberResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := mgrcli.GetSubscriberOnly(ctx, &mgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetInfo().GetAppID(),
		},
		EmailAddress: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetInfo().GetEmailAddress(),
		},
	})
	if err != nil {
		return &npool.CreateSubscriberResponse{}, status.Error(codes.Internal, err.Error())
	}
	if info != nil {
		info1, err := subscriber1.GetSubscriber(ctx, info.ID)
		if err != nil {
			return &npool.CreateSubscriberResponse{}, status.Error(codes.Internal, err.Error())
		}
		return &npool.CreateSubscriberResponse{
			Info: info1,
		}, nil
	}

	info1, err := subscriber1.CreateSubscriber(ctx, in.GetInfo())
	if err != nil {
		return &npool.CreateSubscriberResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSubscriberResponse{
		Info: info1,
	}, nil
}
