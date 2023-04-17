package app

import (
	"context"

	servicename "github.com/NpoolPlatform/appuser-middleware/pkg/servicename"
	tracer "github.com/NpoolPlatform/appuser-middleware/pkg/tracer/app"
	"go.opentelemetry.io/otel"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	capp "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/app"
	mw "github.com/NpoolPlatform/appuser-middleware/pkg/mw/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateApp(ctx context.Context, in *npool.CreateAppRequest) (*npool.CreateAppResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "CreateApp")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if err := validate(ctx, in.GetInfo()); err != nil {
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
