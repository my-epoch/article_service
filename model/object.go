package model

import "gorm.io/gorm"

type Object struct {
	gorm.Model

	Title         string
	MainImageUUID string
	Description   string

	Latitude  float32
	Longitude float32
}
