// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysDeviceCategory is the golang structure of table sys_device_category for DAO operations like Where/Data.
type SysDeviceCategory struct {
	g.Meta    `orm:"table:sys_device_category, do:true"`
	Id        interface{} //
	KindId    interface{} // 类型ID
	Name      interface{} //
	Mark      interface{} //
	DataType  interface{} // 数据类型
	LogicType interface{} //
	Unit      interface{} // 单位
	Ratio     interface{} // 比例
	Format    interface{} //
	HomeShow  interface{} //
	Remark    interface{} //
	CreatedAt interface{} //
	UpdatedAt interface{} //
}
