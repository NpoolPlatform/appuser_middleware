package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"

	cuser "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/user"
	muser "github.com/NpoolPlatform/appuser-middleware/pkg/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	signmethod "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

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
	if in.GetAccount() == "" {
		logger.Sugar().Errorw("VerifyUser", "error", err)
		return &npool.VerifyUserResponse{}, status.Error(codes.InvalidArgument, "Account is invalid")
	}

	switch in.GetAccountType() {
	case signmethod.SignMethodType_Email:
	case signmethod.SignMethodType_Mobile:
	default:
		logger.Sugar().Errorw("VerifyUser", "AccountType", in.GetAccountType())
		return &npool.VerifyUserResponse{}, status.Error(codes.InvalidArgument, "AccountType is invalid")
	}

	span = commontracer.TraceInvoker(span, "user", "middleware", "VerifyUser")

	info, err := muser.VerifyUser(ctx, in.GetAppID(), in.GetAccount(), in.GetAccountType(), in.GetPasswordHash())
	if err != nil {
		logger.Sugar().Errorw("VerifyUser", "error", err)
		return &npool.VerifyUserResponse{}, status.Error(codes.Internal, "fail verify user")
	}

	return &npool.VerifyUserResponse{
		Info: cuser.Ent2Grpc(info),
	}, nil
}
