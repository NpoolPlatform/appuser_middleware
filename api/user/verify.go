package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	servicename "github.com/NpoolPlatform/appuser-middleware/pkg/servicename"

	cuser "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/user"
	user1 "github.com/NpoolPlatform/appuser-middleware/pkg/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) VerifyAccount(ctx context.Context, in *npool.VerifyAccountRequest) (*npool.VerifyAccountResponse, error) {
	handler, err := user1.NewHandler(
		ctx,
		user1.WithAppID(in.GetAppID()),
		user1.WithAccount(in.GetAccount(), in.GetAccountType()),
		user1.WithPasswordHash(&in.PasswordHash),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"VerifyAccount",
			"In", in,
			"error", err,
		)
		return &npool.VerifyAccountResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.VerifyAccount(ctx)
	if err != nil {
		return &npool.VerifyAccountResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.VerifyAccountResponse{
		Info: info,
	}, nil
}

func (s *Server) VerifyUser(ctx context.Context, in *npool.VerifyUserRequest) (*npool.VerifyUserResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "VerifyUser")
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

	info, err := user1.VerifyUser(ctx, in.GetAppID(), in.GetUserID(), in.GetPasswordHash())
	if err != nil {
		logger.Sugar().Errorw("VerifyUser", "error", err)
		return &npool.VerifyUserResponse{}, status.Error(codes.Internal, "fail verify user")
	}

	return &npool.VerifyUserResponse{
		Info: cuser.Ent2Grpc(info),
	}, nil
}
