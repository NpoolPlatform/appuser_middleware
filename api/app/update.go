package app

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
)

func (s *Server) UpdateApp(ctx context.Context, in *npool.UpdateAppRequest) (*npool.UpdateAppResponse, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (s *Server) BanApp(ctx context.Context, in *npool.BanAppRequest) (*npool.BanAppResponse, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (s *Server) SetSignMethods(ctx context.Context, in *npool.SetSignMethodsRequest) (*npool.SetSignMethodsResponse, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (s *Server) SetRecaptcha(ctx context.Context, in *npool.SetRecaptchaRequest) (*npool.SetRecaptchaResponse, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (s *Server) SetKyc(ctx context.Context, in *npool.SetKycRequest) (*npool.SetKycResponse, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (s *Server) SetSigninVerify(ctx context.Context, in *npool.SetSigninVerifyRequest) (*npool.SetSigninVerifyResponse, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}

func (s *Server) SetInvitationCodeMust(ctx context.Context, in *npool.SetInvitationCodeMustRequest) (*npool.SetInvitationCodeMustResponse, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED")
}
