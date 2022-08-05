package admin

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/admin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(info *admin.CreateGenesisUserRequest) error {
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "GetAppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	if info.UserID == nil {
		logger.Sugar().Errorw("validate", "UserID", info.UserID)
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Errorw("validate", "UserID", info.GetUserID(), "error", err)
		return status.Error(codes.InvalidArgument, "UserID is invalid")
	}

	if info.GetRole() == "" {
		logger.Sugar().Errorw("validate", "GetRole", info.GetRole())
		return status.Error(codes.InvalidArgument, "Role is empty")
	}

	if info.GetEmailAddress() == "" {
		logger.Sugar().Errorw("validate", "GetEmailAddress", info.GetEmailAddress())
		return status.Error(codes.InvalidArgument, "EmailAddress is empty")
	}

	if info.GetPasswordHash() == "" {
		logger.Sugar().Errorw("validate", "GetPasswordHash", info.GetPasswordHash())
		return status.Error(codes.InvalidArgument, "PasswordHash is empty")
	}

	return nil
}
