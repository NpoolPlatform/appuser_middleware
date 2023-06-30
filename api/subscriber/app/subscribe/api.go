package appsubscribe

import (
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber/app/subscribe"
	"google.golang.org/grpc"
)

type Server struct {
	subscribe.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	subscribe.RegisterMiddlewareServer(server, &Server{})
}
