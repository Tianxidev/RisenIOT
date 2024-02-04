// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserOnline is the golang structure for table sys_user_online.
type SysUserOnline struct {
	Id         uint64      `json:"id"         ` //
	Uuid       string      `json:"uuid"       ` // 用户标识
	Token      string      `json:"token"      ` // 用户token
	CreateTime *gtime.Time `json:"createTime" ` // 登录时间
	UserName   string      `json:"userName"   ` // 用户名
	Ip         string      `json:"ip"         ` // 登录ip
	Explorer   string      `json:"explorer"   ` // 浏览器
	Os         string      `json:"os"         ` // 操作系统
}
