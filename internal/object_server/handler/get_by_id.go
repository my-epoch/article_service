package handler

import (
	"context"
	pb "github.com/my-epoch/object_service/gen/go/api/proto/v1"
	"github.com/my-epoch/object_service/internal/database"
	"github.com/my-epoch/object_service/internal/dto"
)

func GetById(_ context.Context, request *pb.GetByIdRequest) (*pb.GetByIdResponse, error) {
	object, err := objectRepository.GetById(request.Id)
	if err != nil {
		return nil, database.ErrorTransfer(err, "Object")
	}

	response := &pb.GetByIdResponse{Object: dto.ObjectModelToPb(object)}
	return response, nil
}
