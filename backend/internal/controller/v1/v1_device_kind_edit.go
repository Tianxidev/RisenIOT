package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) KindEdit(ctx context.Context, req *device.KindEditReq) (res *device.KindEditRes, err error) {
	err = service.DeviceKind().Edit(ctx, req)
	liberr.ErrIsNil(ctx, err, "编辑设备类型失败")
	service.UserCtx().GetCtx(ctx).Message = "编辑设备类型成功"
	return
}
