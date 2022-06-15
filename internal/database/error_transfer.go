package database

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func ErrorTransfer(err error, modelName string) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return status.Errorf(codes.NotFound, "Target '%s' record not found", modelName)
	}
	return status.Error(codes.Internal, "Something went wrong")
}
