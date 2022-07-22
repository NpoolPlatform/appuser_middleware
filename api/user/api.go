package user

import (
	"github.com/NpoolPlatform/message/npool/appusermw/user"
	"google.golang.org/grpc"
)

type Service struct {
	user.UnimplementedAppUserMiddlewareUserServer
}

func Register(server grpc.ServiceRegistrar) {
	user.RegisterAppUserMiddlewareUserServer(server, &Service{})
}
