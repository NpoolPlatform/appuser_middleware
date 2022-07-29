package user

import (
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
	"google.golang.org/grpc"
)

type Server struct {
	user.UnimplementedUserMwServer
}

func Register(server grpc.ServiceRegistrar) {
	user.RegisterUserMwServer(server, &Server{})
}
