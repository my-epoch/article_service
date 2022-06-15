package handler

import (
	"context"
	pb "github.com/my-epoch/object_service/gen/go/api/proto/v1"
	"github.com/my-epoch/object_service/internal/database"
)

func GetById(_ context.Context, request *pb.GetByIdRequest) (*pb.GetByIdResponse, error) {
	object, err := objectRepository.GetById(request.Id)
	if err != nil {
		return nil, database.ErrorTransfer(err, "Object")
	}

	response := &pb.GetByIdResponse{Object: &pb.Object{
		Id:            object.ID,
		Title:         object.Title,
		MainImageUuid: object.MainImageUUID.String(),
		Description:   object.Description,
		Latitude:      object.Latitude,
		Longitude:     object.Longitude,
	}}

	return response, nil
}
