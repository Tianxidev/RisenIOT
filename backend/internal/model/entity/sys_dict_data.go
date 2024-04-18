// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictData is the golang structure for table sys_dict_data.
type SysDictData struct {
	DictCode  int64       `json:"dictCode"  orm:"dict_code"  ` // 字典编码
	DictSort  int         `json:"dictSort"  orm:"dict_sort"  ` // 字典排序
	DictLabel string      `json:"dictLabel" orm:"dict_label" ` // 字典标签
	DictValue string      `json:"dictValue" orm:"dict_value" ` // 字典键值
	DictType  string      `json:"dictType"  orm:"dict_type"  ` // 字典类型
	CssClass  string      `json:"cssClass"  orm:"css_class"  ` // 样式属性（其他样式扩展）
	ListClass string      `json:"listClass" orm:"list_class" ` // 表格回显样式
	IsDefault int         `json:"isDefault" orm:"is_default" ` // 是否默认（1是 0否）
	Status    int         `json:"status"    orm:"status"     ` // 状态（0正常 1停用）
	CreateBy  uint64      `json:"createBy"  orm:"create_by"  ` // 创建者
	UpdateBy  uint64      `json:"updateBy"  orm:"update_by"  ` // 更新者
	Remark    string      `json:"remark"    orm:"remark"     ` // 备注
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 修改时间
}
