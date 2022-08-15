//nolint:nolintlint,dupl
package app

import (
	"context"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	mapp "github.com/NpoolPlatform/appuser-middleware/pkg/app"
	capp "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/app"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) GetApp(ctx context.Context, in *npool.GetAppRequest) (*npool.GetAppResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetApp")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetAppID())

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetApp", "error", err)
		return &npool.GetAppResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	span = commontracer.TraceInvoker(span, "app", "middleware", "GetApp")

	info, err := mapp.GetApp(ctx, in.GetAppID())
	if err != nil {
		logger.Sugar().Errorw("GetApp", "error", err)
		return &npool.GetAppResponse{}, status.Error(codes.Internal, "fail get app")
	}

	ginfo, err := capp.Ent2Grpc(info)
	if err != nil {
		logger.Sugar().Errorw("GetApp", "error", err)
		return &npool.GetAppResponse{}, status.Error(codes.Internal, "invalid value")
	}

	return &npool.GetAppResponse{
		Info: ginfo,
	}, nil
}

func (s *Server) GetApps(ctx context.Context, in *npool.GetAppsRequest) (*npool.GetAppsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetApps")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "app", "middleware", "GetApps")

	infos, err := mapp.GetApps(ctx, in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetApps", "error", err)
		return &npool.GetAppsResponse{}, status.Error(codes.Internal, "fail get apps")
	}

	resp := []*npool.App{}
	for _, val := range infos {
		ginfo, err := capp.Ent2Grpc(val)
		if err != nil {
			logger.Sugar().Errorw("GetApps", "error", err)
			return &npool.GetAppsResponse{}, status.Error(codes.Internal, "invalid value")
		}
		resp = append(resp, ginfo)
	}

	return &npool.GetAppsResponse{
		Infos: resp,
		Total: uint32(len(resp)),
	}, nil
}

func (s *Server) GetUserApps(ctx context.Context, in *npool.GetUserAppsRequest) (*npool.GetUserAppsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUserApps")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.String("UserID", in.GetUserID()))

	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("GetUserApps", "error", err)
		return &npool.GetUserAppsResponse{}, status.Error(codes.InvalidArgument, "UserID is invalid")
	}

	span = commontracer.TraceInvoker(span, "app", "middleware", "GetUserApps")

	infos, total, err := mapp.GetUserApps(ctx, in.GetUserID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetUserApps", "error", err)
		return &npool.GetUserAppsResponse{}, status.Error(codes.Internal, "fail get user apps")
	}

	resp := []*npool.App{}
	for _, val := range infos {
		ginfo, err := capp.Ent2Grpc(val)
		if err != nil {
			logger.Sugar().Errorw("GetUserApps", "error", err)
			return &npool.GetUserAppsResponse{}, status.Error(codes.Internal, "invalid value")
		}
		resp = append(resp, ginfo)
	}

	return &npool.GetUserAppsResponse{
		Infos: resp,
		Total: uint32(total),
	}, nil
}

func (s *Server) GetManyApps(ctx context.Context, in *npool.GetManyAppsRequest) (*npool.GetManyAppsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUserApps")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.StringSlice("IDs", in.GetIDs()))

	for _, val := range in.GetIDs() {
		if _, err := uuid.Parse(val); err != nil {
			logger.Sugar().Errorw("GetUserApps", "error", err)
			return &npool.GetManyAppsResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
		}
	}

	span = commontracer.TraceInvoker(span, "app", "middleware", "GetUserApps")

	infos, total, err := mapp.GetManyApps(ctx, in.GetIDs())
	if err != nil {
		logger.Sugar().Errorw("GetUserApps", "error", err)
		return &npool.GetManyAppsResponse{}, status.Error(codes.Internal, "fail get user apps")
	}

	resp := []*npool.App{}
	for _, val := range infos {
		ginfo, err := capp.Ent2Grpc(val)
		if err != nil {
			logger.Sugar().Errorw("GetUserApps", "error", err)
			return &npool.GetManyAppsResponse{}, status.Error(codes.Internal, "invalid value")
		}
		resp = append(resp, ginfo)
	}

	return &npool.GetManyAppsResponse{
		Infos: resp,
		Total: uint32(total),
	}, nil
}
