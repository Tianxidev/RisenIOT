package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) CategoryAdd(ctx context.Context, req *device.CategoryAddReq) (res *device.CategoryAddRes, err error) {
	res = &device.CategoryAddRes{}

	err = service.DeviceCategory().Add(ctx, req)
	liberr.ErrIsNil(ctx, err, "添加设备数据类型失败")
	service.UserCtx().GetCtx(ctx).Message = "添加设备数据类型成功"
	return
}
