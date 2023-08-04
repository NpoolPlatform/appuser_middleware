package oauththirdparty

import (
	"context"

	oauththirdparty1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/oauth/oauththirdparty"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/oauth/oauththirdparty"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DeleteOAuthThirdParty(ctx context.Context, in *npool.DeleteOAuthThirdPartyRequest) (*npool.DeleteOAuthThirdPartyResponse, error) {
	req := in.GetInfo()
	handler, err := oauththirdparty1.NewHandler(
		ctx,
		oauththirdparty1.WithID(req.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteOAuthThirdParty",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOAuthThirdPartyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteOAuthThirdParty(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteOAuthThirdParty",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOAuthThirdPartyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteOAuthThirdPartyResponse{
		Info: info,
	}, nil
}
