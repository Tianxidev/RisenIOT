package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/v1/device"
)

func (c *ControllerDevice) CategoryDataGet(ctx context.Context, req *device.CategoryDataGetReq) (res *device.CategoryDataGetRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
