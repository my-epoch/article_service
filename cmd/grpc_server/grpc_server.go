package main

import (
	"github.com/my-epoch/api-gateway/pkg/logger"
	"object_service/internal/gserver"
	"object_service/pkg/consul"
	"object_service/pkg/service_config"
)

func main() {
	logger.Info("initializing")

	consul.RegisterService(service_config.Get())
	defer consul.DeregisterService()

	gserver.Serve()
}
