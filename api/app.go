package api

import (
	"context"
	"encoding/json"

	constant "github.com/NpoolPlatform/appuser-middleware/pkg/message/const"
	"github.com/google/uuid"

	appmw "github.com/NpoolPlatform/appuser-middleware/pkg/middleware/app"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/appusermiddleware/app"
	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *AppService) GetAppInfo(ctx context.Context, in *app.GetAppInfoRequest) (*app.GetAppInfoResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAppWithBanApp")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Error("ID is invalid")
		return &app.GetAppInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	resp, err := appmw.GetAppInfo(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("fail get app with ban app: %v", err)
		return &app.GetAppInfoResponse{}, status.Error(codes.Internal, err.Error())
	}

	isBanApp := false
	if resp.BanAppAppID != "" {
		isBanApp = true
	}

	var signupMethods []string
	err = json.Unmarshal([]byte(resp.SignupMethods), &signupMethods)
	if err != nil {
		logger.Sugar().Errorw("json unmarshal fail: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	var externSigninMethods []string
	err = json.Unmarshal([]byte(resp.ExternSigninMethods), &externSigninMethods)
	if err != nil {
		logger.Sugar().Errorw("json unmarshal fail: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	kycEnable := false
	if resp.KycEnable == 1 {
		kycEnable = true
	}

	signinVerifyEnable := false
	if resp.SigninVerifyEnable == 1 {
		signinVerifyEnable = true
	}

	invitationCodeMust := false
	if resp.InvitationCodeMust == 1 {
		invitationCodeMust = true
	}

	return &app.GetAppInfoResponse{
		Info: &app.AppInfo{
			Id:                  resp.ID,
			CreatedBy:           resp.CreatedBy,
			Name:                resp.Name,
			Logo:                resp.Logo,
			CreatedAt:           resp.CreatedAt,
			Description:         resp.Description,
			IsBanApp:            isBanApp,
			BanAppMessage:       resp.BanAppMessage,
			SignupMethods:       signupMethods,
			ExternSigninMethods: externSigninMethods,
			RecaptchaMethod:     resp.RecaptchaMethod,
			KycEnable:           kycEnable,
			SigninVerifyEnable:  signinVerifyEnable,
			InvitationCodeMust:  invitationCodeMust,
		},
	}, nil
}
