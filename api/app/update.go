package app

import (
	"context"

	mw "github.com/NpoolPlatform/appuser-middleware/pkg/app"
	capp "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/app"
	servicename "github.com/NpoolPlatform/appuser-middleware/pkg/servicename"
	tracer "github.com/NpoolPlatform/appuser-middleware/pkg/tracer/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
)

func (s *Server) UpdateApp(ctx context.Context, in *npool.UpdateAppRequest) (*npool.UpdateAppResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "UpdateApp")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateApp", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAppResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := mw.UpdateApp(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateApp", "error", err)
		return &npool.UpdateAppResponse{}, status.Error(codes.Internal, err.Error())
	}

	ginfo, err := capp.Ent2Grpc(info)
	if err != nil {
		logger.Sugar().Errorw("UpdateApp", "error", err)
		return &npool.UpdateAppResponse{}, status.Error(codes.Internal, "invalid value")
	}

	return &npool.UpdateAppResponse{
		Info: ginfo,
	}, nil
}
