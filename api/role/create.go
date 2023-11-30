package role

import (
	"context"

	role1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/role"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateRole(ctx context.Context, in *npool.CreateRoleRequest) (*npool.CreateRoleResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateRole",
			"In", in,
		)
		return &npool.CreateRoleResponse{}, status.Error(codes.Aborted, "Info is empty")
	}
	handler, err := role1.NewHandler(
		ctx,
		role1.WithEntID(req.EntID, false),
		role1.WithAppID(req.AppID, true),
		role1.WithCreatedBy(req.CreatedBy, true),
		role1.WithRole(req.Role, true),
		role1.WithDescription(req.Description, true),
		role1.WithDefault(req.Default, true),
		role1.WithGenesis(req.Genesis, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRole",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRoleResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.CreateRole(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRole",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRoleResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateRoleResponse{
		Info: info,
	}, nil
}

func (s *Server) CreateRoles(ctx context.Context, in *npool.CreateRolesRequest) (*npool.CreateRolesResponse, error) {
	handler, err := role1.NewHandler(
		ctx,
		role1.WithReqs(in.GetInfos(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRoles",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRolesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.CreateRoles(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateRoles",
			"In", in,
			"Error", err,
		)
		return &npool.CreateRolesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateRolesResponse{
		Infos: infos,
	}, nil
}
