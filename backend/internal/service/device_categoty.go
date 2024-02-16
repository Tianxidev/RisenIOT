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
	IDeviceCategory interface {
		// List 列表
		List(ctx context.Context, req *device.CategorySearchReq) (total, page int, list []*entity.SysDeviceCategory, err error)
		// KindGet 根据产品类型获取产品数据类型
		KindGet(ctx context.Context, kindId int) (list []*entity.SysDeviceCategory, err error)
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
