package user

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
)

func trace(span trace1.Span, in *npool.UserReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("EmailAddress.%v", index), in.GetEmailAddress()),
		attribute.String(fmt.Sprintf("PhoneNO.%v", index), in.GetPhoneNO()),
		attribute.String(fmt.Sprintf("ImportedFromAppID.%v", index), in.GetImportedFromAppID()),
		attribute.String(fmt.Sprintf("Username.%v", index), in.GetUsername()),
		attribute.StringSlice(fmt.Sprintf("AddressFields.%v", index), in.GetAddressFields()),
		attribute.String(fmt.Sprintf("Gender.%v", index), in.GetGender()),
		attribute.String(fmt.Sprintf("PostalCode.%v", index), in.GetPostalCode()),
		attribute.Int(fmt.Sprintf("Age.%v", index), int(in.GetAge())),
		attribute.Int(fmt.Sprintf("Birthday.%v", index), int(in.GetBirthday())),
		attribute.String(fmt.Sprintf("Avatar.%v", index), in.GetAvatar()),
		attribute.String(fmt.Sprintf("Organization.%v", index), in.GetOrganization()),
		attribute.String(fmt.Sprintf("FirstName.%v", index), in.GetFirstName()),
		attribute.String(fmt.Sprintf("LastName.%v", index), in.GetLastName()),
		attribute.String(fmt.Sprintf("IDNumber.%v", index), in.GetIDNumber()),
		attribute.Bool(fmt.Sprintf("GoogleAuthVerified.%v", index), in.GetGoogleAuthVerified()),
		attribute.String(fmt.Sprintf("PasswordHash.%v", index), in.GetPasswordHash()),
		attribute.String(fmt.Sprintf("GoogleSecret.%v", index), in.GetGoogleSecret()),
		attribute.String(fmt.Sprintf("ThirdPartyID.%v", index), in.GetThirdPartyID()),
		attribute.String(fmt.Sprintf("ThirdPartyUserID.%v", index), in.GetThirdPartyUserID()),
		attribute.String(fmt.Sprintf("ThirdPartyUsername.%v", index), in.GetThirdPartyUsername()),
		attribute.String(fmt.Sprintf("ThirdPartyAvatar.%v", index), in.GetThirdPartyAvatar()),
		attribute.Bool(fmt.Sprintf("Banned.%v", index), in.GetBanned()),
		attribute.String(fmt.Sprintf("BanMessage.%v", index), in.GetBanMessage()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.UserReq) trace1.Span {
	return trace(span, in, 0)
}
