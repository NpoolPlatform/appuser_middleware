package kyc

import (
	"context"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	scodes "go.opentelemetry.io/otel/codes"

	ckyc "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/kyc"
	mkyc "github.com/NpoolPlatform/appuser-middleware/pkg/kyc"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) GetKyc(ctx context.Context, in *npool.GetKycRequest) (*npool.GetKycResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetKyc")
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

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetKycs")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	span.SetAttributes(attribute.String("AppID", in.GetAppID()))

	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetKycs", "error", err)
		return &npool.GetKycsResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	span = commontracer.TraceInvoker(span, "kyc", "middleware", "GetKycs")

	infos, total, err := mkyc.GetKycs(ctx, in.GetAppID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetKycs", "error", err)
		return &npool.GetKycsResponse{}, status.Error(codes.Internal, "fail get kycs")
	}

	return &npool.GetKycsResponse{
		Infos: ckyc.Ent2GrpcMany(infos),
		Total: total,
	}, nil
}
