package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) StrategySearch(ctx context.Context, req *device.StrategySearchReq) (res *device.StrategySearchRes, err error) {
	res, err = service.DeviceStrategy().List(ctx, req)
	liberr.ErrIsNil(ctx, err, "查询策略列表失败")
	service.UserCtx().GetCtx(ctx).Message = "查询策略列表成功"
	return
}
