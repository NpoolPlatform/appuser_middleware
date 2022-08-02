package general

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/admin"
)

func trace(span trace1.Span, in *npool.CreateGenesisUserRequest, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("EmailAddress.%v", index), in.GetEmailAddress()),
		attribute.String(fmt.Sprintf("PasswordHash.%v", index), in.GetPasswordHash()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.CreateGenesisUserRequest) trace1.Span {
	return trace(span, in, 0)
}
