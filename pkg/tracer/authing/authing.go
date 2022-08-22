package authing

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing"
)

func trace(span trace1.Span, in *npool.ExistAuthRequest, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("Resource.%v", index), in.GetResource()),
		attribute.String(fmt.Sprintf("Method.%v", index), in.GetMethod()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.ExistAuthRequest) trace1.Span {
	return trace(span, in, 0)
}
