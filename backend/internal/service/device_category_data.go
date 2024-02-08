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
	IDeviceCategoryData interface {
		Add(ctx context.Context, req *device.CategoryDataAddReq) (err error)
	}
)

var (
	localDeviceCategoryData IDeviceCategoryData
)

func DeviceCategoryData() IDeviceCategoryData {
	if localDeviceCategoryData == nil {
		panic("implement not found for interface IDeviceCategoryData, forgot register?")
	}
	return localDeviceCategoryData
}

func RegisterDeviceCategoryData(i IDeviceCategoryData) {
	localDeviceCategoryData = i
}
