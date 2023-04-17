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
	handler, err := user1.NewHandler(
		ctx,
		user1.WithID(req.ID),
		user1.WithAppID(req.GetAppID()),
		user1.WithPhoneNO(req.PhoneNO),
		user1.WithEmailAddress(req.EmailAddress),
		user1.WithImportedFromAppID(req.ImportedFromAppID),
		user1.WithPasswordHash(req.PasswordHash),
		user1.WithFirstName(req.FirstName),
		user1.WithLastName(req.LastName),
		user1.WithBirthday(req.Birthday),
		user1.WithGender(req.Gender),
		user1.WithAvatar(req.Avatar),
		user1.WithUsername(req.Username),
		user1.WithPostalCode(req.PostalCode),
		user1.WithAge(req.Age),
		user1.WithOrganization(req.Organization),
		user1.WithIDNumber(req.IDNumber),
		user1.WithAddressFields(req.AddressFields),
		user1.WithGoogleSecret(req.GoogleSecret),
		user1.WithGoogleAuthVerified(req.GoogleAuthVerified),
		user1.WithKol(req.Kol),
		user1.WithKolConfirmed(req.KolConfirmed),
		user1.WithActionCredits(req.ActionCredits),
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
