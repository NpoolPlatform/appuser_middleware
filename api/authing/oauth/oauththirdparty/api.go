package oauththirdparty

import (
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/oauth/oauththirdparty"
	"google.golang.org/grpc"
)

type Server struct {
	oauththirdparty.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	oauththirdparty.RegisterMiddlewareServer(server, &Server{})
}
