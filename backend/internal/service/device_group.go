// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/api/v1/device"
	"context"
)

type (
	IDeviceGroup interface {
		// List 查询设备分组列表
		List(ctx context.Context) (res *device.GroupListRes, err error)
		// Add 添加设备分组
		Add(ctx context.Context, req *device.GroupAddReq) (err error)
		// Edit 编辑
		Edit(ctx context.Context, req *device.GroupEditReq) (err error)
		// Del 删除设备分组
		Del(ctx context.Context, ids []int) (err error)
	}
)

var (
	localDeviceGroup IDeviceGroup
)

func DeviceGroup() IDeviceGroup {
	if localDeviceGroup == nil {
		panic("implement not found for interface IDeviceGroup, forgot register?")
	}
	return localDeviceGroup
}

func RegisterDeviceGroup(i IDeviceGroup) {
	localDeviceGroup = i
}
