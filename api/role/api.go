package role

import (
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/role"
	"google.golang.org/grpc"
)

type Server struct {
	role.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	role.RegisterMiddlewareServer(server, &Server{})
}
