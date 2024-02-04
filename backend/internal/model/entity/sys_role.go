// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id        uint        `json:"id"        ` //
	Status    uint        `json:"status"    ` // 状态;0:禁用;1:正常
	ListOrder uint        `json:"listOrder" ` // 排序
	Name      string      `json:"name"      ` // 角色名称
	Remark    string      `json:"remark"    ` // 备注
	DataScope uint        `json:"dataScope" ` // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
