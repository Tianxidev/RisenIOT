// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDeviceKind is the golang structure for table sys_device_kind.
type SysDeviceKind struct {
	Id        int         `json:"id"        ` //
	Name      string      `json:"name"      ` //
	Mark      string      `json:"mark"      ` //
	TimeOut   int         `json:"timeOut"   ` //
	CreatedAt *gtime.Time `json:"createdAt" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" ` //
}
