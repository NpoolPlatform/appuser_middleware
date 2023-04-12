package authing

import (
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing"
	"google.golang.org/grpc"
)

type Server struct {
	authing.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	authing.RegisterMiddlewareServer(server, &Server{})
}
