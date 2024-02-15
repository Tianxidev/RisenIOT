package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) GroupAdd(ctx context.Context, req *device.GroupAddReq) (res *device.GroupAddRes, err error) {
	err = service.DeviceGroup().Add(ctx, req)
	liberr.ErrIsNil(ctx, err, "添加设备分组失败")
	service.UserCtx().GetCtx(ctx).Message = "添加设备分组成功"
	return
}
