package user

import (
	"context"

	commontracer "github.com/NpoolPlatform/appuser-manager/pkg/tracer"
	cuser "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/user"
	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	tracer "github.com/NpoolPlatform/appuser-middleware/pkg/tracer/user"
	mw "github.com/NpoolPlatform/appuser-middleware/pkg/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"

	subscriber "github.com/NpoolPlatform/appuser-middleware/pkg/subscriber"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	commonpb "github.com/NpoolPlatform/message/npool"
	submgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/subscriber"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(ctx context.Context, in *npool.CreateUserRequest) (*npool.CreateUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if err := validate(ctx, in.GetInfo()); err != nil {
		logger.Sugar().Errorw("CreateUser", "error", err)
		return &npool.CreateUserResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "user", "middleware", "CreateUser")

	info, err := mw.CreateUser(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateUser", "error", err)
		return &npool.CreateUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	sub, err := subscriber.GetSubscriberOnly(ctx, &submgrpb.Conds{
		AppID: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetInfo().GetAppID(),
		},
		EmailAddress: &commonpb.StringVal{
			Op:    cruder.EQ,
			Value: in.GetInfo().GetEmailAddress(),
		},
	})
	if err != nil {
		logger.Sugar().Errorw("CreateUser", "error", err)
		return &npool.CreateUserResponse{
			Info: cuser.Ent2Grpc(info),
		}, nil
	}
	if sub != nil && !sub.Registered {
		registered := true

		_, err = subscriber.UpdateSubscriber(ctx, &submgrpb.SubscriberReq{
			ID:         &sub.ID,
			Registered: &registered,
		})
		if err != nil {
			logger.Sugar().Errorw("CreateUser", "error", err)
			return &npool.CreateUserResponse{}, status.Error(codes.Internal, err.Error())
		}
	}

	return &npool.CreateUserResponse{
		Info: cuser.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateUserRevert(ctx context.Context, in *npool.CreateUserRequest) (*npool.CreateUserResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateUser")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if err := validate(ctx, in.GetInfo()); err != nil {
		logger.Sugar().Errorw("CreateUserRevert", "error", err)
		return &npool.CreateUserResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "user", "middleware", "DeleteUser")

	err = mw.DeleteUser(ctx, uuid.MustParse(in.GetInfo().GetID()))
	if err != nil {
		logger.Sugar().Errorw("CreateUserRevert", "error", err)
		return &npool.CreateUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateUserResponse{}, nil
}
