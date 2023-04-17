package history

import (
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/history"
	"google.golang.org/grpc"
)

type Server struct {
	history.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	history.RegisterMiddlewareServer(server, &Server{})
}
