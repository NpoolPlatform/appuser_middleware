package appoauththirdparty

import (
	"context"

	oauththirdparty1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/oauth/appoauththirdparty"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/oauth/appoauththirdparty"

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
		oauththirdparty1.WithAppID(req.AppID, true),
		oauththirdparty1.WithThirdPartyID(req.ThirdPartyID, true),
		oauththirdparty1.WithClientID(req.ClientID, true),
		oauththirdparty1.WithClientSecret(req.ClientSecret, true),
		oauththirdparty1.WithCallbackURL(req.CallbackURL, true),
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
