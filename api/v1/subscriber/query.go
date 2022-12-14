package subscriber

import (
	"context"

	subscribermgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/subscriber"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/const"
	subscriber1 "github.com/NpoolPlatform/appuser-middleware/pkg/subscriber"

	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetSubscriber(ctx context.Context, in *npool.GetSubscriberRequest) (*npool.GetSubscriberResponse, error) {
	if _, err := uuid.Parse(in.GetID()); err != nil {
		return &npool.GetSubscriberResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := subscriber1.GetSubscriber(ctx, in.GetID())
	if err != nil {
		return &npool.GetSubscriberResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSubscriberResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSubscriberes(ctx context.Context, in *npool.GetSubscriberesRequest) (*npool.GetSubscriberesResponse, error) {
	conds := in.GetConds()
	if conds == nil {
		conds = &subscribermgrpb.Conds{}
	}

	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			return &npool.GetSubscriberesResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.AppID != nil {
		if _, err := uuid.Parse(conds.GetAppID().GetValue()); err != nil {
			return &npool.GetSubscriberesResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if conds.EmailAddress != nil && conds.GetEmailAddress().GetValue() == "" {
		return &npool.GetSubscriberesResponse{}, status.Error(codes.InvalidArgument, "EmailAddress is invalid")
	}

	limit := constant.DefaultRowLimit
	if in.GetLimit() > 0 {
		limit = in.GetLimit()
	}

	infos, total, err := subscriber1.GetSubscriberes(ctx, conds, in.GetOffset(), limit)
	if err != nil {
		return &npool.GetSubscriberesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSubscriberesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
