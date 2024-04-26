package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) StrategyAdd(ctx context.Context, req *device.StrategyAddReq) (res *device.StrategyAddRes, err error) {
	res, err = service.DeviceStrategy().Add(ctx, req)
	liberr.ErrIsNil(ctx, err, "添加策略失败")
	return
}
