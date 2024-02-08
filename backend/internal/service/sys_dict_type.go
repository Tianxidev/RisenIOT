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
	ISysDictType interface {
		// List 字典类型列表
		List(ctx context.Context, req *system.DictTypeSearchReq) (res *system.DictTypeSearchRes, err error)
		// Add 添加字典类型
		Add(ctx context.Context, req *system.DictTypeAddReq, userId uint64) (err error)
		// Edit 修改字典类型
		Edit(ctx context.Context, req *system.DictTypeEditReq, userId uint64) (err error)
		Get(ctx context.Context, req *system.DictTypeGetReq) (dictType *entity.SysDictType, err error)
		// ExistsDictType 检查类型是否已经存在
		ExistsDictType(ctx context.Context, dictType string, dictId ...int64) (err error)
		// Delete 删除字典类型
		Delete(ctx context.Context, dictIds []int) (err error)
		// GetAllDictType 获取所有正常状态下的字典类型
		GetAllDictType(ctx context.Context) (list []*entity.SysDictType, err error)
	}
)

var (
	localSysDictType ISysDictType
)

func SysDictType() ISysDictType {
	if localSysDictType == nil {
		panic("implement not found for interface ISysDictType, forgot register?")
	}
	return localSysDictType
}

func RegisterSysDictType(i ISysDictType) {
	localSysDictType = i
}
