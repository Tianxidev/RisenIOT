// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDeviceKind is the golang structure for table sys_device_kind.
type SysDeviceKind struct {
	Id        int         `json:"id"        orm:"id"         ` // 产品类型ID
	Name      string      `json:"name"      orm:"name"       ` // 产品名称
	Mark      string      `json:"mark"      orm:"mark"       ` // 产品标记
	TimeOut   int         `json:"timeOut"   orm:"time_out"   ` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
	CreateBy  int         `json:"createBy"  orm:"create_by"  ` // 创建人
	Public    int         `json:"public"    orm:"public"     ` // 是否公开产品
}
