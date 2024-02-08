// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/internal/model/entity"
	"context"
)

type (
	IDeviceCategory interface {
		KindGet(ctx context.Context, kindId int) (list []*entity.SysDeviceCategoty, err error)
	}
)

var (
	localDeviceCategory IDeviceCategory
)

func DeviceCategory() IDeviceCategory {
	if localDeviceCategory == nil {
		panic("implement not found for interface IDeviceCategory, forgot register?")
	}
	return localDeviceCategory
}

func RegisterDeviceCategory(i IDeviceCategory) {
	localDeviceCategory = i
}
