package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) GroupList(ctx context.Context, req *device.GroupListReq) (res *device.GroupListRes, err error) {
	res, err = service.DeviceGroup().GroupList(ctx)
	liberr.ErrIsNil(ctx, err, "获取设备分组列表失败")
	service.UserCtx().GetCtx(ctx).Message = "获取设备分组列表成功"
	return
}
