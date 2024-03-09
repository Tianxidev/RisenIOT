package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/home"
)

func (c *ControllerHome) Overview(ctx context.Context, req *home.OverviewReq) (res *home.OverviewRes, err error) {
	res = &home.OverviewRes{}

	// 查询总设备数量
	totalDevice, err := service.DeviceInfo().Count(ctx)
	liberr.ErrIsNil(ctx, err, "获取设备数量失败")

	// 查询在线设备数量
	totalOnlineDevice, err := service.DeviceInfo().OnlineCount(ctx)
	liberr.ErrIsNil(ctx, err, "获取在线设备数量失败")

	// 查询离线设备数量
	totalOfflineDevice, err := service.DeviceInfo().OfflineCount(ctx)
	liberr.ErrIsNil(ctx, err, "获取离线设备数量失败")

	// 返回数据
	res.TotalDevice = totalDevice
	res.Online = totalOnlineDevice
	res.Offline = totalOfflineDevice

	return
}
