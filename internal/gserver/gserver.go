package gserver

import (
	pb "github.com/my-epoch/api-gateway/gen/go/api/proto/v1"
	"github.com/my-epoch/api-gateway/pkg/logger"
	"github.com/my-epoch/api-gateway/pkg/service_config"
	"google.golang.org/grpc"
	grpcHealth "google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

func Serve() {
	grpcServer := grpc.NewServer()

	grpcHealth.RegisterHealthServer(grpcServer, &HealthServer{})
	pb.RegisterObjectServiceServer(grpcServer, &ObjectServiceServer{})
	cfg := service_config.Get()
	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		logger.Fatalf("cannot start server: %e", err)
	}

	logger.Infof("gRPC server starts listening on '%s'", cfg.Addr)
	defer grpcServer.Stop()
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("cannot start gRPC server: %e", err)
	}

}
