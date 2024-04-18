// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDeviceStatus is the golang structure for table sys_device_status.
type SysDeviceStatus struct {
	StatusId           int         `json:"statusId"           orm:"status_id"             ` //
	DeviceId           int         `json:"deviceId"           orm:"device_id"             ` //
	Status             int         `json:"status"             orm:"status"                ` //
	TimeOut            int         `json:"timeOut"            orm:"time_out"              ` //
	UpTime             int         `json:"upTime"             orm:"up_time"               ` //
	DownTime           int         `json:"downTime"           orm:"down_time"             ` //
	LastDataUpdateTime *gtime.Time `json:"lastDataUpdateTime" orm:"last_data_update_time" ` //
}
