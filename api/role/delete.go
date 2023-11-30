package role

import (
	"context"

	role1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/role"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteRole(ctx context.Context, in *npool.DeleteRoleRequest) (*npool.DeleteRoleResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteRole",
			"In", in,
		)
		return &npool.DeleteRoleResponse{}, status.Error(codes.Aborted, "Info is empty")
	}
	handler, err := role1.NewHandler(
		ctx,
		role1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRole",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRoleResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteRole(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteRole",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteRoleResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteRoleResponse{
		Info: info,
	}, nil
}
