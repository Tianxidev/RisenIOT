package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/v1/device"
)

func (c *ControllerDevice) CategoryEdit(ctx context.Context, req *device.CategoryEditReq) (res *device.CategoryEditRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
