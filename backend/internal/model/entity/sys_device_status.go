// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDeviceStatus is the golang structure for table sys_device_status.
type SysDeviceStatus struct {
	StatusId           int         `json:"statusId"           ` //
	DeviceId           int         `json:"deviceId"           ` //
	Status             int         `json:"status"             ` //
	TimeOut            int         `json:"timeOut"            ` //
	UpTime             int         `json:"upTime"             ` //
	DownTime           int         `json:"downTime"           ` //
	LastDataUpdateTime *gtime.Time `json:"lastDataUpdateTime" ` //
}
