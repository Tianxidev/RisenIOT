package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/v1/device"
)

func (c *ControllerDevice) InfoGet(ctx context.Context, req *device.InfoGetReq) (res *device.InfoGetRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
