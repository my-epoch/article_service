package main

import (
	_ "github.com/my-epoch/object_service/internal/database"
	"github.com/my-epoch/object_service/internal/grpc_server"
	"github.com/my-epoch/object_service/pkg/consul"
	"github.com/my-epoch/object_service/pkg/service_config"
)

func main() {
	consul.RegisterService(service_config.Get())
	defer consul.DeregisterService()

	grpc_server.Serve()
}
