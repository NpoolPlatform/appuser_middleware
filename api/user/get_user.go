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

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call middleware GetUserInfo")
	info, err := mw.GetUserInfo(ctx, appID, userID)
	if err != nil {
		logger.Sugar().Errorw("fail get app infos: %v", err)
		return &user.GetUserInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	resp, err := InfoRowToObject(info)
	return &user.GetUserInfoResponse{
		Info: resp,
	}, nil

}

func (s *Service) GetUserInfos(ctx context.Context, in *user.GetUserInfosRequest) (*user.GetUserInfosResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUserInfo")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call middleware GetUserInfos")
	infos, err := mw.GetUserInfos(ctx, appID, in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("fail get app infos: %v", err)
		return &user.GetUserInfosResponse{}, status.Error(codes.Internal, err.Error())
	}

	var userInfos []*user.AppUserInfo
	for _, val := range infos {
		info, err := InfoRowToObject(val)
		if err != nil {
			return nil, err
		}
		userInfos = append(userInfos, info)
	}

	return &user.GetUserInfosResponse{
		Infos: userInfos,
	}, nil

}
