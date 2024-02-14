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
		// GroupList 查询设备分组列表
		GroupList(ctx context.Context) (res *device.GroupListRes, err error)
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
