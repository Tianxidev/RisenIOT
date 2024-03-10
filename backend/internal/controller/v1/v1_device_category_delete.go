package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) CategoryDelete(ctx context.Context, req *device.CategoryDeleteReq) (res *device.CategoryDeleteRes, err error) {
	err = service.DeviceCategory().Del(ctx, req)
	liberr.ErrIsNil(ctx, err, "删除设备数据类型失败")
	service.UserCtx().GetCtx(ctx).Message = "删除设备数据类型成功"
	return
}
