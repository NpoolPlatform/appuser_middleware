package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"

	cuser "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/user"
	muser "github.com/NpoolPlatform/appuser-middleware/pkg/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) VerifyAccount(ctx context.Context, in *npool.VerifyAccountRequest) (*npool.VerifyAccountResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "VerifyAccount")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.String("AppID", in.GetAppID()))

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("VerifyAccount", "error", err)
		return &npool.VerifyAccountResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}
	if in.GetAccount() == "" {
		logger.Sugar().Errorw("VerifyAccount", "error", err)
		return &npool.VerifyAccountResponse{}, status.Error(codes.InvalidArgument, "Account is invalid")
	}

	switch in.GetAccountType() {
	case basetypes.SignMethod_Email:
	case basetypes.SignMethod_Mobile:
	default:
		logger.Sugar().Errorw("VerifyAccount", "AccountType", in.GetAccountType())
		return &npool.VerifyAccountResponse{}, status.Error(codes.InvalidArgument, "AccountType is invalid")
	}

	if in.GetPasswordHash() == "" {
		logger.Sugar().Errorw("VerifyAccount", "error", err)
		return &npool.VerifyAccountResponse{}, status.Error(codes.InvalidArgument, "PasswordHash is invalid")
	}

	span = commontracer.TraceInvoker(span, "user", "middleware", "VerifyAccount")

	info, err := muser.VerifyAccount(ctx, in.GetAppID(), in.GetAccount(), in.GetAccountType(), in.GetPasswordHash())
	if err != nil {
		logger.Sugar().Errorw("VerifyAccount", "error", err)
		return &npool.VerifyAccountResponse{}, status.Error(codes.Internal, "fail verify user")
	}

	return &npool.VerifyAccountResponse{
		Info: cuser.Ent2Grpc(info),
	}, nil
}

func (s *Server) VerifyUser(ctx context.Context, in *npool.VerifyUserRequest) (*npool.VerifyUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "VerifyUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span.SetAttributes(attribute.String("AppID", in.GetAppID()))

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("VerifyUser", "error", err)
		return &npool.VerifyUserResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("VerifyUser", "error", err)
		return &npool.VerifyUserResponse{}, status.Error(codes.InvalidArgument, "UserID is invalid")
	}
	if in.GetPasswordHash() == "" {
		logger.Sugar().Errorw("VerifyUser", "error", err)
		return &npool.VerifyUserResponse{}, status.Error(codes.InvalidArgument, "PasswordHash is invalid")
	}

	span = commontracer.TraceInvoker(span, "user", "middleware", "VerifyUser")

	info, err := muser.VerifyUser(ctx, in.GetAppID(), in.GetUserID(), in.GetPasswordHash())
	if err != nil {
		logger.Sugar().Errorw("VerifyUser", "error", err)
		return &npool.VerifyUserResponse{}, status.Error(codes.Internal, "fail verify user")
	}

	return &npool.VerifyUserResponse{
		Info: cuser.Ent2Grpc(info),
	}, nil
}
