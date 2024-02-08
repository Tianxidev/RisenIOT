// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/api/v1/system"
	"backend/internal/model"
	"backend/internal/model/entity"
	"context"
)

type (
	ISysDept interface {
		GetList(ctx context.Context, req *system.DeptSearchReq) (list []*entity.SysDept, err error)
		GetFromCache(ctx context.Context) (list []*entity.SysDept, err error)
		// Add 添加部门
		Add(ctx context.Context, req *system.DeptAddReq) (err error)
		// Edit 部门修改
		Edit(ctx context.Context, req *system.DeptEditReq) (err error)
		Delete(ctx context.Context, id uint64) (err error)
		FindSonByParentId(deptList []*entity.SysDept, deptId uint64) []*entity.SysDept
		// GetListTree 部门树形菜单
		GetListTree(pid uint64, list []*entity.SysDept) (deptTree []*model.SysDeptTreeRes)
		// GetByDeptId 通过部门id获取部门信息
		GetByDeptId(ctx context.Context, deptId uint64) (dept *entity.SysDept, err error)
	}
)

var (
	localSysDept ISysDept
)

func SysDept() ISysDept {
	if localSysDept == nil {
		panic("implement not found for interface ISysDept, forgot register?")
	}
	return localSysDept
}

func RegisterSysDept(i ISysDept) {
	localSysDept = i
}
