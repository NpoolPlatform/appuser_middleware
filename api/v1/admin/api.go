package admin

import (
	"github.com/NpoolPlatform/message/npool/appuser/mw/v1/admin"
	"google.golang.org/grpc"
)

type Service struct {
	admin.UnimplementedAdminMwServer
}

func Register(server grpc.ServiceRegistrar) {
	admin.RegisterAdminMwServer(server, &Service{})
}
