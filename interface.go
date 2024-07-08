package gormx

import "gorm.io/gorm"

type UPO map[string]interface{}

type Modeler interface {
	GetByID(db *gorm.DB, id int64) (DataObjecter, error)
	GetByIDWithLock(db *gorm.DB, id int64, lock Lock) (DataObjecter, error)
	InsertBatch(db *gorm.DB, doList interface{}) error
}

type DataObjecter interface {
	GetIDer() interface{}
	Updates(db *gorm.DB, values UPO) error
	Insert(db *gorm.DB) error
	Delete(db *gorm.DB) error
}
