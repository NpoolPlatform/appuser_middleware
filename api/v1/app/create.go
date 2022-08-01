package app

import (
	"context"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-middleware/pkg/tracer/app"
	"go.opentelemetry.io/otel"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	mw "github.com/NpoolPlatform/appuser-middleware/pkg/app"
	capp "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateApp(ctx context.Context, in *npool.CreateAppRequest) (*npool.CreateAppResponse, error) {
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

	if err := validate(in.GetInfo()); err != nil {
		logger.Sugar().Errorw("CreateApp", "error", err)
		return &npool.CreateAppResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "app", "middleware", "CreateApp")

	info, err := mw.CreateApp(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateApp", "error", err)
		return &npool.CreateAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	ginfo, err := capp.Ent2Grpc(info)
	if err != nil {
		logger.Sugar().Errorw("CreateApp", "error", err)
		return &npool.CreateAppResponse{}, status.Error(codes.Internal, "invalid value")
	}

	return &npool.CreateAppResponse{
		Info: ginfo,
	}, nil
}
