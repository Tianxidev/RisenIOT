// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDeviceInfo is the golang structure for table sys_device_info.
type SysDeviceInfo struct {
	Id                 int         `json:"id"                 ` //
	Name               string      `json:"name"               ` //
	Group              int         `json:"group"              ` //
	Sn                 string      `json:"sn"                 ` //
	Pwd                string      `json:"pwd"                ` //
	Kind               int         `json:"kind"               ` //
	Logo               string      `json:"logo"               ` //
	Monitor            int         `json:"monitor"            ` //
	Location           int         `json:"location"           ` //
	CreatedAt          *gtime.Time `json:"createdAt"          ` //
	UpdatedAt          *gtime.Time `json:"updatedAt"          ` //
	Status             int         `json:"status"             ` //
	TimeOut            int         `json:"timeOut"            ` //
	StatusId           int         `json:"statusId"           ` //
	UpTime             *gtime.Time `json:"upTime"             ` // 上线时间
	DownTime           *gtime.Time `json:"downTime"           ` // 离线时间
	LastDataUpdateTime *gtime.Time `json:"lastDataUpdateTime" ` // 最新一次数据更新时间
	CreateBy           int         `json:"createBy"           ` // 创建人ID
}
