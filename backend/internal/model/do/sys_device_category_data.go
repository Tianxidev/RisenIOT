// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDeviceCategoryData is the golang structure of table sys_device_category_data for DAO operations like Where/Data.
type SysDeviceCategoryData struct {
	g.Meta     `orm:"table:sys_device_category_data, do:true"`
	Id         interface{} //
	CategoryId interface{} //
	DeviceId   interface{} //
	DataInt    interface{} //
	DataStr    interface{} //
	DataDouble interface{} //
	CreatedAt  *gtime.Time //
}
