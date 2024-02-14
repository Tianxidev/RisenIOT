// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysDeviceGroup is the golang structure of table sys_device_group for DAO operations like Where/Data.
type SysDeviceGroup struct {
	g.Meta  `orm:"table:sys_device_group, do:true"`
	Id      interface{} // 分组ID
	Name    interface{} // 分组名称
	Remarks interface{} // 备注
}
