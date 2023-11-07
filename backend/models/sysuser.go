package models

import (
	"RisenIOT/backend/common/global"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// SysUserId 用户ID
type SysUserId struct {
	UserId int `gorm:"column:uid" gorm:"primary_key;"  json:"uid"` // 用户ID
}

// SysUsername 用户名
type SysUsername struct {
	Username string `gorm:"column:username" gorm:"size:64" json:"username"` // 用户名
}

// SysPassword 密码
type SysPassword struct {
	Password string `gorm:"column:password" gorm:"size:128" json:"password"` // 密码
}

// LoginM 登录模型
type LoginM struct {
	SysUsername
	SysPassword
}

// SysUser 用户表结构
type SysUser struct {
	SysUserId
	LoginM
	RoleId int `gorm:"column:role_id" gorm:"size:11" json:"role_id"` // 角色ID
}

// TableName 表名
func (SysUser) TableName() string {
	return "sys_user"
}

// Encrypt 加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

// Insert 添加用户
func (e SysUser) Insert() (id int, err error) {
	if err = e.Encrypt(); err != nil {
		return
	}

	// check 用户名
	var count int64
	global.Eloquent.Table(e.TableName()).Where("username = ?", e.Username).Count(&count)
	if count > 0 {
		err = errors.New("账户已存在！")
		return
	}

	//添加数据
	if err = global.Eloquent.Table(e.TableName()).Create(&e).Error; err != nil {
		return
	}
	id = e.UserId
	return
}

// GetUserFromId 根据用户ID获取用户信息
func (e *SysUser) GetUserFromId() (user SysUser, err error) {

	// 查询用户信息
	err = global.Eloquent.Table("sys_user").Where("uid = ? ", e.UserId).Find(&user).Error
	if e != nil {
		return
	}

	return
}
