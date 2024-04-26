package v1

import (
	"backend/internal/service"
	"backend/library/liberr"
	"context"

	"backend/api/v1/device"
)

func (c *ControllerDevice) StrategyEdit(ctx context.Context, req *device.StrategyEditReq) (res *device.StrategyEditRes, err error) {
	res, err = service.DeviceStrategy().Edit(ctx, req)
	liberr.ErrIsNil(ctx, err, "编辑策略失败")
	return
}
