package role

import (
	"context"

	role1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/role"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetRole(ctx context.Context, in *npool.GetRoleRequest) (*npool.GetRoleResponse, error) {
	handler, err := role1.NewHandler(
		ctx,
		role1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRole",
			"In", in,
			"Error", err,
		)
		return &npool.GetRoleResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.GetRole(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRole",
			"In", in,
			"Error", err,
		)
		return &npool.GetRoleResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetRoleResponse{
		Info: info,
	}, nil
}

func (s *Server) GetRoles(ctx context.Context, in *npool.GetRolesRequest) (*npool.GetRolesResponse, error) {
	handler, err := role1.NewHandler(
		ctx,
		role1.WithConds(in.GetConds()),
		role1.WithOffset(in.GetOffset()),
		role1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRoles",
			"In", in,
			"Error", err,
		)
		return &npool.GetRolesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := handler.GetRoles(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetRoles",
			"In", in,
			"Error", err,
		)
		return &npool.GetRolesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetRolesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
