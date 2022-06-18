package handler

import (
	"context"
	"fmt"
	pb "github.com/my-epoch/object_service/gen/go/api/proto/v1"
	"github.com/my-epoch/object_service/internal/database"
	"github.com/my-epoch/object_service/internal/dto"
)

func GetNearest(_ context.Context, request *pb.GetNearestRequest) (*pb.GetNearestResponse, error) {
	objects, err := objectRepository.GetNearest(request.Latitude, request.Longitude, request.Radius)
	if err != nil {
		fmt.Println(err)
		return nil, database.ErrorTransfer(err, "Object")
	}

	return &pb.GetNearestResponse{Objects: dto.ObjectModelSliceToPbShort(objects)}, nil
}
