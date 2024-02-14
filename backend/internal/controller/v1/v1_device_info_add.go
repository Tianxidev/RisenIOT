package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) InfoAdd(ctx context.Context, req *device.InfoAddReq) (res *device.InfoAddRes, err error) {
	err = service.DeviceInfo().Add(ctx, req)
	liberr.ErrIsNil(ctx, err, "添加设备信息失败")
	service.UserCtx().GetCtx(ctx).Message = "添加设备信息成功"
	return
}
