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
		// Get 通过id获取
		Get(ctx context.Context, id int) (info *entity.SysDeviceKind, err error)
		// List 获取设备种类列表
		List(ctx context.Context, req *device.KindSearchReq) (total, page int, list []*entity.SysDeviceKind, err error)
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
