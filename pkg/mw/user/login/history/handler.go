package history

import (
	"context"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

type Handler struct {
	ID        *uuid.UUID
	AppID     uuid.UUID
	UserID    uuid.UUID
	ClinetIP  string
	UserAgent string
	Location  string
	LoginType basetypes.LoginType
	Conds     *historycrud.Conds
	Offset    int32
	Limit     int32
}
