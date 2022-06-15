package dto

import (
	pb "github.com/my-epoch/object_service/gen/go/api/proto/v1"
	"github.com/my-epoch/object_service/model"
)

func ObjectModelToPb(object *model.Object) *pb.Object {
	return &pb.Object{
		Id:            object.ID,
		Title:         object.Title,
		MainImageUuid: object.MainImageUUID.String(),
		Description:   object.Description,
		Latitude:      object.Latitude,
		Longitude:     object.Longitude,
	}
}
