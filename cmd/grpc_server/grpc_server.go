package main

import (
	"github.com/my-epoch/api-gateway/pkg/consul"
	"github.com/my-epoch/api-gateway/pkg/logger"
	"github.com/my-epoch/api-gateway/pkg/service_config"
	"object_service/internal/gserver"
)

func main() {
	logger.Info("initializing")

	consul.RegisterService(service_config.Get())
	defer consul.DeregisterService()

	gserver.Serve()
}
