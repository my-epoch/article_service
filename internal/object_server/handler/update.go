package handler

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/my-epoch/object_service/gen/go/api/proto/v1"
	"github.com/my-epoch/object_service/internal/database"
	"github.com/my-epoch/object_service/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Update(_ context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	object, err := objectRepository.GetById(request.Id)
	if err != nil {
		return nil, database.ErrorTransfer(err, "Object")
	}

	mainImageUuid, err := uuid.Parse(request.MainImageUuid)
	if err != nil {
		return nil, status.Error(codes.Internal, "`mainImageUuid` has wrong format")
	}

	object.Title = request.Title
	object.MainImageUUID = mainImageUuid
	object.Description = request.Description
	object.Latitude = request.Latitude
	object.Longitude = request.Longitude

	if err := objectRepository.Save(object); err != nil {
		return nil, database.ErrorTransfer(err, "Object")
	}

	response := &pb.UpdateResponse{Object: dto.ObjectModelToPb(object)}
	return response, nil
}
