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

	"github.com/gogf/gf/v2/os/gtime"
)

type (
	IDeviceInfo interface {
		// List 获取设备信息列表
		List(ctx context.Context, req *device.InfoSearchReq) (total, page int, list []*entity.SysDeviceInfo, err error)
		// Auth 设备授权
		Auth(ctx context.Context, sn, pwd string) (status bool, err error)
		// Info 获取设备信息
		Info(ctx context.Context, id int, sn string) (deviceInfo *model.DeviceAllInfo, err error)
		// Add 添加设备信息
		Add(ctx context.Context, req *device.InfoAddReq) (err error)
		// Edit 编辑设备信息
		Edit(ctx context.Context, req *device.InfoEditReq) (err error)
		// Delete 删除设备信息
		Delete(ctx context.Context, ids []int) (err error)
		// UpdateDataLastTime 更新设备数据最后上报时间
		UpdateDataLastTime(ctx context.Context, id int) (err error)
		// Count 统计设备数量
		Count(ctx context.Context) (total int, err error)
		// OnlineCount 在线设备数量
		OnlineCount(ctx context.Context) (total int, err error)
		// OfflineCount 离线设备数量
		OfflineCount(ctx context.Context) (total int, err error)
		// KeepAlive 设备保活状态检查
		KeepAlive(ctx context.Context, time *gtime.Time)
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
