// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLoginLog is the golang structure for table sys_login_log.
type SysLoginLog struct {
	InfoId        int64       `json:"infoId"        ` // 访问ID
	LoginName     string      `json:"loginName"     ` // 登录账号
	Ipaddr        string      `json:"ipaddr"        ` // 登录IP地址
	LoginLocation string      `json:"loginLocation" ` // 登录地点
	Browser       string      `json:"browser"       ` // 浏览器类型
	Os            string      `json:"os"            ` // 操作系统
	Status        int         `json:"status"        ` // 登录状态（0成功 1失败）
	Msg           string      `json:"msg"           ` // 提示消息
	LoginTime     *gtime.Time `json:"loginTime"     ` // 登录时间
	Module        string      `json:"module"        ` // 登录模块
}
