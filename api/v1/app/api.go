package app

import (
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	"google.golang.org/grpc"
)

type Server struct {
	app.UnimplementedAppMwServer
}

func Register(server grpc.ServiceRegistrar) {
	app.RegisterAppMwServer(server, &Server{})
}
