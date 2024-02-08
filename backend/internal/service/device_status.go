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
	IDeviceStatus interface {
		// List 获取状态列表
		List(ctx context.Context, req *device.StatusSearchReq) (total, page int, list []*entity.SysDeviceStatus, err error)
		// ChangeStatus 修改状态
		ChangeStatus(ctx context.Context, deviceId int, status int) error
	}
)

var (
	localDeviceStatus IDeviceStatus
)

func DeviceStatus() IDeviceStatus {
	if localDeviceStatus == nil {
		panic("implement not found for interface IDeviceStatus, forgot register?")
	}
	return localDeviceStatus
}

func RegisterDeviceStatus(i IDeviceStatus) {
	localDeviceStatus = i
}
