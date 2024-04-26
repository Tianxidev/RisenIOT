package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) StrategyDel(ctx context.Context, req *device.StrategyDelReq) (res *device.StrategyDelRes, err error) {
	res, err = service.DeviceStrategy().Del(ctx, req)
	liberr.ErrIsNil(ctx, err, "删除策略失败")
	service.UserCtx().GetCtx(ctx).Message = "删除策略成功"
	return
}
