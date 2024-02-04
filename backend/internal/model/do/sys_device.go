// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysDevice is the golang structure of table sys_device for DAO operations like Where/Data.
type SysDevice struct {
	g.Meta     `orm:"table:sys_device, do:true"`
	Id         interface{} //
	DeviceName interface{} // 设备名称
	DeviceUuid interface{} // 设备UUID
	DeviceUser interface{} // 设备所属用户
	DeviceType interface{} // 设备类型
}
