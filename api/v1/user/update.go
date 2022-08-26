package user

import (
	"context"

	cuser "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/user"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-middleware/pkg/tracer/user"
	mw "github.com/NpoolPlatform/appuser-middleware/pkg/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) UpdateUser(ctx context.Context, in *npool.UpdateUserRequest) (*npool.UpdateUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	span = commontracer.TraceInvoker(span, "user", "middleware", "UpdateUser")

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Infow("UpdateUser", "ID", in.GetInfo().GetID())
		return &npool.UpdateUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetInfo().GetAppID()); err != nil {
		logger.Sugar().Infow("UpdateUser", "AppID", in.GetInfo().GetAppID())
		return &npool.UpdateUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := mw.UpdateUser(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateUser", "error", err)
		return &npool.UpdateUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateUserResponse{
		Info: cuser.Ent2Grpc(info),
	}, nil
}
