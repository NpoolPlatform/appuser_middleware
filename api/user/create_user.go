package user

import (
	"context"

	grpc "github.com/NpoolPlatform/appuser-manager/pkg/client"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	mw "github.com/NpoolPlatform/appuser-middleware/pkg/middleware/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/appusermgrv2/appuser"
	"github.com/NpoolPlatform/message/npool/appusermw/user"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func appUserRowToObject(row *ent.AppUser) *appuser.AppUser {
	return &appuser.AppUser{
		ID:            row.ID.String(),
		AppID:         row.AppID.String(),
		EmailAddress:  row.EmailAddress,
		PhoneNo:       row.PhoneNo,
		ImportFromApp: row.ImportFromApp.String(),
	}
}

func (s *Service) CreateUserWithSecret(ctx context.Context, in *user.CreateUserWithSecretRequest) (*user.CreateUserWithSecretResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateUserWithSecret")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	err = validateUserSecret(in)
	if err != nil {
		logger.Sugar().Errorw("fail validate user secret : %v", err)
		return nil, err
	}

	if in.GetUser().GetID() != "" {
		span.AddEvent("call grpc ExistAppUserV2")
		exist, err := grpc.ExistAppUserV2(ctx, in.GetUser().GetID())
		if err != nil {
			logger.Sugar().Errorw("fail check ban app: %v", err)
			return &user.CreateUserWithSecretResponse{}, status.Error(codes.Internal, err.Error())
		}

		if exist {
			return &user.CreateUserWithSecretResponse{}, status.Error(codes.AlreadyExists, "user already exists")
		}
	}

	if _, err := uuid.Parse(in.GetUser().GetAppID()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call middleware GetUserInfos")
	info, err := mw.CreateUserWithSecret(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app infos: %v", err)
		return &user.CreateUserWithSecretResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &user.CreateUserWithSecretResponse{
		Info: appUserRowToObject(info),
	}, nil
}

func (s *Service) CreateUserWithThirdParty(ctx context.Context, in *user.CreateUserWithThirdPartyRequest) (*user.CreateUserWithThirdPartyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetUserInfo")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	err = validateUserThirdParty(in)
	if err != nil {
		logger.Sugar().Errorw("fail validate user third party : %v", err)
		return nil, err
	}

	if in.GetUser().GetID() != "" {
		span.AddEvent("call grpc ExistAppUserV2")
		exist, err := grpc.ExistAppUserV2(ctx, in.GetUser().GetID())
		if err != nil {
			logger.Sugar().Errorw("fail check app user : %v", err)
			return &user.CreateUserWithThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
		}

		if exist {
			return &user.CreateUserWithThirdPartyResponse{}, status.Error(codes.AlreadyExists, "user already exists")
		}
	}

	if _, err := uuid.Parse(in.GetUser().GetAppID()); err != nil {
		logger.Sugar().Error("AppID is invalid")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	span.AddEvent("call middleware CreateUserWithThirdParty")
	info, err := mw.CreateUserWithThirdParty(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("fail get app infos: %v", err)
		return &user.CreateUserWithThirdPartyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &user.CreateUserWithThirdPartyResponse{
		Info: appUserRowToObject(info),
	}, nil
}
