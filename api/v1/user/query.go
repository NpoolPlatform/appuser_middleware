package user

import (
	"context"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"

	cuser "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/user"
	muser "github.com/NpoolPlatform/appuser-middleware/pkg/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) GetUser(ctx context.Context, in *npool.GetUserRequest) (*npool.GetUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	commontracer.TraceID(span, in.GetUserID())

	span.SetAttributes(attribute.String("AppID", in.GetAppID()))

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetUser", "error", err)
		return &npool.GetUserResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("GetUser", "error", err)
		return &npool.GetUserResponse{}, status.Error(codes.InvalidArgument, "UserID is invalid")
	}

	span = commontracer.TraceInvoker(span, "user", "middleware", "GetUser")

	info, err := muser.GetUser(ctx, in.GetAppID(), in.GetUserID())
	if err != nil {
		logger.Sugar().Errorw("GetUser", "error", err)
		return &npool.GetUserResponse{}, status.Error(codes.Internal, "fail get user")
	}

	return &npool.GetUserResponse{
		Info: cuser.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetUsers(ctx context.Context, in *npool.GetUsersRequest) (*npool.GetUsersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.String("AppID", in.GetAppID()))

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetUsers", "error", err)
		return &npool.GetUsersResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	span = commontracer.TraceInvoker(span, "user", "middleware", "GetUsers")

	infos, total, err := muser.GetUsers(ctx, in.GetAppID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetUsers", "error", err)
		return &npool.GetUsersResponse{}, status.Error(codes.Internal, "fail get users")
	}

	return &npool.GetUsersResponse{
		Infos: cuser.Ent2GrpcMany(infos),
		Total: uint32(total),
	}, nil
}

func (s *Server) GetManyUsers(ctx context.Context, in *npool.GetManyUsersRequest) (*npool.GetManyUsersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetManyUsers")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.StringSlice("UserIDs", in.GetIDs()))

	if len(in.IDs) == 0 {
		logger.Sugar().Errorw("GetManyUsers", "ids empty")
		return &npool.GetManyUsersResponse{}, status.Error(codes.InvalidArgument, "ids empty")
	}

	for _, val := range in.GetIDs() {
		if _, err := uuid.Parse(val); err != nil {
			logger.Sugar().Errorw("GetManyUsers", "error", err)
			return &npool.GetManyUsersResponse{}, status.Error(codes.InvalidArgument, "IDs is invalid")
		}
	}

	span = commontracer.TraceInvoker(span, "user", "middleware", "GetManyUsers")

	infos, total, err := muser.GetManyUsers(ctx, in.GetIDs())
	if err != nil {
		logger.Sugar().Errorw("GetManyUsers", "error", err)
		return &npool.GetManyUsersResponse{}, status.Error(codes.Internal, "fail get many users")
	}

	return &npool.GetManyUsersResponse{
		Infos: cuser.Ent2GrpcMany(infos),
		Total: uint32(total),
	}, nil
}
