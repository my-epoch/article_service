package repository

import (
	"fmt"
	"github.com/my-epoch/object_service/internal/database"
	"github.com/my-epoch/object_service/model"
	"gorm.io/gorm"
	"strings"
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
	result := or.db.Take(&object, "id = ?", id)
	return &object, result.Error
}

func (or *ObjectRepository) Save(object *model.Object) error {
	result := or.db.Save(&object)
	return result.Error
}

func (or *ObjectRepository) GetList(quantity int32, offset int32) ([]model.Object, error) {
	var objects []model.Object
	result := or.db.
		Limit(int(quantity)).
		Offset(int(offset)).
		Order("title").
		Find(&objects)

	return objects, result.Error
}

func (or *ObjectRepository) GetNearest(latitude float32, longitude float32, radius float32) ([]model.Object, error) {
	var objects []model.Object
	result := or.db.Where(
		"sqrt( pow( 69.1 * (latitude - ?), 2 ) + pow( 69.1 * (? - longitude) * cos(latitude / 57.3), 2 ) ) < ? * 1609",
		latitude,
		longitude,
		radius,
	).Find(&objects)

	return objects, result.Error
}

func (or *ObjectRepository) FTS(query string) ([]model.Object, error) {
	words := strings.Fields(query)
	indexFields := []string{"title", "description", "address"}

	var parameters []interface{}
	var whereStatements []string
	for _, word := range words {
		for _, indexField := range indexFields {
			stmt := fmt.Sprintf("to_tsvector(\"%s\") @@ plainto_tsquery(?)", indexField)
			whereStatements = append(whereStatements, stmt)
			parameters = append(parameters, word)
		}
	}

	whereStatement := strings.Join(whereStatements, " OR ")

	var objects []model.Object
	result := or.db.Where(whereStatement, parameters...).Find(&objects)
	return objects, result.Error
}
