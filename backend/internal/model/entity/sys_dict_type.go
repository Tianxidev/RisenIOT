// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictType is the golang structure for table sys_dict_type.
type SysDictType struct {
	DictId    uint64      `json:"dictId"    ` // 字典主键
	DictName  string      `json:"dictName"  ` // 字典名称
	DictType  string      `json:"dictType"  ` // 字典类型
	Status    uint        `json:"status"    ` // 状态（0正常 1停用）
	CreateBy  uint        `json:"createBy"  ` // 创建者
	UpdateBy  uint        `json:"updateBy"  ` // 更新者
	Remark    string      `json:"remark"    ` // 备注
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建日期
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 修改日期
}
