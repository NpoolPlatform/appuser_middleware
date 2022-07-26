package user

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/appusermw/user"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateUserSecret(in *user.CreateUserWithSecretRequest) error {
	if _, err := uuid.Parse(in.GetUser().GetAppID()); err != nil {
		logger.Sugar().Error("AppID is invalid")
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	if in.GetSecret().GetPasswordHash() == "" {
		logger.Sugar().Error("PasswordHash empty")
		return status.Error(codes.InvalidArgument, "PasswordHash empty")
	}

	if in.GetSecret().GetSalt() == "" {
		logger.Sugar().Error("Salt empty")
		return status.Error(codes.InvalidArgument, "Salt empty")
	}

	if in.GetSecret().GetGoogleSecret() == "" {
		logger.Sugar().Error("GoogleSecret is empty")
		return status.Error(codes.InvalidArgument, "GoogleSecret is empty")
	}
	return nil
}

func validateUserThirdParty(in *user.CreateUserWithThirdPartyRequest) error {
	if _, err := uuid.Parse(in.GetUser().GetAppID()); err != nil {
		logger.Sugar().Error("AppID is invalid")
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	if in.GetThirdParty().GetThirdPartyUserID() == "" {
		logger.Sugar().Error("ThirdPartyUserID empty")
		return status.Error(codes.InvalidArgument, "ThirdPartyUserID empty")
	}

	if in.GetThirdParty().GetThirdPartyID() == "" {
		logger.Sugar().Error("ThirdPartyID is invalid")
		return status.Error(codes.InvalidArgument, "ThirdPartyID is invalid")
	}

	return nil
}
