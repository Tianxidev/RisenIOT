// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDeviceCategoryData is the golang structure for table sys_device_category_data.
type SysDeviceCategoryData struct {
	Id         int         `json:"id"         ` //
	CategoryId int         `json:"categoryId" ` //
	DeviceId   int         `json:"deviceId"   ` //
	DataInt    int         `json:"dataInt"    ` //
	DataStr    string      `json:"dataStr"    ` //
	DataDouble float64     `json:"dataDouble" ` //
	CreatedAt  *gtime.Time `json:"createdAt"  ` //
}
