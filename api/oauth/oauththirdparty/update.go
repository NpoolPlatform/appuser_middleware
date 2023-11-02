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

func (s *Server) UpdateOAuthThirdParty(ctx context.Context, in *npool.UpdateOAuthThirdPartyRequest) (*npool.UpdateOAuthThirdPartyResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateOAuthThirdParty",
			"In", in,
		)
		return &npool.UpdateOAuthThirdPartyResponse{}, status.Error(codes.Aborted, "Info is empty")
	}
	handler, err := oauththirdparty1.NewHandler(
		ctx,
		oauththirdparty1.WithID(req.ID, true),
		oauththirdparty1.WithClientName(req.ClientName, false),
		oauththirdparty1.WithClientTag(req.ClientTag, false),
		oauththirdparty1.WithClientLogoURL(req.ClientLogoURL, false),
		oauththirdparty1.WithClientOAuthURL(req.ClientOAuthURL, false),
		oauththirdparty1.WithResponseType(req.ResponseType, false),
		oauththirdparty1.WithScope(req.Scope, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOAuthThirdParty",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOAuthThirdPartyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateOAuthThirdParty(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOAuthThirdParty",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOAuthThirdPartyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateOAuthThirdPartyResponse{
		Info: info,
	}, nil
}
