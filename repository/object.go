package repository

import (
	"github.com/my-epoch/object_service/internal/database"
	"github.com/my-epoch/object_service/model"
	"gorm.io/gorm"
)

type ObjectRepository struct {
	db *gorm.DB
}

func NewObjectRepository() *ObjectRepository {
	return &ObjectRepository{db: database.GetConnection()}
}

func (or *ObjectRepository) Create(object *model.Object) error {
	result := or.db.Create(&object)
	return result.Error
}

func (or *ObjectRepository) GetById(id uint32) (*model.Object, error) {
	var object model.Object
	result := or.db.First(&object, "id = ?", id)
	return &object, result.Error
}

func (or *ObjectRepository) Save(object *model.Object) error {
	result := or.db.Save(&object)
	return result.Error
}
