// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDept is the golang structure for table sys_dept.
type SysDept struct {
	DeptId    int64       `json:"deptId"    orm:"dept_id"    ` // 部门id
	ParentId  int64       `json:"parentId"  orm:"parent_id"  ` // 父部门id
	Ancestors string      `json:"ancestors" orm:"ancestors"  ` // 祖级列表
	DeptName  string      `json:"deptName"  orm:"dept_name"  ` // 部门名称
	OrderNum  int         `json:"orderNum"  orm:"order_num"  ` // 显示顺序
	Leader    string      `json:"leader"    orm:"leader"     ` // 负责人
	Phone     string      `json:"phone"     orm:"phone"      ` // 联系电话
	Email     string      `json:"email"     orm:"email"      ` // 邮箱
	Status    uint        `json:"status"    orm:"status"     ` // 部门状态（0正常 1停用）
	CreatedBy uint64      `json:"createdBy" orm:"created_by" ` // 创建人
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" ` // 修改人
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 修改时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" ` // 删除时间
}
