package kyc

import (
	"context"

	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	kyc.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	kyc.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return kyc.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
