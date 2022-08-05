package admin

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-middleware/pkg/tracer/admin"
	"go.opentelemetry.io/otel"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	mw "github.com/NpoolPlatform/appuser-middleware/pkg/admin"
	cuser "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/admin"
	scodes "go.opentelemetry.io/otel/codes"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateGenesisUser(ctx context.Context, in *npool.CreateGenesisUserRequest) (*npool.CreateGenesisUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateGenesisUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	span = commontracer.TraceInvoker(span, "app", "middleware", "CreateGenesisUser")

	if err := validate(in); err != nil {
		logger.Sugar().Errorw("CreateGenesisUser", "error", err)
		return &npool.CreateGenesisUserResponse{}, err
	}

	info, err := mw.CreateGenesisUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("CreateGenesisUser", "error", err)
		return &npool.CreateGenesisUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateGenesisUserResponse{
		Info: cuser.Ent2Grpc(info),
	}, nil
}

func (s *Server) AuthorizeGenesis(ctx context.Context, in *npool.AuthorizeGenesisRequest) (*npool.AuthorizeGenesisResponse, error) {
	// TODO: Wait for authing-gateway refactoring to complete the API
	return &npool.AuthorizeGenesisResponse{}, status.Error(codes.Internal, fmt.Errorf("NOT IMPLEMENTED").Error())
}
