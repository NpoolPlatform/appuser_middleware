//nolint:dupl
package oauththirdparty

import (
	"context"

	oauththirdparty1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/oauth/oauththirdparty"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/oauth/oauththirdparty"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateOAuthThirdParty(ctx context.Context, in *npool.CreateOAuthThirdPartyRequest) (*npool.CreateOAuthThirdPartyResponse, error) {
	req := in.GetInfo()
	handler, err := oauththirdparty1.NewHandler(
		ctx,
		oauththirdparty1.WithID(req.ID),
		oauththirdparty1.WithClientName(req.ClientName),
		oauththirdparty1.WithClientTag(req.ClientTag),
		oauththirdparty1.WithClientLogoURL(req.ClientLogoURL),
		oauththirdparty1.WithClientOAuthURL(req.ClientOAuthURL),
		oauththirdparty1.WithResponseType(req.ResponseType),
		oauththirdparty1.WithScope(req.Scope),
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
