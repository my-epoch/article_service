package grpc_server

import (
	"fmt"
	"github.com/ShuffleBoy/goapp/logger"
	"github.com/my-epoch/article_service/internal/infrastructure/config"
	"github.com/my-epoch/proto/article_gen"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	Server   *grpc.Server
	Listener net.Listener
}

func NewArticleGrpcServer(config *config.Config, log *logger.Logger) *GrpcServer {
	grpcServer := grpc.NewServer()

	listenAddr := fmt.Sprintf("%s:%d", config.GrpcServerConfig.Addr, config.GrpcServerConfig.Port)
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("cannot create listener: %e", err)
	}
	article_gen.RegisterArticleServer(grpcServer, new(article_gen.UnimplementedArticleServer))
	return &GrpcServer{
		Server:   grpcServer,
		Listener: listener,
	}
}

func StartGrpcServer(server *GrpcServer, log *logger.Logger) {
	log.Info("starting grpc server")
	err := server.Server.Serve(server.Listener)
	if err != nil {
		log.Fatalf("cannot start grpc server: %e", err)
	}
	defer server.Server.Stop()
}
