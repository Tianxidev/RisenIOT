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
		// Add 添加策略
		Add(ctx context.Context, req *device.StrategyAddReq) (*device.StrategyAddRes, error)
		// List 查询策略列表
		List(ctx context.Context, req *device.StrategySearchReq) (*device.StrategySearchRes, error)
		// Edit 编辑策略
		Edit(ctx context.Context, req *device.StrategyEditReq) (*device.StrategyEditRes, error)
		// Del 删除策略
		Del(ctx context.Context, req *device.StrategyDelReq) (*device.StrategyDelRes, error)
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
