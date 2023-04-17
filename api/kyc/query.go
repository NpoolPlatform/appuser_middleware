package kyc

import (
	"context"

	tracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer/kyc"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	servicename "github.com/NpoolPlatform/appuser-middleware/pkg/servicename"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	ckyc "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/kyc"
	mkyc "github.com/NpoolPlatform/appuser-middleware/pkg/mw/kyc"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) GetKyc(ctx context.Context, in *npool.GetKycRequest) (*npool.GetKycResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetKyc")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetKyc", "error", err)
		return &npool.GetKycResponse{}, status.Error(codes.InvalidArgument, "ID is invalid")
	}

	span = commontracer.TraceInvoker(span, "kyc", "middleware", "GetKyc")

	info, err := mkyc.GetKyc(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetKyc", "error", err)
		return &npool.GetKycResponse{}, status.Error(codes.Internal, "fail get kyc")
	}

	return &npool.GetKycResponse{
		Info: ckyc.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetKycs(ctx context.Context, in *npool.GetKycsRequest) (*npool.GetKycsResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "GetKycs")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	tracer.TraceConds(span, in.GetConds().GetConds())

	span = commontracer.TraceInvoker(span, "kyc", "middleware", "GetKycs")

	infos, total, err := mkyc.GetKycs(ctx, in.GetConds(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetKycs", "error", err)
		return &npool.GetKycsResponse{}, status.Error(codes.Internal, "fail get kycs")
	}

	return &npool.GetKycsResponse{
		Infos: ckyc.Ent2GrpcMany(infos),
		Total: total,
	}, nil
}
