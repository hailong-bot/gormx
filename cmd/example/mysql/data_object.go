package mysql

import (
	"github.com/hailong-bot/gormx"
	"time"
)

type CpeDO struct {
	gormx.DOBase

	FlowID     string    `gorm:"type:varchar(50);NOT NULL;column:flow_id;"`      // 流程ID
	ContractID string    `gorm:"type:varchar(50);NOT NULLL;column:contract_id;"` // 合同ID(普通合同和变更合同)
	UserID     int64     `gorm:"type:int(11);NOT NULL;column:user_id;"`          // 用户ID
	CreateTime time.Time `gorm:"type:datetime;NOT NULL;column:create_time;"`     // 创建时间
	UpdateTime time.Time `gorm:"type:datetime;NOT NULL;column:update_time;"`     // 更新时间
	IsAlter    int8      `gorm:"type:tinyint(4);NOT NULL;column:is_alter;"`      // 0：普通合同 1：变更合同
	IsDeleted  int8      `gorm:"type:tinyint(4);NOT NULL;column:is_deleted;"`    // 0 正常 1删除
}

func (d *CpeDO) TableName() string {
	return "cpe"
}
