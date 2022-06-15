package handler

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/my-epoch/object_service/gen/go/api/proto/v1"
	"github.com/my-epoch/object_service/internal/database"
	"github.com/my-epoch/object_service/model"
	"github.com/my-epoch/object_service/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var objectRepository = repository.NewObjectRepository()

func Create(_ context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	mainImageUuid, err := uuid.Parse(request.MainImageUuid)
	if err != nil {
		return nil, status.Error(codes.Internal, "`mainImageUuid` has wrong format")
	}

	object := model.Object{
		Title:         request.Title,
		MainImageUUID: mainImageUuid,
		Description:   request.Description,
		Latitude:      request.Latitude,
		Longitude:     request.Longitude,
	}

	if err := objectRepository.Create(&object); err != nil {
		return nil, database.ErrorTransfer(err, "Object")
	}

	response := &pb.CreateResponse{Object: &pb.Object{
		Id:            object.ID,
		Title:         object.Title,
		MainImageUuid: object.MainImageUUID.String(),
		Description:   object.Description,
		Latitude:      object.Latitude,
		Longitude:     object.Longitude,
	}}

	return response, nil
}
