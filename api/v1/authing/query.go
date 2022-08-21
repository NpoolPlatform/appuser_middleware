package authing

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing"

	authing1 "github.com/NpoolPlatform/appuser-middleware/pkg/authing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

func (s *Server) ExistAuth(ctx context.Context, in *npool.ExistAuthRequest) (*npool.ExistAuthResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("ExistAuth", "AppID", in.GetAppID(), "error", err)
		return &npool.ExistAuthResponse{}, status.Error(codes.InvalidArgument, "AppID is invalid")
	}
	if in.UserID != nil {
		if _, err := uuid.Parse(in.GetUserID()); err != nil {
			logger.Sugar().Errorw("ExistAuth", "UserID", in.GetUserID(), "error", err)
			return &npool.ExistAuthResponse{}, status.Error(codes.InvalidArgument, "UserID is invalid")
		}
	}

	if in.GetResource() == "" {
		logger.Sugar().Errorw("ExistAuth", "Resource", in.GetResource)
		return &npool.ExistAuthResponse{}, status.Error(codes.InvalidArgument, "Resource is invalid")
	}

	if in.GetMethod() == "" {
		logger.Sugar().Errorw("ExistAuth", "Method", in.GetMethod)
		return &npool.ExistAuthResponse{}, status.Error(codes.InvalidArgument, "Method is invalid")
	}

	exist, err := authing1.ExistAuth(ctx, in.GetAppID(), in.UserID, in.GetResource(), in.GetMethod())
	if err != nil {
		logger.Sugar().Errorw("ExistAuth", "error", err)
		return &npool.ExistAuthResponse{}, status.Error(codes.InvalidArgument, "fail check auth")
	}

	return &npool.ExistAuthResponse{
		Info: exist,
	}, nil
}

func (s *Server) GetAuths(ctx context.Context, in *npool.GetAuthsRequest) (*npool.GetAuthsResponse, error) {
	return &npool.GetAuthsResponse{}, status.Error(codes.Unimplemented, "NOT IMPLEMENTED")
}

func (s *Server) GetHistories(ctx context.Context, in *npool.GetHistoriesRequest) (*npool.GetHistoriesResponse, error) {
	return &npool.GetHistoriesResponse{}, status.Error(codes.Unimplemented, "NOT IMPLEMENTED")
}
