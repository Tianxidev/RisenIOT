package models

import (
	"RisenIOT/backend/common/global"
	"RisenIOT/backend/utils"
)

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// GetUser 获取登录用户信息
func (u *Login) GetUser() (user SysUser, e error) {

	// 查询用户信息
	e = global.Eloquent.Table("sys_user").Where("username = ? ", u.Username).Find(&user).Error
	if e != nil {
		return
	}

	// 校验密码
	_, e = utils.CompareHashAndPassword(user.Password, u.Password)
	if e != nil {
		return
	}

	return
}
