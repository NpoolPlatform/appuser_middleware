package recoverycode

import (
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/user/recoverycode"
	"google.golang.org/grpc"
)

type Server struct {
	recoverycode.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	recoverycode.RegisterMiddlewareServer(server, &Server{})
}
