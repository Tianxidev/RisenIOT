// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDeviceInfo is the golang structure of table sys_device_info for DAO operations like Where/Data.
type SysDeviceInfo struct {
	g.Meta             `orm:"table:sys_device_info, do:true"`
	Id                 interface{} //
	Name               interface{} //
	Group              interface{} //
	Sn                 interface{} //
	Pwd                interface{} //
	Kind               interface{} //
	Logo               interface{} //
	Monitor            interface{} //
	Location           interface{} //
	CreatedAt          *gtime.Time //
	UpdatedAt          *gtime.Time //
	Status             interface{} //
	TimeOut            interface{} //
	StatusId           interface{} //
	UpTime             *gtime.Time // 上线时间
	DownTime           *gtime.Time // 离线时间
	LastDataUpdateTime *gtime.Time // 最新一次数据更新时间
	CreateBy           interface{} // 创建人ID
}
