package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Object struct {
	gorm.Model
	ID            uint32 `gorm:"primarykey"`
	Title         string
	MainImageUUID uuid.UUID
	Description   string

	Latitude  float32
	Longitude float32
}
