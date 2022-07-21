package app

import (
	"context"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/google/uuid"

	appmw "github.com/NpoolPlatform/appuser-middleware/pkg/middleware/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/appusermw/app"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetAppInfo(ctx context.Context, in *app.GetAppInfoRequest) (*app.GetAppInfoResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppInfo")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Error("ID is invalid")
		return &app.GetAppInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	span.AddEvent("call middleware GetAppInfo")
	info, err := appmw.GetAppInfo(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("fail get app info: %v", err)
		return &app.GetAppInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	resp, err := appInfoRowToObject(info)
	if err != nil {
		return nil, err
	}
	return &app.GetAppInfoResponse{
		Info: resp,
	}, nil
}

func (s *Service) GetAppInfos(ctx context.Context, in *app.GetAppInfosRequest) (*app.GetAppInfosResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppInfos")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.AddEvent("call middleware GetAppInfos")
	infos, err := appmw.GetAppInfos(ctx, in.GetLimit(), in.GetOffset())
	if err != nil {
		logger.Sugar().Errorw("fail get app infos: %v", err)
		return &app.GetAppInfosResponse{}, status.Error(codes.Internal, err.Error())
	}

	resp := []*app.AppInfo{}

	for _, val := range infos {
		info, err := appInfoRowToObject(val)
		if err != nil {
			return nil, err
		}
		resp = append(resp, info)
	}

	return &app.GetAppInfosResponse{
		Infos: resp,
	}, nil
}

func (s *Service) GetAppInfosByCreator(ctx context.Context, in *app.GetAppInfosByCreatorRequest) (*app.GetAppInfosByCreatorResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppInfosByCreator")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	uid, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call middleware GetAppInfosByCreator")
	infos, err := appmw.GetAppInfosByCreator(ctx, uid, in.GetLimit(), in.GetOffset())
	if err != nil {
		logger.Sugar().Errorw("fail get app infos: %v", err)
		return &app.GetAppInfosByCreatorResponse{}, status.Error(codes.Internal, err.Error())
	}

	resp := []*app.AppInfo{}

	for _, val := range infos {
		info, err := appInfoRowToObject(val)
		if err != nil {
			return nil, err
		}
		resp = append(resp, info)
	}

	return &app.GetAppInfosByCreatorResponse{
		Infos: resp,
	}, nil
}
