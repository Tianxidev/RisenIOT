package models

import (
	"RisenIOT/backend/common/global"
	"errors"
)

// CasbinRule casbin_rule 表结构
type CasbinRule struct {
	ID    int    `gorm:"column:id" gorm:"primary_key;" json:"id"`
	PType string `gorm:"column:ptype" gorm:"size:100;" json:"ptype"` // 类型，可以是 p 策略，g 角色等等
	V0    string `gorm:"column:v0" gorm:"size:100;" json:"v0"`       // 角色 roleName/roleId   sub
	V1    string `gorm:"column:v1" gorm:"size:100;" json:"v1"`       // Path 路径              obj
	V2    string `gorm:"column:v2" gorm:"size:100;" json:"v2"`       // Method 请求方式         act
	V3    string `gorm:"column:v3" gorm:"size:100;" json:"v3"`       // 允许读/写        read/write
	V4    string `gorm:"column:v4" gorm:"size:100;" json:"v4"`       // 允许/拒绝        allow/deny
	V5    string `gorm:"column:v5" gorm:"size:100;" json:"v5"`
}

func (CasbinRule) TableName() string {
	return "casbin_rule"
}

func (e CasbinRule) Insert() (id int, err error) {
	var count int64

	// 判断v0,v1,v2 是否为空
	if e.V0 == "" || e.V1 == "" || e.V2 == "" {
		err = errors.New("参数错误！")
	}

	// 判断是否已存在该策略
	global.Eloquent.Table(e.TableName()).Where("v0 = ?", e.V0).Where("v1 = ?", e.V1).Where("v2 =?", e.V2).Count(&count)
	if count > 0 {
		err = errors.New("已存在该策略！")
		return
	}

	//添加数据
	if err = global.Eloquent.Table(e.TableName()).Create(&e).Error; err != nil {
		return
	}
	id = e.ID
	return
}
