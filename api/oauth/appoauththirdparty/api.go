package appoauththirdparty

import (
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/oauth/appoauththirdparty"
	"google.golang.org/grpc"
)

type Server struct {
	appoauththirdparty.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	appoauththirdparty.RegisterMiddlewareServer(server, &Server{})
}
