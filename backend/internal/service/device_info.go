// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/api/v1/device"
	"backend/internal/model"
	"backend/internal/model/entity"
	"context"
)

type (
	IDeviceInfo interface {
		// List 获取设备信息列表
		List(ctx context.Context, req *device.InfoSearchReq) (total, page int, list []*entity.SysDeviceInfo, err error)
		Auth(ctx context.Context, sn, pwd string) (status bool, err error)
		// Info 获取设备所有信息
		Info(ctx context.Context, id int, sn string) (info *model.DeviceAllInfo, err error)
	}
)

var (
	localDeviceInfo IDeviceInfo
)

func DeviceInfo() IDeviceInfo {
	if localDeviceInfo == nil {
		panic("implement not found for interface IDeviceInfo, forgot register?")
	}
	return localDeviceInfo
}

func RegisterDeviceInfo(i IDeviceInfo) {
	localDeviceInfo = i
}
