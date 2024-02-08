// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/api/v1/system"
	"context"
)

type (
	ISysDictData interface {
		// GetDictWithDataByType 通过字典键类型获取选项
		GetDictWithDataByType(ctx context.Context, req *system.GetDictReq) (dict *system.GetDictRes, err error)
		// List 获取字典数据
		List(ctx context.Context, req *system.DictDataSearchReq) (res *system.DictDataSearchRes, err error)
		Add(ctx context.Context, req *system.DictDataAddReq, userId uint64) (err error)
		// Get 获取字典数据
		Get(ctx context.Context, dictCode uint) (res *system.DictDataGetRes, err error)
		// Edit 修改字典数据
		Edit(ctx context.Context, req *system.DictDataEditReq, userId uint64) (err error)
		// Delete 删除字典数据
		Delete(ctx context.Context, ids []int) (err error)
	}
)

var (
	localSysDictData ISysDictData
)

func SysDictData() ISysDictData {
	if localSysDictData == nil {
		panic("implement not found for interface ISysDictData, forgot register?")
	}
	return localSysDictData
}

func RegisterSysDictData(i ISysDictData) {
	localSysDictData = i
}
