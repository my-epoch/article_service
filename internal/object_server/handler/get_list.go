package handler

import (
	"context"
	pb "github.com/my-epoch/object_service/gen/go/api/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetList(context.Context, *pb.GetListRequest) (*pb.GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
