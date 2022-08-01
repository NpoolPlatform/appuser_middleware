package user

import (
	"context"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/google/uuid"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	capp "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/user"
	tracer "github.com/NpoolPlatform/appuser-middleware/pkg/tracer/user"
	mw "github.com/NpoolPlatform/appuser-middleware/pkg/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(ctx context.Context, in *npool.CreateUserRequest) (*npool.CreateUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateApp")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if in.Info.ID == nil {
		id := uuid.NewString()
		in.Info.ID = &id
	}

	if err := validate(in.GetInfo()); err != nil {
		logger.Sugar().Errorw("CreateApp", "error", err)
		return &npool.CreateUserResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "user", "middleware", "CreateApp")

	info, err := mw.CreateUser(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateApp", "error", err)
		return &npool.CreateUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateUserResponse{
		Info: capp.Ent2Grpc(info),
	}, nil
}
