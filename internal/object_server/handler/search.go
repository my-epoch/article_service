package handler

import (
	"context"
	"fmt"
	pb "github.com/my-epoch/object_service/gen/go/api/proto/v1"
	"github.com/my-epoch/object_service/internal/database"
	"github.com/my-epoch/object_service/internal/dto"
)

func Search(_ context.Context, request *pb.SearchRequest) (*pb.SearchResponse, error) {
	objects, err := objectRepository.FTS(request.Query)
	if err != nil {
		fmt.Println(err)
		return nil, database.ErrorTransfer(err, "Object")
	}

	return &pb.SearchResponse{Objects: dto.ObjectModelSliceToPbShort(objects)}, nil
}
