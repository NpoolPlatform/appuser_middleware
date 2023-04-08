package kyc

import (
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/kyc"
	"google.golang.org/grpc"
)

type Server struct {
	kyc.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	kyc.RegisterMiddlewareServer(server, &Server{})
}
