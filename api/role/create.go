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
	handler, err := role1.NewHandler(
		ctx,
		role1.WithID(req.ID),
		role1.WithAppID(req.GetAppID()),
		role1.WithCreatedBy(req.CreatedBy),
		role1.WithRole(req.Role),
		role1.WithDescription(req.Description),
		role1.WithDefault(req.Default),
		role1.WithGenesis(req.Genesis),
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
