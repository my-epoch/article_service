package main

import (
	"github.com/my-epoch/api-gateway/pkg/logger"
	"github.com/my-epoch/object_service/internal/grpc_server"
	"github.com/my-epoch/object_service/pkg/consul"
	"github.com/my-epoch/object_service/pkg/service_config"
)

func main() {
	logger.Info("initializing")

	consul.RegisterService(service_config.Get())
	defer consul.DeregisterService()

	grpc_server.Serve()
}
