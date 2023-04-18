package auth

import (
	"context"

	handler "github.com/NpoolPlatform/appuser-middleware/pkg/mw/authing/handler"
)

type Handler struct {
	*handler.Handler
}

func (h *Handler) VerifyConds(ctx context.Context, conds interface{}) error {
	return nil
}
