package general

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
)

func trace(span trace1.Span, in *npool.AppReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("CreatedBy.%v", index), in.GetCreatedBy()),
		attribute.String(fmt.Sprintf("Name.%v", index), in.GetName()),
		attribute.String(fmt.Sprintf("Logo.%v", index), in.GetLogo()),
		attribute.String(fmt.Sprintf("Description.%v", index), in.GetDescription()),
		attribute.Bool(fmt.Sprintf("Banned.%v", index), in.GetBanned()),
		attribute.String(fmt.Sprintf("BanMessage.%v", index), in.GetBanMessage()),
		attribute.String(fmt.Sprintf("RecaptchaMethod.%v", index), in.GetRecaptchaMethod().String()),
		attribute.Bool(fmt.Sprintf("KycEnable.%v", index), in.GetKycEnable()),
		attribute.Bool(fmt.Sprintf("SigninVerifyEnable.%v", index), in.GetSigninVerifyEnable()),
		attribute.Bool(fmt.Sprintf("InvitationCodeMust.%v", index), in.GetInvitationCodeMust()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.AppReq) trace1.Span {
	return trace(span, in, 0)
}
