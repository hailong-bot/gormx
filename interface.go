package gormx

import "gorm.io/gorm"

type UPO map[string]interface{}

type Modeler interface {
	GetByID(db *gorm.DB, id int64) (DataObjecter, error)
	GetByIDWithLock(db *gorm.DB, id int64, lock Lock) (DataObjecter, error)
	InsertBatch(db *gorm.DB, doList any) error
	GetByConditions(db *gorm.DB, where string, values ...any) (DataObjecter, error)
	GetByConditionsWithLock(db *gorm.DB, lock Lock, where string, values ...any) (DataObjecter, error)
	List(db *gorm.DB, offset int, limit int, sortField string, sort Sort, where string, values ...any) (DataObjecterList, error)
	ListAll(db *gorm.DB, sortField string, sort Sort, where string, values ...any) (DataObjecterList, error)
	Exist(db *gorm.DB, where string, values ...any) (bool, error)
	Count(db *gorm.DB, where string, values ...any) (int64, error)
	DeleteBatch(db *gorm.DB, where string, values ...any) error
	UpdateBatch(db *gorm.DB, updateParams UPO, where string, values ...any) error
}

type DataObjecter interface {
	GetIDer() any
	Updates(db *gorm.DB, values UPO) error
	Insert(db *gorm.DB) error
	Delete(db *gorm.DB) error
}
