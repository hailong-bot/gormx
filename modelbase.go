package gormx

import (
	"reflect"

	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ModelBase struct {
	DataObjecter DataObjecter
}

func (m *ModelBase) GetByID(db *gorm.DB, id int64) (DataObjecter, error) {
	dataObjectType := reflect.TypeOf(m.DataObjecter)
	for dataObjectType.Kind() == reflect.Ptr {
		dataObjectType = dataObjectType.Elem()
	}

	dataObjecterValue := reflect.New(dataObjectType)
	result := dataObjecterValue.Interface().(DataObjecter)

	if dataObjecterValue.Elem().Kind() == reflect.Struct {
		doerField := dataObjecterValue.Elem().FieldByName("DataObjecter")
		if doerField.IsValid() && doerField.CanSet() {
			doerField.Set(dataObjecterValue)
		}
	}

	if err := db.Take(result, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return result, nil
}

func (m *ModelBase) GetByIDWithLock(db *gorm.DB, id int64, lock Lock) (DataObjecter, error) {
	// 1.生成该对象
	dataObjecterType := reflect.TypeOf(m.DataObjecter)
	for dataObjecterType.Kind() == reflect.Ptr {
		dataObjecterType = dataObjecterType.Elem()
	}
	dataObjecterValue := reflect.New(dataObjecterType)
	result := dataObjecterValue.Interface().(DataObjecter)

	// 2.为生成的对象设置 DataObjecter 值
	if dataObjecterValue.Elem().Kind() == reflect.Struct {
		doerField := dataObjecterValue.Elem().FieldByName("DataObjecter")
		if doerField.IsValid() && doerField.CanSet() {
			doerField.Set(dataObjecterValue)
		}
	}

	// 3.查找该对象
	query := db
	switch lock {
	case NoLock:
	case IS:
		query = query.Clauses(clause.Locking{Strength: "SHARE"})
	case IX:
		query = query.Clauses(clause.Locking{Strength: "UPDATE"})
	}
	if err := query.Take(result, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.WithStack(err)
	}
	return result, nil
}
