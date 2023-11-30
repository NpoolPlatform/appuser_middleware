package user

import (
	"context"

	user1 "github.com/NpoolPlatform/appuser-middleware/pkg/mw/user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateUser(ctx context.Context, in *npool.UpdateUserRequest) (*npool.UpdateUserResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateUser",
			"Req", req,
		)
		return &npool.UpdateUserResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}
	handler, err := user1.NewHandler(
		ctx,
		user1.WithID(req.ID, true),
		user1.WithAppID(req.AppID, false),
		user1.WithPhoneNO(req.PhoneNO, false),
		user1.WithEmailAddress(req.EmailAddress, false),
		user1.WithImportFromAppID(req.ImportedFromAppID, false),
		user1.WithPasswordHash(req.PasswordHash, false),
		user1.WithFirstName(req.FirstName, false),
		user1.WithLastName(req.LastName, false),
		user1.WithBirthday(req.Birthday, false),
		user1.WithGender(req.Gender, false),
		user1.WithAvatar(req.Avatar, false),
		user1.WithUsername(req.Username, false),
		user1.WithPostalCode(req.PostalCode, false),
		user1.WithAge(req.Age, false),
		user1.WithOrganization(req.Organization, false),
		user1.WithIDNumber(req.IDNumber, false),
		user1.WithAddressFields(req.AddressFields, false),
		user1.WithGoogleSecret(req.GoogleSecret, false),
		user1.WithGoogleAuthVerified(req.GoogleAuthVerified, false),
		user1.WithKol(req.Kol, false),
		user1.WithKolConfirmed(req.KolConfirmed, false),
		user1.WithSelectedLangID(req.SelectedLangID, false),
		user1.WithActionCredits(req.ActionCredits, false),
		user1.WithSigninVerifyType(req.SigninVerifyType, false),
		user1.WithBanned(req.Banned, false),
		user1.WithBanMessage(req.BanMessage, false),
		user1.WithThirdPartyID(req.ThirdPartyID, false),
		user1.WithThirdPartyUserID(req.ThirdPartyUserID, false),
		user1.WithThirdPartyUsername(req.ThirdPartyUsername, false),
		user1.WithThirdPartyAvatar(req.ThirdPartyAvatar, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateUser",
			"Req", req,
			"error", err,
		)
		return &npool.UpdateUserResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.UpdateUser(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateUser",
			"Req", req,
			"error", err,
		)
		return &npool.UpdateUserResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateUserResponse{
		Info: info,
	}, nil
}
