package gormx

import (
	"reflect"

	"gorm.io/gorm"
)

type ModelBase struct {
	DataObjecter DataObjecter
}

func (m *ModelBase) GetByID(db *gorm.DB, id int64) (DataObjecter, error) {
	dataObjectType := reflect.TypeOf(m.DataObjecter)
	for dataObjectType.Kind() == reflect.Ptr {
		dataObjectType = dataObjectType.Elem()
	}
}
