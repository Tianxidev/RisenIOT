// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RiseniotUser is the golang structure for table riseniot_user.
type RiseniotUser struct {
	Uuid       string      `json:"uuid"       ` // 用户uuid
	Username   string      `json:"username"   ` // 用户名, 唯一索引
	Password   string      `json:"password"   ` // 密码
	Nickname   string      `json:"nickname"   ` // 昵称
	Role       string      `json:"role"       ` // 用户角色
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
}
