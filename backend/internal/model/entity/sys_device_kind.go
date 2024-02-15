// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDeviceKind is the golang structure for table sys_device_kind.
type SysDeviceKind struct {
	Id        int         `json:"id"        ` // 产品类型ID
	Name      string      `json:"name"      ` // 产品名称
	Mark      string      `json:"mark"      ` // 产品标记
	TimeOut   int         `json:"timeOut"   ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
	CreateBy  int         `json:"createBy"  ` // 创建人
	Public    int         `json:"public"    ` // 是否公开产品
}
