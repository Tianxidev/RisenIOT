package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/v1/device"
)

func (c *ControllerDevice) GatewayAdd(ctx context.Context, req *device.GatewayAddReq) (res *device.GatewayAddRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
