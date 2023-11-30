package role

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	role1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/role"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ExistRole(ctx context.Context, in *npool.ExistRoleRequest) (*npool.ExistRoleResponse, error) {
	handler, err := role1.NewHandler(
		ctx,
		role1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistRole",
			"In", in,
			"error", err,
		)
		return &npool.ExistRoleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.ExistRole(ctx)
	if err != nil {
		return &npool.ExistRoleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistRoleResponse{
		Info: info,
	}, nil
}

func (s *Server) ExistRoleConds(ctx context.Context, in *npool.ExistRoleCondsRequest) (*npool.ExistRoleCondsResponse, error) {
	handler, err := role1.NewHandler(
		ctx,
		role1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistRoleConds",
			"In", in,
			"error", err,
		)
		return &npool.ExistRoleCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	exist, err := handler.ExistRoleConds(ctx)
	if err != nil {
		return &npool.ExistRoleCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistRoleCondsResponse{
		Info: exist,
	}, nil
}
