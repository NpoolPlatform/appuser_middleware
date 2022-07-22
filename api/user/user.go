package user

import (
	"context"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	mw "github.com/NpoolPlatform/appuser-middleware/pkg/middleware/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/appusermw/user"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetUserInfo(ctx context.Context, in *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUserInfo")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	uid, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call middleware GetAppInfosByCreator")
	resp, err := mw.GetUserInfo(ctx, uid)
	if err != nil {
		logger.Sugar().Errorw("fail get app infos: %v", err)
		return &user.GetUserInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &user.GetUserInfoResponse{
		Info: resp,
	}, nil
}
