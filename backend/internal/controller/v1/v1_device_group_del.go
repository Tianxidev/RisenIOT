package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) GroupDel(ctx context.Context, req *device.GroupDelReq) (res *device.GroupDelRes, err error) {
	err = service.DeviceGroup().Del(ctx, req.Ids)
	liberr.ErrIsNil(ctx, err, "删除设备分组失败")
	service.UserCtx().GetCtx(ctx).Message = "删除设备分组成功"
	return
}
