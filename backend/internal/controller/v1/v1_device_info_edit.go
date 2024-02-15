package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) InfoEdit(ctx context.Context, req *device.InfoEditReq) (res *device.InfoEditRes, err error) {
	err = service.DeviceInfo().Edit(ctx, req)
	liberr.ErrIsNil(ctx, err, "编辑设备信息失败")
	service.UserCtx().GetCtx(ctx).Message = "编辑设备信息成功"
	return
}
