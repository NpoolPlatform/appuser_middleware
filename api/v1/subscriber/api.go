package subscriber

import (
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"
	"google.golang.org/grpc"
)

type Server struct {
	subscriber.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	subscriber.RegisterMiddlewareServer(server, &Server{})
}
