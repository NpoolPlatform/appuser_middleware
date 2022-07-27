package api

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/api/admin"
	"github.com/NpoolPlatform/appuser-middleware/api/app"
	"github.com/NpoolPlatform/appuser-middleware/api/user"
	"github.com/NpoolPlatform/message/npool/appusermw"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Service struct {
	appusermw.UnimplementedAppUserMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	appusermw.RegisterAppUserMiddlewareServer(server, &Service{})
	app.Register(server)
	user.Register(server)
	admin.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := appusermw.RegisterAppUserMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
