// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDeviceInfo is the golang structure for table sys_device_info.
type SysDeviceInfo struct {
	Id                 int         `json:"id"                 orm:"id"                    ` //
	Name               string      `json:"name"               orm:"name"                  ` //
	Group              int         `json:"group"              orm:"group"                 ` //
	Sn                 string      `json:"sn"                 orm:"sn"                    ` //
	Pwd                string      `json:"pwd"                orm:"pwd"                   ` //
	Kind               int         `json:"kind"               orm:"kind"                  ` //
	Logo               string      `json:"logo"               orm:"logo"                  ` //
	Monitor            int         `json:"monitor"            orm:"monitor"               ` //
	Location           int         `json:"location"           orm:"location"              ` //
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"            ` //
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"            ` //
	Status             int         `json:"status"             orm:"status"                ` //
	TimeOut            int         `json:"timeOut"            orm:"time_out"              ` //
	StatusId           int         `json:"statusId"           orm:"status_id"             ` //
	UpTime             *gtime.Time `json:"upTime"             orm:"up_time"               ` // 上线时间
	DownTime           *gtime.Time `json:"downTime"           orm:"down_time"             ` // 离线时间
	LastDataUpdateTime *gtime.Time `json:"lastDataUpdateTime" orm:"last_data_update_time" ` // 最新一次数据更新时间
	CreateBy           int         `json:"createBy"           orm:"create_by"             ` // 创建人ID
}
