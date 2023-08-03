//nolint:nolintlint,dupl
package appoauththirdparty

import (
	"context"

	oauththirdparty1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/oauth/appoauththirdparty"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/oauth/appoauththirdparty"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateOAuthThirdParty(ctx context.Context, in *npool.UpdateOAuthThirdPartyRequest) (*npool.UpdateOAuthThirdPartyResponse, error) {
	req := in.GetInfo()
	handler, err := oauththirdparty1.NewHandler(
		ctx,
		oauththirdparty1.WithID(req.ID),
		oauththirdparty1.WithClientID(req.ClientID),
		oauththirdparty1.WithClientSecret(req.ClientSecret),
		oauththirdparty1.WithCallbackURL(req.CallbackURL),
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
