// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictType is the golang structure for table sys_dict_type.
type SysDictType struct {
	DictId    uint64      `json:"dictId"    orm:"dict_id"    ` // 字典主键
	DictName  string      `json:"dictName"  orm:"dict_name"  ` // 字典名称
	DictType  string      `json:"dictType"  orm:"dict_type"  ` // 字典类型
	Status    uint        `json:"status"    orm:"status"     ` // 状态（0正常 1停用）
	CreateBy  uint        `json:"createBy"  orm:"create_by"  ` // 创建者
	UpdateBy  uint        `json:"updateBy"  orm:"update_by"  ` // 更新者
	Remark    string      `json:"remark"    orm:"remark"     ` // 备注
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建日期
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 修改日期
}
