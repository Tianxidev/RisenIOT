package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) KindDelete(ctx context.Context, req *device.KindDeleteReq) (res *device.KindDeleteRes, err error) {
	err = service.DeviceKind().Del(ctx, req.Ids)
	liberr.ErrIsNil(ctx, err, "删除产品类型失败")
	service.UserCtx().GetCtx(ctx).Message = "删除产品类型成功"
	return
}
