package v1

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/v1/device"
)

func (c *ControllerDevice) StatusSearch(ctx context.Context, req *device.StatusSearchReq) (res *device.StatusSearchRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
