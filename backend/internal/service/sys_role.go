// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/api/v1/system"
	"backend/internal/model/entity"
	"context"
)

type (
	ISysRole interface {
		GetRoleListSearch(ctx context.Context, req *system.RoleListReq) (res *system.RoleListRes, err error)
		// GetRoleList 获取角色列表
		GetRoleList(ctx context.Context) (list []*entity.SysRole, err error)
		// AddRoleRule 添加角色权限
		AddRoleRule(ctx context.Context, ruleIds []uint, roleId int64) (err error)
		// DelRoleRule 删除角色权限
		DelRoleRule(ctx context.Context, roleId int64) (err error)
		AddRole(ctx context.Context, req *system.RoleAddReq) (err error)
		Get(ctx context.Context, id uint) (res *entity.SysRole, err error)
		// GetFilteredNamedPolicy 获取角色关联的菜单规则
		GetFilteredNamedPolicy(ctx context.Context, id uint) (gpSlice []int, err error)
		// EditRole 修改角色
		EditRole(ctx context.Context, req *system.RoleEditReq) (err error)
		// DeleteByIds 删除角色
		DeleteByIds(ctx context.Context, ids []int64) (err error)
	}
)

var (
	localSysRole ISysRole
)

func SysRole() ISysRole {
	if localSysRole == nil {
		panic("implement not found for interface ISysRole, forgot register?")
	}
	return localSysRole
}

func RegisterSysRole(i ISysRole) {
	localSysRole = i
}
