package grpc_server

import (
	"fmt"
	"github.com/my-epoch/api-gateway/pkg/logger"
	gen "github.com/my-epoch/object_service/gen/go/api/proto/v1"
	"github.com/my-epoch/object_service/internal/health_server"
	"github.com/my-epoch/object_service/internal/object_server"
	"github.com/my-epoch/object_service/pkg/service_config"
	"google.golang.org/grpc"
	grpcHealth "google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

func Serve() {
	grpcServer := grpc.NewServer()

	grpcHealth.RegisterHealthServer(grpcServer, &health_server.HealthServer{})

	gen.RegisterObjectServiceAPIServer(grpcServer, &object_server.ObjectServiceServer{})

	cfg := service_config.Get()
	listenAddr := fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port)

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		logger.Fatalf("cannot start server: %e", err)
	}

	logger.Infof("gRPC server starts listening on '%s'", cfg.Addr)
	defer grpcServer.Stop()
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("cannot start gRPC server: %e", err)
	}

}
