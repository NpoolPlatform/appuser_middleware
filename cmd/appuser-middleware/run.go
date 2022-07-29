package main

import (
	"github.com/NpoolPlatform/appuser-middleware/api"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	apimgrcli "github.com/NpoolPlatform/api-manager/pkg/client"

	cli "github.com/urfave/cli/v2"

	"google.golang.org/grpc"
)

var runCmd = &cli.Command{
	Name:    "run",
	Aliases: []string{"s"},
	Usage:   "Run the daemon",
	Action: func(c *cli.Context) error {
		if err := grpc2.RunGRPC(rpcRegister); err != nil {
			logger.Sugar().Errorf("fail to run grpc server: %v", err)
		}

		return nil
	},
}

func rpcRegister(server grpc.ServiceRegistrar) error {
	api.Register(server)

	apimgrcli.RegisterGRPC(server)

	return nil
}
