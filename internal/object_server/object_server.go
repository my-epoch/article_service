package object_server

import (
	"context"
	pb "github.com/my-epoch/object_service/gen/go/api/proto/v1"
	"github.com/my-epoch/object_service/internal/object_server/handler"
)

type ObjectServiceServer struct {
	pb.UnimplementedObjectServiceAPIServer
}

func (ObjectServiceServer) Status(ctx context.Context, request *pb.StatusRequest) (*pb.StatusResponse, error) {
	return handler.Status(ctx, request)
}
func (ObjectServiceServer) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	return handler.Create(ctx, request)
}
func (ObjectServiceServer) GetById(ctx context.Context, request *pb.GetByIdRequest) (*pb.GetByIdResponse, error) {
	return handler.GetById(ctx, request)
}
func (ObjectServiceServer) GetList(ctx context.Context, request *pb.GetListRequest) (*pb.GetListResponse, error) {
	return handler.GetList(ctx, request)
}
func (ObjectServiceServer) Search(ctx context.Context, request *pb.SearchRequest) (*pb.SearchResponse, error) {
	return handler.Search(ctx, request)
}
func (ObjectServiceServer) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	return handler.Update(ctx, request)
}
func (ObjectServiceServer) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return handler.Delete(ctx, request)
}
