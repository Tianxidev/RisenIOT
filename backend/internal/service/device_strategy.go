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
	IDeviceStrategy interface {
		// List 查询策略列表
		List(ctx context.Context, req *device.StrategySearchReq) (*device.StrategySearchRes, error)
	}
)

var (
	localDeviceStrategy IDeviceStrategy
)

func DeviceStrategy() IDeviceStrategy {
	if localDeviceStrategy == nil {
		panic("implement not found for interface IDeviceStrategy, forgot register?")
	}
	return localDeviceStrategy
}

func RegisterDeviceStrategy(i IDeviceStrategy) {
	localDeviceStrategy = i
}
