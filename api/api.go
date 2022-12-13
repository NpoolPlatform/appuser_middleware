package api

import (
	"context"

	"github.com/NpoolPlatform/appuser-middleware/api/v1/kyc"

	"github.com/NpoolPlatform/appuser-middleware/api/v1/app"
	"github.com/NpoolPlatform/appuser-middleware/api/v1/authing"
	"github.com/NpoolPlatform/appuser-middleware/api/v1/role"
	"github.com/NpoolPlatform/appuser-middleware/api/v1/subscriber"
	"github.com/NpoolPlatform/appuser-middleware/api/v1/user"

	appusermw "github.com/NpoolPlatform/message/npool/appuser/mw/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Service struct {
	appusermw.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	appusermw.RegisterMiddlewareServer(server, &Service{})
	app.Register(server)
	subscriber.Register(server)
	user.Register(server)
	role.Register(server)
	authing.Register(server)
	kyc.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := appusermw.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
