// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysDeviceStatus is the golang structure of table sys_device_status for DAO operations like Where/Data.
type SysDeviceStatus struct {
	g.Meta             `orm:"table:sys_device_status, do:true"`
	StatusId           interface{} //
	DeviceId           interface{} //
	Status             interface{} //
	TimeOut            interface{} //
	UpTime             interface{} //
	DownTime           interface{} //
	LastDataUpdateTime interface{} //
}
