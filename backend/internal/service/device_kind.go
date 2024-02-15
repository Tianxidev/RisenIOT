// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/api/v1/device"
	"backend/internal/model/entity"
	"context"
)

type (
	IDeviceKind interface {
		Get(ctx context.Context, id int) (info *entity.SysDeviceKind, err error)
		List(ctx context.Context, req *device.KindSearchReq) (total, page int, list []*entity.SysDeviceKind, err error)
		// Add 添加产品类型
		Add(ctx context.Context, req *device.KindAddReq) (err error)
		// Edit 编辑产品类型
		Edit(ctx context.Context, req *device.KindEditReq) (err error)
		// Del 删除产品类型
		Del(ctx context.Context, id []int) (err error)
	}
)

var (
	localDeviceKind IDeviceKind
)

func DeviceKind() IDeviceKind {
	if localDeviceKind == nil {
		panic("implement not found for interface IDeviceKind, forgot register?")
	}
	return localDeviceKind
}

func RegisterDeviceKind(i IDeviceKind) {
	localDeviceKind = i
}
