package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) KindAdd(ctx context.Context, req *device.KindAddReq) (res *device.KindAddRes, err error) {
	err = service.DeviceKind().Add(ctx, req)
	liberr.ErrIsNil(ctx, err, "添加产品类型失败")
	service.UserCtx().GetCtx(ctx).Message = "添加产品类型成功"
	return
}
