package user

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	cuser "github.com/NpoolPlatform/appuser-middleware/pkg/converter/v1/user"
	muser "github.com/NpoolPlatform/appuser-middleware/pkg/user"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) GetUser(ctx context.Context, in *npool.GetUserRequest) (*npool.GetUserResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetUser", "error", err)
		return &npool.GetUserResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("GetUser", "error", err)
		return &npool.GetUserResponse{}, status.Error(codes.InvalidArgument, "UserID is invalid")
	}

	info, err := muser.GetUser(ctx, in.GetAppID(), in.GetUserID())
	if err != nil {
		logger.Sugar().Errorw("GetUser", "error", err)
		return &npool.GetUserResponse{}, status.Error(codes.Internal, "fail get user")
	}

	return &npool.GetUserResponse{
		Info: cuser.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetUsers(ctx context.Context, in *npool.GetUsersRequest) (*npool.GetUsersResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("GetUsers", "error", err)
		return &npool.GetUsersResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	infos, err := muser.GetUsers(ctx, in.GetAppID(), in.GetOffset(), in.GetLimit())
	if err != nil {
		logger.Sugar().Errorw("GetUsers", "error", err)
		return &npool.GetUsersResponse{}, status.Error(codes.Internal, "fail get users")
	}

	return &npool.GetUsersResponse{
		Infos: cuser.Ent2GrpcMany(infos),
	}, nil
}

func (s *Server) GetManyUsers(ctx context.Context, in *npool.GetManyUsersRequest) (*npool.GetManyUsersResponse, error) {
	if len(in.IDs) == 0 {
		logger.Sugar().Errorw("GetManyUsers", "ids empty")
		return &npool.GetManyUsersResponse{}, status.Error(codes.InvalidArgument, "ids empty")
	}

	// TODO: parse id

	infos, err := muser.GetManyUsers(ctx, in.GetIDs())
	if err != nil {
		logger.Sugar().Errorw("GetManyUsers", "error", err)
		return &npool.GetManyUsersResponse{}, status.Error(codes.Internal, "fail get many users")
	}

	return &npool.GetManyUsersResponse{
		Infos: cuser.Ent2GrpcMany(infos),
	}, nil
}
