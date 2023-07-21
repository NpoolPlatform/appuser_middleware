//nolint:dupl
package appoauththirdparty

import (
	"context"

	oauththirdparty1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/oauth/appoauththirdparty"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/oauth/appoauththirdparty"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetOAuthThirdParty(ctx context.Context, in *npool.GetOAuthThirdPartyRequest) (*npool.GetOAuthThirdPartyResponse, error) {
	_handler, err := oauththirdparty1.NewHandler(
		ctx,
		oauththirdparty1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOAuthThirdParty",
			"In", in,
			"Error", err,
		)
		return &npool.GetOAuthThirdPartyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := _handler.GetOAuthThirdParty(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOAuthThirdParty",
			"In", in,
			"Error", err,
		)
		return &npool.GetOAuthThirdPartyResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetOAuthThirdPartyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetOAuthThirdParties(ctx context.Context, in *npool.GetOAuthThirdPartiesRequest) (*npool.GetOAuthThirdPartiesResponse, error) {
	_handler, err := oauththirdparty1.NewHandler(
		ctx,
		oauththirdparty1.WithConds(in.GetConds()),
		oauththirdparty1.WithOffset(in.GetOffset()),
		oauththirdparty1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOAuthThirdParties",
			"In", in,
			"Error", err,
		)
		return &npool.GetOAuthThirdPartiesResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, total, err := _handler.GetOAuthThirdParties(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOAuthThirdParties",
			"In", in,
			"Error", err,
		)
		return &npool.GetOAuthThirdPartiesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.GetOAuthThirdPartiesResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) GetOAuthThirdPartyDecryptOnly(ctx context.Context, in *npool.GetOAuthThirdPartyDecryptOnlyRequest) (*npool.GetOAuthThirdPartyDecryptOnlyResponse, error) {
	const limit = 2
	_handler, err := oauththirdparty1.NewHandler(
		ctx,
		oauththirdparty1.WithConds(in.GetConds()),
		oauththirdparty1.WithOffset(0),
		oauththirdparty1.WithLimit(limit),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOAuthThirdParties",
			"In", in,
			"Error", err,
		)
		return &npool.GetOAuthThirdPartyDecryptOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, _, err := _handler.GetOAuthThirdPartiesDecrypt(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOAuthThirdParties",
			"In", in,
			"Error", err,
		)
		return &npool.GetOAuthThirdPartyDecryptOnlyResponse{}, status.Error(codes.Aborted, err.Error())
	}
	if len(infos) > 1 {
		logger.Sugar().Errorw(
			"GetOAuthThirdParties",
			"In", in,
			"Error", err,
		)
		return &npool.GetOAuthThirdPartyDecryptOnlyResponse{}, status.Error(codes.Aborted, "too many records")
	}

	return &npool.GetOAuthThirdPartyDecryptOnlyResponse{
		Info: infos[0],
	}, nil
}
