package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) GroupEdit(ctx context.Context, req *device.GroupEditReq) (res *device.GroupEditRes, err error) {
	err = service.DeviceGroup().Edit(ctx, req)
	liberr.ErrIsNil(ctx, err, "编辑设备分组失败")
	service.UserCtx().GetCtx(ctx).Message = "编辑设备分组成功"
	return
}
