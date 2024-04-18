// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDeviceCategoryData is the golang structure for table sys_device_category_data.
type SysDeviceCategoryData struct {
	Id         int         `json:"id"         orm:"id"          ` //
	CategoryId int         `json:"categoryId" orm:"category_id" ` //
	DeviceId   int         `json:"deviceId"   orm:"device_id"   ` //
	DataInt    int         `json:"dataInt"    orm:"data_int"    ` //
	DataStr    string      `json:"dataStr"    orm:"data_str"    ` //
	DataDouble float64     `json:"dataDouble" orm:"data_double" ` //
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  ` //
}
