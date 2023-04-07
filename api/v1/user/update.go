package user

import (
	"context"

	appusercrud "github.com/NpoolPlatform/appuser-manager/pkg/crud/appuser"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npoolpb "github.com/NpoolPlatform/message/npool"
	appuserpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"

	cuser "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/user"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	servicename "github.com/NpoolPlatform/appuser-middleware/pkg/servicename"
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

//nolint:dupl
func (s *Server) UpdateUser(ctx context.Context, in *npool.UpdateUserRequest) (*npool.UpdateUserResponse, error) {
	var err error

	_, span := otel.Tracer(servicename.ServiceDomain).Start(ctx, "UpdateUser")
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

	if in.Info.PhoneNO != nil && in.Info.GetPhoneNO() != "" {
		phoneExist, err := appusercrud.ExistConds(ctx, &appuserpb.Conds{
			PhoneNO: &npoolpb.StringVal{
				Op:    cruder.EQ,
				Value: in.Info.GetPhoneNO(),
			},
			AppID: &npoolpb.StringVal{
				Op:    cruder.EQ,
				Value: in.Info.GetAppID(),
			},
		})

		if err != nil {
			logger.Sugar().Errorw("validate", "err", err)
			return &npool.UpdateUserResponse{}, status.Error(codes.Internal, err.Error())
		}

		if phoneExist {
			logger.Sugar().Errorw("validate", "phoneExsit", phoneExist)
			return &npool.UpdateUserResponse{}, status.Error(codes.AlreadyExists, "phone already exists")
		}
	}

	if in.Info.EmailAddress != nil && in.Info.GetEmailAddress() != "" {
		emailExist, err := appusercrud.ExistConds(ctx, &appuserpb.Conds{
			EmailAddress: &npoolpb.StringVal{
				Op:    cruder.EQ,
				Value: in.Info.GetEmailAddress(),
			},
			AppID: &npoolpb.StringVal{
				Op:    cruder.EQ,
				Value: in.Info.GetAppID(),
			},
		})
		if err != nil {
			logger.Sugar().Errorw("validate", "err", err)
			return &npool.UpdateUserResponse{}, status.Error(codes.Internal, err.Error())
		}

		if emailExist {
			logger.Sugar().Errorw("validate", "emailExist", emailExist)
			return &npool.UpdateUserResponse{}, status.Error(codes.AlreadyExists, "email already exists")
		}
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
