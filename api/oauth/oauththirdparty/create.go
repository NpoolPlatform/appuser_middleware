//nolint:dupl
package oauththirdparty

import (
	"context"

	oauththirdparty1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/oauth/oauththirdparty"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/oauth/oauththirdparty"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateOAuthThirdParty(ctx context.Context, in *npool.CreateOAuthThirdPartyRequest) (*npool.CreateOAuthThirdPartyResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateOAuthThirdParty",
			"In", in,
		)
		return &npool.CreateOAuthThirdPartyResponse{}, status.Error(codes.Aborted, "Info is empty")
	}
	handler, err := oauththirdparty1.NewHandler(
		ctx,
		oauththirdparty1.WithEntID(req.EntID, false),
		oauththirdparty1.WithClientName(req.ClientName, true),
		oauththirdparty1.WithClientTag(req.ClientTag, true),
		oauththirdparty1.WithClientLogoURL(req.ClientLogoURL, true),
		oauththirdparty1.WithClientOAuthURL(req.ClientOAuthURL, true),
		oauththirdparty1.WithResponseType(req.ResponseType, true),
		oauththirdparty1.WithScope(req.Scope, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOAuthThirdParty",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOAuthThirdPartyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.CreateOAuthThirdParty(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOAuthThirdParty",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOAuthThirdPartyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateOAuthThirdPartyResponse{
		Info: info,
	}, nil
}
