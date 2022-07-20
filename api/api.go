package api

import (
	"github.com/NpoolPlatform/message/npool/appusermiddleware/app"
	"google.golang.org/grpc"
)

type AppService struct {
	app.UnimplementedAppUserMiddlewareAppServer
}

func Register(server grpc.ServiceRegistrar) {
	app.RegisterAppUserMiddlewareAppServer(server, &AppService{})
}
