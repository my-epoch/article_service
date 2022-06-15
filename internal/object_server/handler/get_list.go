package handler

import (
	"context"
	pb "github.com/my-epoch/object_service/gen/go/api/proto/v1"
	"github.com/my-epoch/object_service/internal/database"
	"github.com/my-epoch/object_service/internal/dto"
)

func GetList(_ context.Context, request *pb.GetListRequest) (*pb.GetListResponse, error) {
	// handling zero value
	if request.Quantity == 0 {
		request.Quantity = 20
	}

	objects, err := objectRepository.GetList(request.Quantity, request.Offset)
	if err != nil {
		return nil, database.ErrorTransfer(err, "Object")
	}

	return &pb.GetListResponse{Objects: dto.ObjectModelSliceToPbShort(objects)}, nil
}
