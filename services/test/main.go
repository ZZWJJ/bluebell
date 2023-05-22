package main

import (
	"bluebell/services/test/handler"
	pb "bluebell/services/test/proto"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	service "github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/micro/micro/v3/service/logger"
)

// consul 2
const (
	ServerName = "consul-test"
	ConsulAddr = "localhost:38500"
)

func main() {
	// consul 3
	consulReg := consul.NewRegistry(
		registry.Addrs(ConsulAddr),
	)

	// Create service
	//srv := service.NewService(
	//	service.Name("test"),
	//)

	srv := service.NewService(
		service.Name(ServerName),    // 服务名字
		service.Registry(consulReg), // 注册中心
	)

	// Register handler
	_ = pb.RegisterTestHandler(srv.Server(), new(handler.Test))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
