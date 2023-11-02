package role

import (
	"context"

	role1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/role"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateRole(ctx context.Context, in *npool.UpdateRoleRequest) (*npool.UpdateRoleResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateRole",
			"In", in,
		)
		return &npool.UpdateRoleResponse{}, status.Error(codes.Aborted, "Info is empty")
	}
	handler, err := role1.NewHandler(
		ctx,
		role1.WithID(req.ID, true),
		role1.WithAppID(req.AppID, false),
		role1.WithRole(req.Role, false),
		role1.WithDescription(req.Description, false),
		role1.WithDefault(req.Default, false),
		role1.WithGenesis(req.Genesis, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRole",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRoleResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateRole(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateRole",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateRoleResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateRoleResponse{
		Info: info,
	}, nil
}
