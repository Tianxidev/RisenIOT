package models

import (
	"RisenIOT/backend/common/global"
	"errors"
)

type SysRole struct {
	RoleID   int    `gorm:"column:rid" gorm:"primary_key;" gorm:"size:11" json:"rid"` // 角色ID
	RoleName string `gorm:"column:role_name" gorm:"size:64" json:"role_name"`         // 角色名称
}

// TableName 表名
func (SysRole) TableName() string {
	return "sys_role"
}

// Insert 添加用户
func (e SysRole) Insert() (id int, err error) {

	// check 用户名
	var count int64
	global.Eloquent.Table(e.TableName()).Where("rid = ?", e.RoleID).Count(&count)
	if count > 0 {
		err = errors.New("角色已存在！")
		return
	}

	// 添加数据
	if err = global.Eloquent.Table(e.TableName()).Create(&e).Error; err != nil {
		return
	}
	id = e.RoleID
	return
}
