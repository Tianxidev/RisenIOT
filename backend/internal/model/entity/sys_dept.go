// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDept is the golang structure for table sys_dept.
type SysDept struct {
	DeptId    int64       `json:"deptId"    ` // 部门id
	ParentId  int64       `json:"parentId"  ` // 父部门id
	Ancestors string      `json:"ancestors" ` // 祖级列表
	DeptName  string      `json:"deptName"  ` // 部门名称
	OrderNum  int         `json:"orderNum"  ` // 显示顺序
	Leader    string      `json:"leader"    ` // 负责人
	Phone     string      `json:"phone"     ` // 联系电话
	Email     string      `json:"email"     ` // 邮箱
	Status    uint        `json:"status"    ` // 部门状态（0正常 1停用）
	CreatedBy uint64      `json:"createdBy" ` // 创建人
	UpdatedBy int64       `json:"updatedBy" ` // 修改人
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 修改时间
	DeletedAt *gtime.Time `json:"deletedAt" ` // 删除时间
}
