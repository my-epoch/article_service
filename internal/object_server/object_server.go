package object_server

import (
	"context"
	pb "github.com/my-epoch/object_service/gen/go/api/proto/v1"
	"time"
)

type ObjectServiceServer struct {
	pb.UnimplementedObjectServiceAPIServer
}

func (ObjectServiceServer) Status(context.Context, *pb.StatusRequest) (*pb.StatusResponse, error) {
	return &pb.StatusResponse{
		Status: "ok",
		Time:   time.Now().Format(time.RFC3339),
	}, nil
}
