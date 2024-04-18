// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDeviceCategory is the golang structure for table sys_device_category.
type SysDeviceCategory struct {
	Id        int         `json:"id"        orm:"id"         ` //
	KindId    int         `json:"kindId"    orm:"kind_id"    ` // 产品类型ID
	Name      string      `json:"name"      orm:"name"       ` //
	Mark      string      `json:"mark"      orm:"mark"       ` //
	DataType  int         `json:"dataType"  orm:"data_type"  ` // 数据类型
	LogicType int         `json:"logicType" orm:"logic_type" ` //
	Unit      string      `json:"unit"      orm:"unit"       ` // 单位
	Ratio     string      `json:"ratio"     orm:"ratio"      ` // 比例
	Format    string      `json:"format"    orm:"format"     ` //
	HomeShow  int         `json:"homeShow"  orm:"home_show"  ` //
	Remark    string      `json:"remark"    orm:"remark"     ` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
	CreateBy  int         `json:"createBy"  orm:"create_by"  ` // 创建人
}
