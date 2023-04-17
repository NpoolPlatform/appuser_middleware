package auth

import (
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/auth"
	"google.golang.org/grpc"
)

type Server struct {
	auth.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	auth.RegisterMiddlewareServer(server, &Server{})
}
