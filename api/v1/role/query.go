//nolint:dupl
package role

import (
	"context"

	croleuser "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/roleuser"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"

	crole "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/role"
	mrole "github.com/NpoolPlatform/appuser-middleware/pkg/role"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) GetRoles(ctx context.Context, in *npool.GetRolesRequest) (*npool.GetRolesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetRoles")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	span.SetAttributes(attribute.String("AppID", in.GetAppID()))

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetRoles", "error", err)
		return &npool.GetRolesResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	span = commontracer.TraceInvoker(span, "role", "middleware", "GetRoles")

	infos, total, err := mrole.GetRoles(ctx, in.GetAppID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetRoles", "error", err)
		return &npool.GetRolesResponse{}, status.Error(codes.Internal, "fail get roles")
	}

	return &npool.GetRolesResponse{
		Infos: crole.Ent2GrpcMany(infos),
		Total: uint32(total),
	}, nil
}

func (s *Server) GetAppRoles(ctx context.Context, in *npool.GetAppRolesRequest) (*npool.GetAppRolesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	span.SetAttributes(attribute.String("TargetAppID", in.GetTargetAppID()))

	if _, err := uuid.Parse(in.GetTargetAppID()); err != nil {
		logger.Sugar().Errorw("GetUser", "error", err)
		return &npool.GetAppRolesResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	span = commontracer.TraceInvoker(span, "user", "middleware", "GetUser")

	infos, total, err := mrole.GetRoles(ctx, in.GetTargetAppID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetUser", "error", err)
		return &npool.GetAppRolesResponse{}, status.Error(codes.Internal, "fail get user")
	}

	return &npool.GetAppRolesResponse{
		Infos: crole.Ent2GrpcMany(infos),
		Total: uint32(total),
	}, nil
}

func (s *Server) GetManyRoles(ctx context.Context, in *npool.GetManyRolesRequest) (*npool.GetManyRolesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetManyRoles")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.StringSlice("UserIDs", in.GetIDs()))

	if len(in.IDs) == 0 {
		logger.Sugar().Errorw("GetManyRoles", "ids empty")
		return &npool.GetManyRolesResponse{}, status.Error(codes.InvalidArgument, "ids empty")
	}

	for _, val := range in.GetIDs() {
		if _, err := uuid.Parse(val); err != nil {
			logger.Sugar().Errorw("GetManyRoles", "error", err)
			return &npool.GetManyRolesResponse{}, status.Error(codes.InvalidArgument, "IDs is invalid")
		}
	}

	span = commontracer.TraceInvoker(span, "user", "middleware", "GetManyRoles")

	infos, err := mrole.GetManyRoles(ctx, in.GetIDs())
	if err != nil {
		logger.Sugar().Errorw("GetManyRoles", "error", err)
		return &npool.GetManyRolesResponse{}, status.Error(codes.Internal, "fail get many users")
	}

	return &npool.GetManyRolesResponse{
		Infos: crole.Ent2GrpcMany(infos),
		Total: uint32(len(infos)),
	}, nil
}

func (s *Server) GetManyRoleUsers(ctx context.Context, in *npool.GetManyRoleUsersRequest) (*npool.GetManyRoleUsersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetManyRoleUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.StringSlice("IDs", in.GetIDs()))

	if len(in.GetIDs()) == 0 {
		logger.Sugar().Errorw("GetManyRoleUsers", "ids empty")
		return &npool.GetManyRoleUsersResponse{}, status.Error(codes.InvalidArgument, "ids empty")
	}

	for _, val := range in.GetIDs() {
		if _, err := uuid.Parse(val); err != nil {
			logger.Sugar().Errorw("GetManyRoleUsers", "error", err)
			return &npool.GetManyRoleUsersResponse{}, status.Error(codes.InvalidArgument, "IDs is invalid")
		}
	}

	span = commontracer.TraceInvoker(span, "user", "middleware", "GetManyRoles")

	infos, err := mrole.GetManyRoleUsers(ctx, in.GetIDs())
	if err != nil {
		logger.Sugar().Errorw("GetManyRoles", "error", err)
		return &npool.GetManyRoleUsersResponse{}, status.Error(codes.Internal, "fail get many users")
	}

	return &npool.GetManyRoleUsersResponse{
		Infos: croleuser.Ent2GrpcMany(infos),
		Total: uint32(len(infos)),
	}, nil
}
