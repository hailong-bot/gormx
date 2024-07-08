package mysql

import "github.com/hailong-bot/gormx"

type CpeModel struct {
	gormx.ModelBase
}

func NewCpeModel() *CpeModel {
	r := CpeModel{ModelBase: gormx.ModelBase{DataObjecter: &CpeDO{}}}
	return &r
}
