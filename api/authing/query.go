//nolint:dupl
package authing

import (
	"context"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	servicename "github.com/NpoolPlatform/appuser-middleware/pkg/servicename"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"

	tracer "github.com/NpoolPlatform/appuser-middleware/pkg/tracer/authing"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing"

	authing1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) ExistAuth(ctx context.Context, in *npool.ExistAuthRequest) (info *npool.ExistAuthResponse, err error) {
	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetRoles")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("ExistAuth", "AppID", in.GetAppID(), "error", err)
		return &npool.ExistAuthResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}
	if in.UserID != nil {
		if _, err := uuid.Parse(in.GetUserID()); err != nil {
			logger.Sugar().Errorw("ExistAuth", "UserID", in.GetUserID(), "error", err)
			return &npool.ExistAuthResponse{}, status.Error(codes.InvalidArgument, "UserID is invalid")
		}
	}

	if in.GetResource() == "" {
		logger.Sugar().Errorw("ExistAuth", "Resource", in.GetResource)
		return &npool.ExistAuthResponse{}, status.Error(codes.InvalidArgument, "Resource is invalid")
	}

	if in.GetMethod() == "" {
		logger.Sugar().Errorw("ExistAuth", "Method", in.GetMethod)
		return &npool.ExistAuthResponse{}, status.Error(codes.InvalidArgument, "Method is invalid")
	}

	span = tracer.Trace(span, in)

	span = commontracer.TraceInvoker(span, "auth", "auth", "ExistAuth")

	exist, err := authing1.ExistAuth(ctx, in.GetAppID(), in.UserID, in.GetResource(), in.GetMethod())
	if err != nil {
		logger.Sugar().Errorw("ExistAuth", "error", err)
		return &npool.ExistAuthResponse{}, status.Error(codes.InvalidArgument, "fail check auth")
	}

	return &npool.ExistAuthResponse{
		Info: exist,
	}, nil
}

func (s *Server) GetAuth(ctx context.Context, in *npool.GetAuthRequest) (resp *npool.GetAuthResponse, err error) {
	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetRoles")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.ID)

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetAuth", "ID", in.GetID(), "error", err)
		return &npool.GetAuthResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	span = commontracer.TraceInvoker(span, "auth", "auth", "GetAuths")

	info, err := authing1.GetAuth(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetAuths", "error", err)
		return &npool.GetAuthResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAuthResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAuths(ctx context.Context, in *npool.GetAuthsRequest) (info *npool.GetAuthsResponse, err error) {
	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetRoles")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.String("AppID", in.GetAppID()))
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetAuths", "AppID", in.GetAppID(), "error", err)
		return &npool.GetAuthsResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	span = commontracer.TraceInvoker(span, "auth", "auth", "GetAuths")

	infos, total, err := authing1.GetAuths(ctx, in.GetAppID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetAuths", "error", err)
		return &npool.GetAuthsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAuthsResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}

func (s *Server) GetHistories(ctx context.Context, in *npool.GetHistoriesRequest) (info *npool.GetHistoriesResponse, err error) {
	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetHistories")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.String("AppID", in.GetAppID()))
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetHistories", "AppID", in.GetAppID(), "error", err)
		return &npool.GetHistoriesResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	span = commontracer.TraceInvoker(span, "auth", "auth", "GetHistories")

	infos, total, err := authing1.GetHistories(ctx, in.GetAppID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetHistories", "error", err)
		return &npool.GetHistoriesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetHistoriesResponse{
		Infos: infos,
		Total: uint32(total),
	}, nil
}
