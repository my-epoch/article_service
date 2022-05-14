package gserver

import (
	"context"
	pb "github.com/my-epoch/api-gateway/gen/go/api/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ObjectServiceServer struct {
	pb.UnimplementedObjectServiceServer
}

func (ObjectServiceServer) GetAll(_ context.Context, _ *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Method not implemented yet")
}

func (ObjectServiceServer) Search(_ context.Context, _ *pb.SearchRequest) (*pb.SearchResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Method not implemented yet")
}
