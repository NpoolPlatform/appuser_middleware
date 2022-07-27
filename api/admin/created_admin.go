package admin

import (
	"context"

	bconst "github.com/NpoolPlatform/appuser-middleware/pkg/const"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/google/uuid"

	mw "github.com/NpoolPlatform/appuser-middleware/pkg/middleware/admin"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/appusermw/admin"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) CreateGenesisRoleUser(ctx context.Context, in *admin.CreateGenesisRoleUserRequest) (*admin.CreateGenesisRoleUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateGenesisRoleUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if _, err := uuid.Parse(in.GetUser().GetAppID()); err != nil {
		logger.Sugar().Error("AppID is invalid")
		return &admin.CreateGenesisRoleUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.GetUser().GetAppID() != bconst.GenesisAppID && in.GetUser().GetAppID() != bconst.ChurchAppID {
		return nil, status.Error(codes.InvalidArgument, "invalid app id for genesis role user")
	}

	span.AddEvent("call middleware CreateGenesisRoleUser")
	user, roleUser, err := mw.CreateGenesisRoleUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app info: %v", err)
		return &admin.CreateGenesisRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &admin.CreateGenesisRoleUserResponse{
		User:     user,
		RoleUser: roleUser,
	}, nil
}
