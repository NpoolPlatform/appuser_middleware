package app

import (
	"github.com/NpoolPlatform/message/npool/appusermw/app"
	"google.golang.org/grpc"
)

type Service struct {
	app.UnimplementedAppUserMiddlewareAppServer
}

func Register(server grpc.ServiceRegistrar) {
	app.RegisterAppUserMiddlewareAppServer(server, &Service{})
}
