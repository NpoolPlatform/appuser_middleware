package admin

import (
	"github.com/NpoolPlatform/message/npool/appusermw/admin"
	"google.golang.org/grpc"
)

type Service struct {
	admin.UnimplementedAppUserMiddlewareAdminServer
}

func Register(server grpc.ServiceRegistrar) {
	admin.RegisterAppUserMiddlewareAdminServer(server, &Service{})
}
