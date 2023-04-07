package user

import (
	"context"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	servicename "github.com/NpoolPlatform/appuser-middleware/pkg/servicename"
	mw "github.com/NpoolPlatform/appuser-middleware/pkg/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteUser(ctx context.Context, in *npool.DeleteUserRequest) (*npool.DeleteUserResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "DeleteUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	span = commontracer.TraceInvoker(span, "user", "middleware", "DeleteUser")

	userID, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("DeleteUser", "error", err)
		return &npool.DeleteUserResponse{}, status.Error(codes.Internal, "UserID is invalid")
	}

	err = mw.DeleteUser(ctx, userID)
	if err != nil {
		logger.Sugar().Errorw("DeleteUser", "error", err)
		return &npool.DeleteUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteUserResponse{}, nil
}
