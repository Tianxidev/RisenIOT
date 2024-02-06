// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDeviceKind is the golang structure of table sys_device_kind for DAO operations like Where/Data.
type SysDeviceKind struct {
	g.Meta    `orm:"table:sys_device_kind, do:true"`
	Id        interface{} //
	Name      interface{} //
	Mark      interface{} //
	TimeOut   interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
