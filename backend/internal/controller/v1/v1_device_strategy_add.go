package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/v1/device"
)

func (c *ControllerDevice) StrategyAdd(ctx context.Context, req *device.StrategyAddReq) (res *device.StrategyAddRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
