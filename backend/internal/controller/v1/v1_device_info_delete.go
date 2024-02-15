package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) InfoDelete(ctx context.Context, req *device.InfoDeleteReq) (res *device.InfoDeleteRes, err error) {
	err = service.DeviceInfo().Delete(ctx, req.Ids)
	liberr.ErrIsNil(ctx, err, "删除设备信息失败")
	service.UserCtx().GetCtx(ctx).Message = "删除设备信息成功"
	return
}
